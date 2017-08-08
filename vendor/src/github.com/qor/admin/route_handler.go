package admin

import (
	"strings"

	"github.com/qor/qor"
	"github.com/qor/roles"
)

var blankPermissionMode roles.PermissionMode

// RouteConfig config for admin routes
type RouteConfig struct {
	Resource       *Resource
	Permissioner   HasPermissioner
	PermissionMode roles.PermissionMode
	Values         map[interface{}]interface{}
}

type requestHandler func(c *Context)

type routeHandler struct {
	Path   string
	Handle requestHandler
	Config *RouteConfig
}

func newRouteHandler(path string, handle requestHandler, configs ...*RouteConfig) *routeHandler {
	handler := routeHandler{
		Path:   "/" + strings.TrimPrefix(path, "/"),
		Handle: handle,
	}

	for _, config := range configs {
		handler.Config = config
	}

	if handler.Config == nil {
		handler.Config = &RouteConfig{}
	}

	if handler.Config.Permissioner == nil && handler.Config.Resource != nil {
		handler.Config.Permissioner = handler.Config.Resource
	}
	return &handler
}

func (handler routeHandler) HasPermission(permissionMode roles.PermissionMode, context *qor.Context) bool {
	if handler.Config.Permissioner == nil {
		return true
	}

	if handler.Config.PermissionMode != "" {
		return handler.Config.Permissioner.HasPermission(handler.Config.PermissionMode, context)
	}

	return handler.Config.Permissioner.HasPermission(permissionMode, context)
}
