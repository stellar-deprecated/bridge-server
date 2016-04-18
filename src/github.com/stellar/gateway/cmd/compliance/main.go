package main

import (
	log "github.com/Sirupsen/logrus"
	"runtime"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/stellar/gateway/compliance"
	"github.com/stellar/gateway/compliance/config"
)

var app *compliance.App
var rootCmd *cobra.Command
var migrateFlag bool

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	rootCmd.Execute()
}

func init() {
	viper.SetConfigName("config_compliance")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	rootCmd = &cobra.Command{
		Use:   "compliance",
		Short: "stellar compliance server",
		Long:  `stellar compliance server`,
		Run:   run,
	}

	rootCmd.Flags().BoolVarP(&migrateFlag, "migrate-db", "", false, "migrate DB to the newest schema version")
}

func run(cmd *cobra.Command, args []string) {
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Error reading config_compliance.toml file: ", err)
	}

	var config config.Config
	err = viper.Unmarshal(&config)

	err = config.Validate()
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	if config.LogFormat == "json" {
		log.SetFormatter(&log.JSONFormatter{})
	}

	app, err = compliance.NewApp(config, migrateFlag)

	if err != nil {
		log.Fatal(err.Error())
		return
	}

	app.Serve()
}
