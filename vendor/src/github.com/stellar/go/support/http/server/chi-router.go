package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/stellar/go/support/log"
	"github.com/tylerb/graceful"
	"golang.org/x/net/http2"
)

// TLSConfig specifies the TLS portion of a config
type TLSConfig struct {
	CertificateFile string `toml:"certificate-file" valid:"required"`
	PrivateKeyFile  string `toml:"private-key-file" valid:"required"`
}

// routeMaker translates the HTTP string method to the chi-equivalent route operation
var routeMaker = map[string]func(*chi.Mux, string, http.HandlerFunc){
	http.MethodGet: func(mux *chi.Mux, route string, fn http.HandlerFunc) {
		mux.Get(route, fn)
	},
	http.MethodPost: func(mux *chi.Mux, route string, fn http.HandlerFunc) {
		mux.Post(route, fn)
	},
}

// NewRouter creates a new router with the provided config
func NewRouter(c *Config) *chi.Mux {
	mux := chi.NewRouter()

	// add middleware
	mux.Use(c.middlewares...)

	// add routes
	for method, routes := range c.router {
		bindFn := routeMaker[method]
		for _, route := range routes {
			bindFn(mux, route.path, route.handler)
		}
	}

	// not found handler
	if c.notFound != nil {
		mux.NotFound(c.notFound)
	}

	return mux
}

// Serve starts a web server by binding it to a socket and setting up the shutdown signals
func Serve(router *chi.Mux, port int, tls *TLSConfig) {
	http.Handle("/", router)

	addr := fmt.Sprintf(":%d", port)

	srv := &graceful.Server{
		Timeout: 10 * time.Second,

		Server: &http.Server{
			Addr:    addr,
			Handler: http.DefaultServeMux,
		},

		ShutdownInitiated: func() {
			log.Info("received signal, gracefully stopping")
		},
	}

	http2.ConfigureServer(srv.Server, nil)

	log.Info("Starting web service on " + addr)

	var err error
	if tls != nil {
		err = srv.ListenAndServeTLS(tls.CertificateFile, tls.PrivateKeyFile)
	} else {
		err = srv.ListenAndServe()
	}

	if err != nil {
		log.Panic(err)
	}

	log.Info("stopped")
}
