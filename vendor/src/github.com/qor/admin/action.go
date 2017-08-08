package admin

import (
	"path"
	"reflect"
	"strings"

	"github.com/qor/qor"
	"github.com/qor/qor/utils"
	"github.com/qor/roles"
)

// Action register action for qor resource
func (res *Resource) Action(action *Action) *Action {
	for _, a := range res.Actions {
		if a.Name == action.Name {
			if action.Label != "" {
				a.Label = action.Label
			}

			if action.Method != "" {
				a.Method = action.Method
			}

			if action.URL != nil {
				a.URL = action.URL
			}

			if action.URLOpenType != "" {
				a.URLOpenType = action.URLOpenType
			}

			if action.Visible != nil {
				a.Visible = action.Visible
			}

			if action.Handler != nil {
				a.Handler = action.Handler
			}

			if len(action.Modes) != 0 {
				a.Modes = action.Modes
			}

			if action.Resource != nil {
				a.Resource = action.Resource
			}

			if action.Permission != nil {
				a.Permission = action.Permission
			}

			*action = *a
			return a
		}
	}

	if action.Label == "" {
		action.Label = utils.HumanizeString(action.Name)
	}

	if action.Method == "" {
		if action.URL != nil {
			action.Method = "GET"
		} else {
			action.Method = "PUT"
		}
	}

	if action.URLOpenType == "" {
		if action.Resource != nil {
			action.URLOpenType = "bottomsheet"
		} else if action.Method == "GET" {
			action.URLOpenType = "_blank"
		}
	}

	res.Actions = append(res.Actions, action)

	// Register Actions into Router
	{
		actionController := &Controller{Admin: res.GetAdmin(), action: action}
		primaryKeyParams := res.ParamIDName()

		// Bulk actions
		res.RegisterRoute("GET", path.Join("!action", action.ToParam()), actionController.Action, &RouteConfig{Permissioner: action, PermissionMode: roles.Update})
		res.RegisterRoute("PUT", path.Join("!action", action.ToParam()), actionController.Action, &RouteConfig{Permissioner: action, PermissionMode: roles.Update})

		// Resource action
		res.RegisterRoute("GET", path.Join(primaryKeyParams, action.ToParam()), actionController.Action, &RouteConfig{Permissioner: action, PermissionMode: roles.Update})
		res.RegisterRoute("PUT", path.Join(primaryKeyParams, action.ToParam()), actionController.Action, &RouteConfig{Permissioner: action, PermissionMode: roles.Update})
	}

	return action
}

// GetAction get defined action
func (res *Resource) GetAction(name string) *Action {
	for _, action := range res.Actions {
		if action.Name == name {
			return action
		}
	}
	return nil
}

// ActionArgument action argument that used in handle
type ActionArgument struct {
	PrimaryValues       []string
	Context             *Context
	Argument            interface{}
	SkipDefaultResponse bool
}

// Action action definiation
type Action struct {
	Name        string
	Label       string
	Method      string
	URL         func(record interface{}, context *Context) string
	URLOpenType string
	Visible     func(record interface{}, context *Context) bool
	Handler     func(argument *ActionArgument) error
	Modes       []string
	Resource    *Resource
	Permission  *roles.Permission
}

// ToParam used to register routes for actions
func (action Action) ToParam() string {
	return utils.ToParamString(action.Name)
}

// IsAllowed check if current user has permission to view the action
func (action Action) IsAllowed(mode roles.PermissionMode, context *Context, records ...interface{}) bool {
	if action.Visible != nil {
		for _, record := range records {
			if !action.Visible(record, context) {
				return false
			}
		}
	}

	if action.Permission != nil {
		return action.HasPermission(mode, context.Context)
	}

	if context.Resource != nil {
		return context.Resource.HasPermission(mode, context.Context)
	}
	return true
}

// HasPermission check if current user has permission for the action
func (action Action) HasPermission(mode roles.PermissionMode, context *qor.Context) bool {
	if action.Permission != nil {
		return action.Permission.HasPermission(mode, context.Roles...)
	}

	return true
}

// FindSelectedRecords find selected records when run bulk actions
func (actionArgument *ActionArgument) FindSelectedRecords() []interface{} {
	var (
		context   = actionArgument.Context
		resource  = context.Resource
		records   = []interface{}{}
		sqls      []string
		sqlParams []interface{}
	)

	clone := context.clone()
	for _, primaryValue := range actionArgument.PrimaryValues {
		primaryQuerySQL, primaryParams := resource.ToPrimaryQueryParams(primaryValue, context.Context)
		sqls = append(sqls, primaryQuerySQL)
		sqlParams = append(sqlParams, primaryParams...)
	}

	if len(sqls) > 0 {
		clone.SetDB(clone.GetDB().Where(strings.Join(sqls, " OR "), sqlParams...))
	}
	results, _ := clone.FindMany()

	resultValues := reflect.Indirect(reflect.ValueOf(results))
	for i := 0; i < resultValues.Len(); i++ {
		records = append(records, resultValues.Index(i).Interface())
	}
	return records
}
