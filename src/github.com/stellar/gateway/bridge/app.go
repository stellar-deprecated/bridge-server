package bridge

import (
	"errors"
	"flag"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"time"

	"github.com/elazarl/go-bindata-assetfs"
	"github.com/facebookgo/inject"
	"github.com/stellar/gateway/bridge/config"
	"github.com/stellar/gateway/bridge/gui"
	"github.com/stellar/gateway/bridge/handlers"
	"github.com/stellar/gateway/db"
	"github.com/stellar/gateway/db/drivers/mysql"
	"github.com/stellar/gateway/db/drivers/postgres"
	"github.com/stellar/gateway/horizon"
	"github.com/stellar/gateway/listener"
	"github.com/stellar/gateway/server"
	"github.com/stellar/gateway/submitter"
	"github.com/stellar/go/clients/federation"
	"github.com/stellar/go/clients/stellartoml"
	"github.com/zenazn/goji/graceful"
	"github.com/zenazn/goji/web"
	"github.com/zenazn/goji/web/middleware"
)

// App is the application object
type App struct {
	config         config.Config
	requestHandler handlers.RequestHandler
}

// NewApp constructs an new App instance from the provided config.
func NewApp(config config.Config, migrateFlag bool, versionFlag bool, version string) (app *App, err error) {
	var g inject.Graph

	var driver db.Driver
	switch config.Database.Type {
	case "mysql":
		driver = &mysql.Driver{}
	case "postgres":
		driver = &postgres.Driver{}
	case "":
		// Allow to start gateway server with a single endpoint: /payment
		break
	default:
		return nil, fmt.Errorf("%s database has no driver", config.Database.Type)
	}

	var entityManager db.EntityManagerInterface
	var repository db.Repository

	if driver != nil {
		err = driver.Init(config.Database.URL)
		if err != nil {
			err = fmt.Errorf("Cannot connect to a DB: %s", err)
			return
		}

		entityManager = db.NewEntityManager(driver)
		repository = db.NewRepository(driver)
	}

	if migrateFlag {
		if driver == nil {
			log.Fatal("No database driver.")
			return
		}

		var migrationsApplied int
		migrationsApplied, err = driver.MigrateUp("gateway")
		if err != nil {
			return
		}

		log.Info("Applied migrations: ", migrationsApplied)
		os.Exit(0)
		return
	}

	if versionFlag {
		fmt.Printf("Bridge Server Version: %s \n", version)
		os.Exit(0)
		return
	}

	h := horizon.New(config.Horizon)

	log.Print("Creating and initializing TransactionSubmitter")
	ts := submitter.NewTransactionSubmitter(&h, entityManager, config.NetworkPassphrase, time.Now)
	if err != nil {
		return
	}

	log.Print("Initializing Authorizing account")

	if config.Accounts.AuthorizingSeed == "" {
		log.Warning("No accounts.authorizing_seed param. Skipping...")
	} else {
		err = ts.InitAccount(config.Accounts.AuthorizingSeed)
		if err != nil {
			return
		}
	}

	if config.Accounts.BaseSeed == "" {
		log.Warning("No accounts.base_seed param. Skipping...")
	} else {
		log.Print("Initializing Base account")
		err = ts.InitAccount(config.Accounts.BaseSeed)
		if err != nil {
			return
		}
	}

	log.Print("TransactionSubmitter created")

	log.Print("Creating and starting PaymentListener")

	var paymentListener listener.PaymentListener

	if config.Accounts.ReceivingAccountID == "" {
		log.Warning("No accounts.receiving_account_id param. Skipping...")
	} else if config.Callbacks.Receive == "" {
		log.Warning("No callbacks.receive param. Skipping...")
	} else {
		paymentListener, err = listener.NewPaymentListener(&config, entityManager, &h, repository, time.Now)
		if err != nil {
			return
		}
		err = paymentListener.Listen()
		if err != nil {
			return
		}

		log.Print("PaymentListener created")
	}

	if len(config.APIKey) > 0 && len(config.APIKey) < 15 {
		err = errors.New("api-key have to be at least 15 chars long")
		return
	}

	requestHandler := handlers.RequestHandler{}

	httpClientWithTimeout := http.Client{
		Timeout: 10 * time.Second,
	}

	stellartomlClient := stellartoml.Client{
		HTTP: &httpClientWithTimeout,
	}

	federationClient := federation.Client{
		HTTP:        &httpClientWithTimeout,
		StellarTOML: &stellartomlClient,
	}

	err = g.Provide(
		&inject.Object{Value: &requestHandler},
		&inject.Object{Value: &config},
		&inject.Object{Value: &stellartomlClient},
		&inject.Object{Value: &federationClient},
		&inject.Object{Value: &h},
		&inject.Object{Value: &repository},
		&inject.Object{Value: driver},
		&inject.Object{Value: &ts},
		&inject.Object{Value: &paymentListener},
		&inject.Object{Value: &httpClientWithTimeout},
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
	portString := fmt.Sprintf(":%d", *a.config.Port)
	flag.Set("bind", portString)

	bridge := web.New()

	bridge.Abandon(middleware.Logger)
	bridge.Use(server.StripTrailingSlashMiddleware())
	bridge.Use(server.HeadersMiddleware())
	if a.config.APIKey != "" {
		bridge.Use(server.APIKeyMiddleware(a.config.APIKey))
	}

	if a.config.Accounts.AuthorizingSeed != "" {
		bridge.Post("/authorize", a.requestHandler.Authorize)
	} else {
		log.Warning("accounts.authorizing_seed not provided. /authorize endpoint will not be available.")
	}

	bridge.Post("/create-keypair", a.requestHandler.CreateKeypair)
	bridge.Post("/builder", a.requestHandler.Builder)
	bridge.Post("/payment", a.requestHandler.Payment)
	bridge.Get("/payment", a.requestHandler.Payment)
	bridge.Post("/reprocess", a.requestHandler.Reprocess)

	bridge.Get("/admin/received-payments", a.requestHandler.AdminReceivedPayments)
	bridge.Get("/admin/received-payments/:id", a.requestHandler.AdminReceivedPayment)
	bridge.Get("/admin/sent-transactions", a.requestHandler.AdminSentTransactions)

	if a.config.Develop {
		// Create a proxy server to localhost:3000 where GUI development server lives.
		staticAdminURL, err := url.Parse("http://localhost:3000")
		if err != nil {
			panic(err)
		}
		bridge.Get("/*", httputil.NewSingleHostReverseProxy(staticAdminURL))
	} else {
		// Load go-bindata files
		fileServerHandler := http.FileServer(
			&assetfs.AssetFS{
				Asset:     gui.Asset,
				AssetDir:  gui.AssetDir,
				AssetInfo: gui.AssetInfo,
			})
		bridge.Get("/admin", func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "/admin/", http.StatusPermanentRedirect)
		})
		bridge.Get("/admin/*", http.StripPrefix("/admin/", fileServerHandler))
	}

	err := graceful.ListenAndServe(portString, bridge)
	if err != nil {
		log.Fatal(err)
	}
}
