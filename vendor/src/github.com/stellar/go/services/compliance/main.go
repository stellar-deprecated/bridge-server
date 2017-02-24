package main

import (
	"fmt"
	"os"

	"goji.io"
	"goji.io/pat"

	"github.com/rs/cors"
	"github.com/spf13/cobra"
	complianceHandler "github.com/stellar/go/handlers/compliance"
	complianceProtocol "github.com/stellar/go/protocols/compliance"
	"github.com/stellar/go/support/app"
	"github.com/stellar/go/support/config"
	"github.com/stellar/go/support/errors"
	"github.com/stellar/go/support/http"
	"github.com/stellar/go/support/log"
)

// Config represents the configuration of a federation server
type Config struct {
	ExternalPort      int    `valid:"required" toml:"external_port"`
	InternalPort      int    `valid:"required" toml:"internal_port"`
	NeedsAuth         bool   `valid:"required" toml:"needs_auth"`
	NetworkPassphrase string `valid:"required" toml:"network_passphrase"`
	Keys              struct {
		SigningSeed string `valid:"stellar_seed,required" toml:"signing_seed"`
	} `valid:"required"`
	Callbacks struct {
		Sanctions   string `valid:"url,optional" toml:"sanctions"`
		AskUser     string `valid:"url,optional" toml:"ask_user"`
		GetUserData string `valid:"url,optional" toml:"get_user_data"`
	} `valid:"optional"`
	TLS struct {
		CertificateFile string `valid:"required" toml:"certificate_file"`
		PrivateKeyFile  string `valid:"required" toml:"private_key_file"`
	} `valid:"required"`
}

func main() {
	rootCmd := &cobra.Command{
		Use:   "compliance",
		Short: "stellar compliance server",
		Long:  "",
		Run:   run,
	}

	rootCmd.PersistentFlags().String("conf", "./compliance.cfg", "config file path")
	rootCmd.Execute()
}

func run(cmd *cobra.Command, args []string) {
	var (
		cfg     Config
		cfgPath = cmd.PersistentFlags().Lookup("conf").Value.String()
	)
	err := config.Read(cfgPath, &cfg)
	log.SetLevel(log.InfoLevel)

	if err != nil {
		switch cause := errors.Cause(err).(type) {
		case *config.InvalidConfigError:
			log.Error("config file: ", cause)
		default:
			log.Error(err)
		}
		os.Exit(1)
	}

	strategy := &complianceHandler.CallbackStrategy{
		SanctionsCheckURL: cfg.Callbacks.Sanctions,
		GetUserDataURL:    cfg.Callbacks.GetUserData,
	}

	mux := initMux(strategy)
	addr := fmt.Sprintf("0.0.0.0:%d", cfg.ExternalPort)

	http.Run(http.Config{
		ListenAddr: addr,
		Handler:    mux,
		TLSCert:    cfg.TLS.CertificateFile,
		TLSKey:     cfg.TLS.PrivateKeyFile,
		OnStarting: func() {
			log.Infof("starting compliance server - %s", app.Version())
			log.Infof("listening on %s", addr)
		},
	})
}

func initMux(strategy complianceHandler.Strategy) *goji.Mux {
	mux := goji.NewMux()

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"*"},
		AllowedMethods: []string{"*"},
	})
	mux.Use(c.Handler)
	mux.Use(log.HTTPMiddleware)

	authHandler := &complianceHandler.AuthHandler{
		Strategy: strategy,
		PersistTransaction: func(data complianceProtocol.AuthData) error {
			fmt.Println("Persist")
			return nil
		},
	}

	mux.Handle(pat.Post("/auth"), authHandler)
	mux.Handle(pat.Post("/auth/"), authHandler)

	return mux
}
