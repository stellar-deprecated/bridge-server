package admin

import (
	"path"

	"github.com/qor/qor"
	"github.com/qor/roles"
)

// GetMenus get all sidebar menus for admin
func (admin Admin) GetMenus() []*Menu {
	return admin.menus
}

// AddMenu add a menu to admin sidebar
func (admin *Admin) AddMenu(menu *Menu) *Menu {
	menu.router = admin.router
	admin.menus = appendMenu(admin.menus, menu.Ancestors, menu)
	return menu
}

// GetMenu get sidebar menu with name
func (admin Admin) GetMenu(name string) *Menu {
	return getMenu(admin.menus, name)
}

////////////////////////////////////////////////////////////////////////////////
// Sidebar Menu
////////////////////////////////////////////////////////////////////////////////

// Menu admin sidebar menu definiation
type Menu struct {
	Name         string
	Link         string
	RelativePath string
	Priority     int
	Ancestors    []string
	Permissioner HasPermissioner
	Permission   *roles.Permission

	subMenus []*Menu
	router   *Router
}

// URL return menu's URL
func (menu Menu) URL() string {
	if menu.Link != "" {
		return menu.Link
	}

	if (menu.router != nil) && (menu.RelativePath != "") {
		return path.Join(menu.router.Prefix, menu.RelativePath)
	}

	return menu.RelativePath
}

// HasPermission check menu has permission or not
func (menu Menu) HasPermission(mode roles.PermissionMode, context *qor.Context) bool {
	if menu.Permission != nil {
		return menu.Permission.HasPermission(mode, context.Roles...)
	}

	if menu.Permissioner != nil {
		return menu.Permissioner.HasPermission(mode, context)
	}

	return true
}

// GetSubMenus get submenus for a menu
func (menu *Menu) GetSubMenus() []*Menu {
	return menu.subMenus
}

func getMenu(menus []*Menu, name string) *Menu {
	for _, m := range menus {
		if m.Name == name {
			return m
		}

		if len(m.subMenus) > 0 {
			if mc := getMenu(m.subMenus, name); mc != nil {
				return mc
			}
		}
	}

	return nil
}

func generateMenu(menus []string, menu *Menu) *Menu {
	menuCount := len(menus)
	for index := range menus {
		menu = &Menu{Name: menus[menuCount-index-1], subMenus: []*Menu{menu}}
	}

	return menu
}

func appendMenu(menus []*Menu, ancestors []string, menu *Menu) []*Menu {
	if len(ancestors) > 0 {
		for _, m := range menus {
			if m.Name != ancestors[0] {
				continue
			}

			if len(ancestors) > 1 {
				m.subMenus = appendMenu(m.subMenus, ancestors[1:], menu)
			} else {
				m.subMenus = appendMenu(m.subMenus, []string{}, menu)
			}

			return menus
		}
	}

	var newMenu = generateMenu(ancestors, menu)
	var added bool
	if len(menus) == 0 {
		menus = append(menus, newMenu)
	} else if newMenu.Priority > 0 {
		for idx, menu := range menus {
			if menu.Priority > newMenu.Priority || menu.Priority <= 0 {
				menus = append(menus[0:idx], append([]*Menu{newMenu}, menus[idx:]...)...)
				added = true
				break
			}
		}
		if !added {
			menus = append(menus, menu)
		}
	} else {
		if newMenu.Priority < 0 {
			for idx := len(menus) - 1; idx >= 0; idx-- {
				menu := menus[idx]
				if menu.Priority < newMenu.Priority || menu.Priority == 0 {
					menus = append(menus[0:idx+1], append([]*Menu{newMenu}, menus[idx+1:]...)...)
					added = true
					break
				}
			}

			if !added {
				menus = append(menus, menu)
			}
		} else {
			for idx := len(menus) - 1; idx >= 0; idx-- {
				menu := menus[idx]
				if menu.Priority >= 0 {
					menus = append(menus[0:idx+1], append([]*Menu{newMenu}, menus[idx+1:]...)...)
					added = true
					break
				}
			}

			if !added {
				menus = append([]*Menu{menu}, menus...)
			}
		}
	}

	return menus
}
