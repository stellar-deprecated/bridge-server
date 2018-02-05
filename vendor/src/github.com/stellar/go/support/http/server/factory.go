package server

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
)

// EmptyConfig gives you a new empty Config
func EmptyConfig() *Config {
	return &Config{
		router:      make(map[string][]route),
		middlewares: []func(http.Handler) http.Handler{},
	}
}

// AddBasicMiddleware is a helper function that augments the passed in Config with some basic middleware components
func AddBasicMiddleware(c *Config) {
	c.Middleware(middleware.RequestID)
	c.Middleware(middleware.Recoverer)
	c.Middleware(BindLoggerMiddleware(middleware.RequestIDKey))
}
