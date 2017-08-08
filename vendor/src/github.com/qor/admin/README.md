## QOR Admin

Instantly create a beautiful, cross platform, configurable Admin Interface and API for managing your data in minutes.

[![GoDoc](https://godoc.org/github.com/qor/admin?status.svg)](https://godoc.org/github.com/qor/admin)

**For security issues, please send us an email to security@getqor.com and give us time to respond BEFORE posting as an issue or reporting on public forums.**

## Documentation

<https://doc.getqor.com/chapter2/setup.html>

## Features

- Admin Interface for managing data
- JSON API
- Association handling
- Search and filtering
- Actions/Batch Actions
- Authentication and Authorization (based on Permissions)
- Extendability

## Quick Start

```go
package main

import (
    "fmt"
    "net/http"

    "github.com/jinzhu/gorm"
    _ "github.com/mattn/go-sqlite3"
    "github.com/qor/qor"
    "github.com/qor/admin"
)

// Create a GORM-backend model
type User struct {
  gorm.Model
  Name string
}

// Create another GORM-backend model
type Product struct {
  gorm.Model
  Name        string
  Description string
}

func main() {
  DB, _ := gorm.Open("sqlite3", "demo.db")
  DB.AutoMigrate(&User{}, &Product{})

  // Initalize
  Admin := admin.New(&qor.Config{DB: DB})

  // Create resources from GORM-backend model
  Admin.AddResource(&User{})
  Admin.AddResource(&Product{})

  // Register route
  mux := http.NewServeMux()
  // amount to /admin, so visit `/admin` to view the admin interface
  Admin.MountTo("/admin", mux)

  fmt.Println("Listening on: 9000")
  http.ListenAndServe(":9000", mux)
}
```

`go run main.go` and visit `localhost:9000/admin` to see the result !

## General Setting

### Site Name

Use `SetSiteName` to set QOR Admin's HTML title, the name will also be used to auto-load javascripts and stylesheet files that you can provide for customizing the admin interface.

For example, say you set the Site Name as `QOR Demo`, admin will look up `qor_demo.js`, `qor_demo.css` in [QOR view paths](#customizing-views), and load them if present.

```go
Admin.SetSiteName("QOR DEMO")
```

### Dashboard

QOR Admin provides a default dashboard page with some dummy text. If you want to customize the dashboard, you can create a file `dashboard.tmpl` in [QOR view paths](#customizing-views), QOR Admin will load it as golang templates when rendering the dashboard.

If you want to disable the dashboard, you can redirect it to some other page, for example:

```go
Admin.GetRouter().Get("/", func(c *admin.Context) {
  http.Redirect(c.Writer, c.Request, "/admin/clients", http.StatusSeeOther)
})
```

### Authentication

QOR Admin provides a flexible authorization solution. With it, you could integrate with your current authorization method.

What you need to do is implement an `Auth` interface like below, and set it in the Admin value.

```go
type Auth interface {
	GetCurrentUser(*Context) qor.CurrentUser // get current user, if don't have permission, then return nil
	LoginURL(*Context) string // get login url, if don't have permission, will redirect to this url
	LogoutURL(*Context) string // get logout url, if click logout link from admin interface, will visit this page
}
```

Here is an example:

```go
type Auth struct{}

func (Auth) LoginURL(*admin.Context) string {
  return "/login"
}

func (Auth) LogoutURL(*Context) string {
  return "/logout"
}

func (Auth) GetCurrentUser(c *admin.Context) qor.CurrentUser {
  if userid, err := c.Request.Cookie("userid"); err == nil {
    var user User
    if !DB.First(&user, "id = ?", userid.Value).RecordNotFound() {
      return &user
    }
  }
  return nil
}

func (u User) DisplayName() string {
  return u.Name
}

// Register Auth for QOR Admin
Admin.SetAuth(&Auth{})
```

### Menu

#### Register a Menu

It is possible to define a nested menu structure for the admin interface.

```go
Admin.AddMenu(&admin.Menu{Name: "Dashboard", Link: "/admin"})

// Register nested menu
Admin.AddMenu(&admin.Menu{Name: "menu", Link: "/link", Ancestors: []string{"Dashboard"}})

// Register menu with permission
Admin.AddMenu(&admin.Menu{Name: "Report", Link: "/admin", Permission: roles.Allow(roles.Read, "admin")})
```

#### Add Resources to a menu

```go
Admin.AddResource(&User{})

Admin.AddResource(&Product{}, &admin.Config{Menu: []string{"Product Management"}})
Admin.AddResource(&Color{}, &admin.Config{Menu: []string{"Product Management"}})
Admin.AddResource(&Size{}, &admin.Config{Menu: []string{"Product Management"}})

Admin.AddResource(&Order{}, &admin.Config{Menu: []string{"Order Management"}})
```

If you don't want a resource to be displayed in the menu, pass the Invisible option:

```go
Admin.AddResource(&User{}, &admin.Config{Invisible: true})
```

### Internationalization

To translate admin interface to a new language, you could use `i18n` [https://github.com/qor/i18n](https://github.com/qor/i18n)

## Working with a Resource

Every QOR Admin Resource needs a [GORM-backend](https://github.com/jinzhu/gorm) model. Once you have defined the model you can create a QOR Admin resource: `Admin.AddResource(&Product{})`

Once a resource has been added, QOR Admin will generate the admin interface to manage it, including a RESTFul JSON API.

So for above example, you could visit `localhost:9000/admin/products` to manage `Product` in the HTML admin interface, or use the RESTFul JSON api `localhost:9000/admin/products.json` to perform CRUD activities.

### Customizing CRUD pages

```go
// Set attributes will be shown in the index page
// show given attributes
order.IndexAttrs("User", "PaymentAmount", "ShippedAt", "CancelledAt", "State", "ShippingAddress")
// show all attributes except `State`
order.IndexAttrs("-State")

// Set attributes will be shown in the new page
order.NewAttrs("User", "PaymentAmount", "ShippedAt", "CancelledAt", "State", "ShippingAddress")
// show all attributes except `State`
order.NewAttrs("-State")
// Structure the new form to make it tidy and clean with `Section`
product.NewAttrs(
  &admin.Section{
		Title: "Basic Information",
		Rows: [][]string{
			{"Name"},
			{"Code", "Price"},
		}
  },
  &admin.Section{
		Title: "Organization",
		Rows: [][]string{
			{"Category", "Collections", "MadeCountry"},
    }
  },
  "Description",
  "ColorVariations",
}

// Set attributes will be shown for the edit page, similiar with new page
order.EditAttrs("User", "PaymentAmount", "ShippedAt", "CancelledAt", "State", "ShippingAddress")

// Set attributes will be shown for the show page, similiar with new page
// If ShowAttrs haven't been configured, there will be no show page generated, by will show the edit from instead
order.ShowAttrs("User", "PaymentAmount", "ShippedAt", "CancelledAt", "State", "ShippingAddress")
```

### Search

It is possible to specify database table columns as search attributes, using `SearchAttrs`, the columns will be used to perform any search queries. It is also possible to specify nested relations.

```go
// Search products with its name, code, category's name, brand's name
product.SearchAttrs("Name", "Code", "Category.Name", "Brand.Name")
```

If you want to fully customize the search function, you could set the `SearchHandler`:

```go
order.SearchHandler = func(keyword string, context *qor.Context) *gorm.DB {
  // search orders
}
```

#### Search Center

You might want to search a broad range of resources from a single web page, in this case `Search Center` is for you!  Simply add resources that you want to be searchable to the Admin value's search center:

```go
// add resource `product`, `user`, `order` to search resources
Admin.AddSearchResource(product, user, order)
```

[Search Center Online Demo](http://demo.getqor.com/admin/!search)

### Scopes

You can define scopes to filter data with given conditions, for example:

```go
// Only show actived users
user.Scope(&admin.Scope{Name: "Active", Handle: func(db *gorm.DB, context *qor.Context) *gorm.DB {
  return db.Where("active = ?", true)
}})
```

#### Group Scopes

```go
order.Scope(&admin.Scope{Name: "Paid", Group: "State", Handle: func(db *gorm.DB, context *qor.Context) *gorm.DB {
  return db.Where("state = ?", "paid")
}})

order.Scope(&admin.Scope{Name: "Shipped", Group: "State", Handle: func(db *gorm.DB, context *qor.Context) *gorm.DB {
  return db.Where("state = ?", "shipped")
}})
```

[Scopes Online Demo](http://demo.getqor.com/admin/products)

### Actions

QOR Admin has the notion of four action modes:

* Bulk actions (will be shown in index page as bulk actions)
* Edit form action (will be shown in edit page)
* Show page action (will be shown in show page)
* Menu item action (will be shown in table's menu)

You can register an Action of any mode using the `Action` method, along with `Modes` values to contol where to show them:

```go
product.Action(&admin.Action{
	Name: "enable",
	Handle: func(actionArgument *admin.ActionArgument) error {
    // `FindSelectedRecords` => return selected record in bulk action mode, return current record in other mode
		for _, record := range actionArgument.FindSelectedRecords() {
			actionArgument.Context.DB.Model(record.(*models.Product)).Update("disabled", false)
		}
		return nil
	},
	Modes: []string{"index", "edit", "show", "menu_item"},
})

// Register Actions need user's input
order.Action(&admin.Action{
  Name: "Ship",
  Handle: func(argument *admin.ActionArgument) error {
    trackingNumberArgument := argument.Argument.(*trackingNumberArgument)
    for _, record := range argument.FindSelectedRecords() {
      argument.Context.GetDB().Model(record).UpdateColumn("tracking_number", trackingNumberArgument.TrackingNumber)
    }
    return nil
  },
  Resource: Admin.NewResource(&trackingNumberArgument{}),
  Modes: []string{"show", "menu_item"},
})

// the ship action's argument
type trackingNumberArgument struct {
  TrackingNumber string
}

// Use `Visible` to hide registered Action in some case
order.Action(&admin.Action{
  Name: "Cancel",
  Handle: func(argument *admin.ActionArgument) error {
    // cancel the order
  },
  Visible: func(record interface{}) bool {
    if order, ok := record.(*models.Order); ok {
      for _, state := range []string{"draft", "checkout", "paid", "processing"} {
        if order.State == state {
          return true
        }
      }
    }
    return false
  },
  Modes: []string{"show", "menu_item"},
})
```

### Customizing the Form

By default, management pages in QOR Admin are rendered based on your resource's fields' data types and relations. The default should satisfy most use cases, however should you need to you can customize the rendering by overwritting the `Meta` definition.

There are some Meta types that have been predefined, including `string`, `password`, `date`, `datetime`, `rich_editor`, `select_one`, `select_many` and so on (see full list here: [qor admin form templates](https://github.com/qor/admin/tree/master/views/metas/form "qor admin form templates")). QOR Admin will auto select a type for `Meta` based on a field's data type. For example, if a field's type is `time.Time`, QOR Admin will determine `datetime` as the type.

```go
// Change the Meta type of `Password` field in User resource from `string` (default value) to `password`
user.Meta(&admin.Meta{Name: "Password", Type: "password"})

// Change the Meta type of `Gender` field in User resource from `string` (default value) to `select_one`, with options `M` | `F`
user.Meta(&admin.Meta{Name: "Gender", Type: "select_one", Collection: []string{"M", "F"}})
```

### Authorization and Permissions

Authorization in QOR Admin is based on setting Permissions per Role. QOR Admin uses [https://github.com/qor/roles](https://github.com/qor/roles) for Permission management, please refer to it's documentation for information on how to define Roles and Permissions.

```go
// CRUD permission for admin users, deny create permission for manager
user := Admin.AddResource(&User{}, &admin.Config{Permission: roles.Allow(roles.CRUD, "admin").Deny(roles.Create, "manager")})

// For user's Email field, allow CRUD for admin users, deny update for manager
user.Meta(&admin.Meta{Name: "Email", Permission: roles.Allow(roles.CRUD, "admin").Deny(roles.Create, "manager")})
```

### An automagic RESTFul API

The RESTFul API shares the same configuration as your admin interface, including actions and permissions - so after you have configured your admin interface, you will get an API for free!

## Extendability

#### Configuring QOR Admin Resources Automatically

If your model has the following two methods defined, they will be called when registering:

```go
func ConfigureQorResourceBeforeInitialize(resource) {
  // resource.(*admin.Resource)
}

func ConfigureQorResource(resource) {
  // resource.(*admin.Resource)
}
```

#### Configuring QOR Admin Meta Automatically

If your field's type has the following two methods defined, they will be called when registering:

```go
func ConfigureQorMetaBeforeInitialize(meta) {
  // resource.(*admin.Meta)
}

func ConfigureMetaInterface(meta) {
  // resource.(*admin.Meta)
}
```

#### Using a Theme

A custom theme can be applied using a custom javascript and css file, for example to make a product page look super fancy. To apply a custom theme, provide the theme name using the `UseTheme` method, this will load `assets/javascripts/fancy.js` and `assets/stylesheets/fancy.css` from [QOR view paths](#customizing-views)

```go
product.UseTheme("fancy")
```

#### Customizing Views

QOR Admin will look up templates in QOR Admin view paths and use them to render any admin page. By placing your own templates in `{current path}/app/views/qor` you can extend your application by customizing it's views. If you want to customize your views from other places, you could register any new paths with `admin.RegisterViewPath`.

Customize Views Rules:

* To overwrite a template, create a file with the same name under `{current path}/app/views/qor`.
* To overwrite templates for a specific resource, put templates with the same name in `{qor view paths}/{resource param}`, for example `{current path}/app/views/qor/products/index.tmpl`.
* To overwrite templates for resources using a theme, put templates with the same name in `{qor view paths}/themes/{theme name}`.

#### Registering HTTP routes

Qor admin uses Qor's Router.

```go
router := Admin.GetRouter()

router.Get("/path", func(context *admin.Context) {
    // do something here
})

router.Post("/path", func(context *admin.Context) {
    // do something here
})

router.Put("/path", func(context *admin.Context) {
    // do something here
})

router.Delete("/path", func(context *admin.Context) {
    // do something here
})

// naming route
router.Get("/path/:name", func(context *admin.Context) {
    context.Request.URL.Query().Get(":name")
})

// regexp support
router.Get("/path/:name[world]", func(context *admin.Context) { // "/hello/world"
    context.Request.URL.Query().Get(":name")
})

router.Get("/path/:name[\\d+]", func(context *admin.Context) { // "/hello/123"
    context.Request.URL.Query().Get(":name")
})
```

#### Plugins

There are a few plugins created for QOR already, you can find some of them at [https://github.com/qor](https://github.com/qor), visit them to learn more about how to extend QOR.

## Live DEMO

* Live Demo [http://demo.getqor.com/admin](http://demo.getqor.com/admin)
* Source Code of Live Demo [https://github.com/qor/qor-example](https://github.com/qor/qor-example)

## Q & A

* How to integrate with beego

```go
mux := http.NewServeMux()
Admin.MountTo("/admin", mux)

beego.Handler("/admin/*", mux)
beego.Run()
```

* How to integrate with Gin

```go
mux := http.NewServeMux()
Admin.MountTo("/admin", mux)

r := gin.Default()
r.Any("/admin/*w", gin.WrapH(mux))
r.Run()
```

## License

Released under the [MIT License](http://opensource.org/licenses/MIT).
