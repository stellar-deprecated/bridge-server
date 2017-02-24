package compliance

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"net/http"
	"os"

	"github.com/facebookgo/inject"
	"github.com/stellar/gateway/compliance/config"
	"github.com/stellar/gateway/compliance/handlers"
	"github.com/stellar/gateway/crypto"
	"github.com/stellar/gateway/db"
	"github.com/stellar/gateway/db/drivers/mysql"
	"github.com/stellar/gateway/db/drivers/postgres"
	"github.com/stellar/gateway/server"
	"github.com/stellar/go/clients/federation"
	"github.com/stellar/go/clients/stellartoml"
	"github.com/zenazn/goji/graceful"
	"github.com/zenazn/goji/web"
)

// App is the application object
type App struct {
	config         config.Config
	requestHandler handlers.RequestHandler
}

// NewApp constructs an new App instance from the provided config.
func NewApp(config config.Config, migrateFlag bool) (app *App, err error) {
	var g inject.Graph

	var driver db.Driver
	switch config.Database.Type {
	case "mysql":
		driver = &mysql.Driver{}
	case "postgres":
		driver = &postgres.Driver{}
	default:
		return nil, fmt.Errorf("%s database has no driver", config.Database.Type)
	}

	err = driver.Init(config.Database.URL)
	if err != nil {
		return
	}

	entityManager := db.NewEntityManager(driver)
	repository := db.NewRepository(driver)

	if migrateFlag {
		var migrationsApplied int
		migrationsApplied, err = driver.MigrateUp("compliance")
		if err != nil {
			return
		}

		log.Info("Applied migrations: ", migrationsApplied)
		os.Exit(0)
		return
	}

	requestHandler := handlers.RequestHandler{}

	federationClient := &federation.Client{
		HTTP:        http.DefaultClient,
		StellarTOML: stellartoml.DefaultClient,
	}

	err = g.Provide(
		&inject.Object{Value: &requestHandler},
		&inject.Object{Value: &config},
		&inject.Object{Value: &entityManager},
		&inject.Object{Value: &repository},
		&inject.Object{Value: &crypto.SignerVerifier{}},
		&inject.Object{Value: stellartoml.DefaultClient},
		&inject.Object{Value: federationClient},
		&inject.Object{Value: &http.Client{}},
		&inject.Object{Value: &handlers.NonceGenerator{}},
	)

	if err != nil {
		log.Fatal("Injector: ", err)
	}

	if err := g.Populate(); err != nil {
		log.Fatal("Injector: ", err)
	}

	app = &App{
		config:         config,
		requestHandler: requestHandler,
	}
	return
}

// Serve starts the server
func (a *App) Serve() {
	// External endpoints
	external := web.New()
	external.Use(server.StripTrailingSlashMiddleware())
	external.Use(server.HeadersMiddleware())
	external.Post("/", a.requestHandler.HandlerAuth)
	externalPortString := fmt.Sprintf(":%d", *a.config.ExternalPort)
	log.Println("Starting external server on", externalPortString)
	go func() {
		var err error
		if a.config.TLS.CertificateFile != "" && a.config.TLS.PrivateKeyFile != "" {
			err = graceful.ListenAndServeTLS(
				externalPortString,
				a.config.TLS.CertificateFile,
				a.config.TLS.PrivateKeyFile,
				external,
			)
		} else {
			err = graceful.ListenAndServe(externalPortString, external)
		}

		if err != nil {
			log.Fatal(err)
		}
	}()

	// Internal endpoints
	internal := web.New()
	internal.Use(server.StripTrailingSlashMiddleware())
	internal.Use(server.HeadersMiddleware())
	internal.Post("/send", a.requestHandler.HandlerSend)
	internal.Post("/receive", a.requestHandler.HandlerReceive)
	internal.Post("/allow_access", a.requestHandler.HandlerAllowAccess)
	internal.Post("/remove_access", a.requestHandler.HandlerRemoveAccess)
	internalPortString := fmt.Sprintf(":%d", *a.config.InternalPort)
	log.Println("Starting internal server on", internalPortString)
	err := graceful.ListenAndServe(internalPortString, internal)
	if err != nil {
		log.Fatal(err)
	}
}
