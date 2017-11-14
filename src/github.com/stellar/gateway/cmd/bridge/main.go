package main

import (
	log "github.com/sirupsen/logrus"
	"runtime"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/stellar/gateway/bridge"
	"github.com/stellar/gateway/bridge/config"
)

var app *bridge.App
var rootCmd *cobra.Command
var migrateFlag bool
var configFile string
var versionFlag bool
var version = "N/A"

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	rootCmd.Execute()
}

func init() {
	rootCmd = &cobra.Command{
		Use:   "bridge",
		Short: "stellar bridge server",
		Long:  `stellar bridge server`,
		Run:   run,
	}

	rootCmd.Flags().BoolVarP(&migrateFlag, "migrate-db", "", false, "migrate DB to the newest schema version")
	rootCmd.Flags().StringVarP(&configFile, "config", "c", "bridge.cfg", "path to config file")
	rootCmd.Flags().BoolVarP(&versionFlag, "version", "v", false, "displays bridge server version")
}

func run(cmd *cobra.Command, args []string) {
	viper.SetConfigFile(configFile)
	viper.SetConfigType("toml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Error reading "+configFile+" file: ", err)
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

	app, err = bridge.NewApp(config, migrateFlag, versionFlag, version)

	if err != nil {
		log.Fatal(err.Error())
		return
	}

	app.Serve()
}
