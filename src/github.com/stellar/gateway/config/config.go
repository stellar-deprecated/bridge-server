package config

import (
	"errors"
	"net/url"
)

type Config struct {
	Port              int
	Horizon           string
	ApiKey            string `mapstructure:"api_key"`
	NetworkPassphrase string `mapstructure:"network_passphrase"`
	Assets            []string
	Database          struct {
		Type string
		Url  string
	}
	Accounts
	Hooks
}

type Accounts struct {
	AuthorizingSeed    string `mapstructure:"authorizing_seed"`
	IssuingSeed        string `mapstructure:"issuing_seed"`
	ReceivingAccountId string `mapstructure:"receiving_account_id"`
}

type Hooks struct {
	Receive string
	Error   string
}

func (c *Config) Validate() (err error) {
	// Add `parseTime=true` param to mysql url
	if c.Database.Type == "mysql" {
		var mysqlUrl *url.URL
		mysqlUrl, err = url.Parse(c.Database.Url)
		if err != nil {
			err = errors.New("Cannot parse database.url parameter")
			return
		}
		query := mysqlUrl.Query()
		query.Set("parseTime", "true")
		mysqlUrl.RawQuery = query.Encode()
		c.Database.Url = mysqlUrl.String()
	}

	return
}
