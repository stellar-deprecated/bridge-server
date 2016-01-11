package main

import (
	"fmt"
	"log"
	"runtime"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/stellar/gateway"
)

var app *gateway.App
var rootCmd *cobra.Command

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	rootCmd.Execute()
}

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	rootCmd = &cobra.Command{
		Use:   "gateway",
		Short: "stellar gateway server",
		Long:  `stellar gateway server`,
		Run:   run,
	}
}

func run(cmd *cobra.Command, args []string) {
	log.Print("Reading config.toml file")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	var config gateway.Config
	err = viper.Unmarshal(&config)
	app, err = gateway.NewApp(config)

	if err != nil {
		log.Fatal(err.Error())
		return
	}

	app.Serve()
}
