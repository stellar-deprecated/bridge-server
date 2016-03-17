package compliance

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"os"

	"github.com/stellar/gateway/compliance/config"
	"github.com/stellar/gateway/compliance/handlers"
	"github.com/stellar/gateway/db"
	"github.com/stellar/gateway/db/drivers/mysql"
	"github.com/stellar/gateway/db/drivers/postgres"
	gatewayHandlers "github.com/stellar/gateway/handlers"
	"github.com/zenazn/goji/graceful"
	"github.com/zenazn/goji/web"
)

type App struct {
	config        config.Config
	entityManager db.EntityManagerInterface
	repository    db.RepositoryInterface
}

// NewApp constructs an new App instance from the provided config.
func NewApp(config config.Config, migrateFlag bool) (app *App, err error) {
	var driver db.Driver
	switch config.Database.Type {
	case "mysql":
		driver = &mysql.MysqlDriver{}
	case "postgres":
		driver = &postgres.PostgresDriver{}
	default:
		return nil, fmt.Errorf("%s database has no driver.", config.Database.Type)
	}

	err = driver.Init(config.Database.Url)
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

	app = &App{
		config:        config,
		entityManager: entityManager,
		repository:    repository,
	}
	return
}

func (a *App) Serve() {
	requestHandlers := &handlers.RequestHandler{
		Config:        &a.config,
		EntityManager: a.entityManager,
		Repository:    a.repository,
	}

	// External endpoints
	external := web.New()
	external.Use(gatewayHandlers.StripTrailingSlashMiddleware())
	external.Use(gatewayHandlers.HeadersMiddleware())
	external.Post("/", requestHandlers.HandlerAuth)
	externalPortString := fmt.Sprintf(":%d", *a.config.ExternalPort)
	log.Println("Starting external server on", externalPortString)
	go graceful.ListenAndServe(externalPortString, external)

	// Internal endpoints
	internal := web.New()
	internal.Use(gatewayHandlers.StripTrailingSlashMiddleware())
	internal.Use(gatewayHandlers.HeadersMiddleware())
	internal.Get("/send_payment", requestHandlers.HandlerSendPayment)
	internal.Get("/receive_payment", requestHandlers.HandlerReceivePayment)
	internalPortString := fmt.Sprintf(":%d", *a.config.InternalPort)
	log.Println("Starting internal server on", internalPortString)
	graceful.ListenAndServe(internalPortString, internal)
}
