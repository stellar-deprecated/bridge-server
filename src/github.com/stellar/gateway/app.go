package gateway

import (
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
func NewApp(config Config) (*App, error) {
	database, err := sqlx.Connect(
		config.Database.Type,
		config.Database.Url,
	)

	if err != nil {
		log.Panic(err)
	}

	h := horizon.Horizon{config.Horizon}

	log.Print("Creating TransactionSubmitter")
	ts, err := NewTransactionSubmitter(&h, config.Accounts.ChannelsSeeds)
	if err != nil {
		log.Panic(err)
	}
	log.Print("TransactionSubmitter created")

	app := &App{
		config: config,
		database: database,
		horizon: &h,
		transactionSubmitter: &ts,
	}
	return app, nil
}

func (a *App) Serve() {
	requestHandlers := &RequestHandler{
		config: &a.config,
		database: a.database,
		transactionSubmitter: a.transactionSubmitter,
	}

	portString := fmt.Sprintf(":%d", a.config.Port)
	flag.Set("bind", portString)

	goji.Use(stripTrailingSlashMiddleware())
	goji.Use(headersMiddleware())

	goji.Get("/authorize", requestHandlers.Authorize)
	goji.Get("/send", requestHandlers.Send)
	goji.Serve()
}
