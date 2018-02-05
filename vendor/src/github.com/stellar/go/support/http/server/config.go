package server

import (
	"net/http"
)

type route struct {
	path    string
	handler http.HandlerFunc
}

// Config is the immutable Config file that will be used to construct a server
type Config struct {
	router      map[string][]route // use a []route to maintain a consistent traversal ordering guarantee
	middlewares []func(http.Handler) http.Handler
	notFound    http.HandlerFunc
}

// Route allows you to set (or override) an existing route by the HTTP method
func (c *Config) Route(method string, path string, handler http.HandlerFunc) {
	c.router[method] = append(c.router[method], route{
		path:    path,
		handler: handler,
	})
}

// Middleware adds a middleware to the list
func (c *Config) Middleware(m func(http.Handler) http.Handler) {
	c.middlewares = append(c.middlewares, m)
}

// NotFound sets the handler for when a resource is not found
func (c *Config) NotFound(handler http.HandlerFunc) {
	c.notFound = handler
}
