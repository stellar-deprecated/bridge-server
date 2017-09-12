package server

import (
	"net/http"
	"strings"
)

// StripTrailingSlashMiddleware strips trailing slash.
// Credit goes to https://github.com/stellar/horizon
func StripTrailingSlashMiddleware() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			path := r.URL.Path

			// Do not change /admin paths
			if strings.HasPrefix(path, "/admin") {
				next.ServeHTTP(w, r)
				return
			}

			l := len(path)

			// if the path is longer than 1 char (i.e., not '/')
			// and has a trailing slash, remove it.
			if l > 1 && path[l-1] == '/' {
				r.URL.Path = path[0 : l-1]
			}

			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}

// HeadersMiddleware sends headers required by servers
func HeadersMiddleware() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			// Do not change admin home
			if r.URL.Path == "/admin/" {
				next.ServeHTTP(w, r)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}

// APIKeyMiddleware checks for apiKey in a request and writes http.StatusForbidden if it's incorrect.
func APIKeyMiddleware(apiKey string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			k := r.PostFormValue("apiKey")
			if k != apiKey {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}
