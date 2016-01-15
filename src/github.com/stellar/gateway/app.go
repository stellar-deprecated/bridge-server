package gateway

import (
	"errors"
	"flag"
	"fmt"
	log "github.com/Sirupsen/logrus"

	"github.com/stellar/gateway/db"
	"github.com/stellar/gateway/horizon"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web/middleware"
)

type App struct {
	config               Config
	entityManager        *db.EntityManager
	horizon              *horizon.Horizon
	transactionSubmitter *TransactionSubmitter
}

// NewApp constructs an new App instance from the provided config.
func NewApp(config Config) (app *App, err error) {
	em, err := db.NewEntityManager(config.Database.Type, config.Database.Url)

	if err != nil {
		return
	}

	h := horizon.New(config.Horizon)

	log.Print("Creating and initializing TransactionSubmitter")
	ts := NewTransactionSubmitter(&h, &em)
	if err != nil {
		return
	}

	log.Print("Initializing Authorizing account")
	err = ts.InitAccount(config.Accounts.AuthorizingSeed)
	if err != nil {
		return
	}

	log.Print("Initializing Issuing account")
	err = ts.InitAccount(config.Accounts.IssuingSeed)
	if err != nil {
		return
	}

	log.Print("TransactionSubmitter created")

	if len(config.ApiKey) > 0 && len(config.ApiKey) < 15 {
		err = errors.New("api-key have to be at least 15 chars long.")
		return
	}

	app = &App{
		config:               config,
		entityManager:        &em,
		horizon:              &h,
		transactionSubmitter: &ts,
	}
	return
}

func (a *App) Serve() {
	requestHandlers := &RequestHandler{
		config:               &a.config,
		transactionSubmitter: a.transactionSubmitter,
	}

	portString := fmt.Sprintf(":%d", a.config.Port)
	flag.Set("bind", portString)

	goji.Abandon(middleware.Logger)
	goji.Use(stripTrailingSlashMiddleware())
	goji.Use(headersMiddleware())
	if a.config.ApiKey != "" {
		goji.Use(apiKeyMiddleware(a.config.ApiKey))
	}

	goji.Get("/authorize", requestHandlers.Authorize)
	goji.Get("/send", requestHandlers.Send)
	goji.Serve()
}
