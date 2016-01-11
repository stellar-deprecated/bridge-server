package gateway

import (
	"errors"
	"flag"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stellar/gateway/horizon"
	"github.com/zenazn/goji"
)

type Database interface {
	Get(dest interface{}, query string, args ...interface{}) error
}

type App struct {
	config               Config
	database             Database
	horizon              *horizon.Horizon
	transactionSubmitter *TransactionSubmitter
}

// NewApp constructs an new App instance from the provided config.
func NewApp(config Config) (app *App, err error) {
	database, err := sqlx.Connect(
		config.Database.Type,
		config.Database.Url,
	)

	if err != nil {
		return
	}

	h := horizon.Horizon{config.Horizon}

	log.Print("Creating TransactionSubmitter")
	ts, err := NewTransactionSubmitter(&h, config.Accounts.ChannelsSeeds)
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
		database:             database,
		horizon:              &h,
		transactionSubmitter: &ts,
	}
	return
}

func (a *App) Serve() {
	requestHandlers := &RequestHandler{
		config:               &a.config,
		database:             a.database,
		transactionSubmitter: a.transactionSubmitter,
	}

	portString := fmt.Sprintf(":%d", a.config.Port)
	flag.Set("bind", portString)

	goji.Use(stripTrailingSlashMiddleware())
	goji.Use(headersMiddleware())
	if a.config.ApiKey != "" {
		goji.Use(apiKeyMiddleware(a.config.ApiKey))
	}

	goji.Get("/authorize", requestHandlers.Authorize)
	goji.Get("/send", requestHandlers.Send)
	goji.Serve()
}
