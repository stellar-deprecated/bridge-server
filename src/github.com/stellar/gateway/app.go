package gateway

import (
	"errors"
	"flag"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"time"

	"github.com/stellar/gateway/db"
	"github.com/stellar/gateway/horizon"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web/middleware"
)

type App struct {
	config               Config
	entityManager        db.EntityManagerInterface
	horizon              horizon.HorizonInterface
	transactionSubmitter *TransactionSubmitter
	repository           db.RepositoryInterface
}

// NewApp constructs an new App instance from the provided config.
func NewApp(config Config) (app *App, err error) {
	entityManager, err := db.NewEntityManager(config.Database.Type, config.Database.Url)
	if err != nil {
		return
	}
	repository, err := db.NewRepository(config.Database.Type, config.Database.Url)
	if err != nil {
		return
	}

	h := horizon.New(config.Horizon)

	log.Print("Creating and initializing TransactionSubmitter")
	ts := NewTransactionSubmitter(&h, &entityManager)
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

	log.Print("Creating and starting PaymentListener")
	paymentListener, err := NewPaymentListener(&config, &entityManager, &h, &repository, time.Now)
	if err != nil {
		return
	}
	err = paymentListener.Listen()
	if err != nil {
		return
	}

	log.Print("PaymentListener created")

	if len(config.ApiKey) > 0 && len(config.ApiKey) < 15 {
		err = errors.New("api-key have to be at least 15 chars long.")
		return
	}

	app = &App{
		config:               config,
		entityManager:        &entityManager,
		horizon:              &h,
		repository:           &repository,
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
