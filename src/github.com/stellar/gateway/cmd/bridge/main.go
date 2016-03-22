package main

import (
	log "github.com/Sirupsen/logrus"
	"runtime"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/stellar/gateway/bridge"
	"github.com/stellar/gateway/bridge/config"
)

var app *bridge.App
var rootCmd *cobra.Command
var migrateFlag bool

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	rootCmd.Execute()
}

func init() {
	viper.SetConfigName("config_bridge")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	rootCmd = &cobra.Command{
		Use:   "bridge",
		Short: "stellar bridge server",
		Long:  `stellar bridge server`,
		Run:   run,
	}

	rootCmd.Flags().BoolVarP(&migrateFlag, "migrate-db", "", false, "migrate DB to the newest schema version")
}

func run(cmd *cobra.Command, args []string) {
	log.Print("Reading config_bridge.toml file")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Error reading config file: ", err)
	}

	var config config.Config
	err = viper.Unmarshal(&config)

	err = config.Validate()
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	app, err = bridge.NewApp(config, migrateFlag)

	if err != nil {
		log.Fatal(err.Error())
		return
	}

	app.Serve()
}
