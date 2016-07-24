package config

import (
	"errors"
	"net/url"

	"github.com/stellar/go-stellar-base/keypair"
)

// Config contains config params of the bridge server
type Config struct {
	Port              *int
	Horizon           string
	Compliance        string
	LogFormat         string `mapstructure:"log_format"`
	APIKey            string `mapstructure:"api_key"`
	NetworkPassphrase string `mapstructure:"network_passphrase"`
	Assets            []Asset
	Database          struct {
		Type string
		URL  string
	}
	Accounts
	Callbacks
}

// Asset represents credit asset
type Asset struct {
	Code   string
	Issuer string
}

// Accounts contains values of `accounts` config group
type Accounts struct {
	AuthorizingSeed    string `mapstructure:"authorizing_seed"`
	BaseSeed           string `mapstructure:"base_seed"`
	IssuingAccountID   string `mapstructure:"issuing_account_id"`
	ReceivingAccountID string `mapstructure:"receiving_account_id"`
}

// Callbacks contains values of `callbacks` config group
type Callbacks struct {
	Receive string
	Error   string
}

// Validate validates config and returns error if any of config values is incorrect
func (c *Config) Validate() (err error) {
	if c.Port == nil {
		err = errors.New("port param is required")
		return
	}

	if c.Horizon == "" {
		err = errors.New("horizon param is required")
		return
	}

	_, err = url.Parse(c.Horizon)
	if err != nil {
		err = errors.New("Cannot parse horizon param")
		return
	}

	if c.NetworkPassphrase == "" {
		err = errors.New("network_passphrase param is required")
		return
	}

	var dbURL *url.URL
	dbURL, err = url.Parse(c.Database.URL)
	if err != nil {
		err = errors.New("Cannot parse database.url param")
		return
	}

	switch c.Database.Type {
	case "mssql":
		break
	case "mysql":
		// Add `parseTime=true` param to mysql url
		query := dbURL.Query()
		query.Set("parseTime", "true")
		dbURL.RawQuery = query.Encode()
		c.Database.URL = dbURL.String()
	case "postgres":
		break
	case "":
		// Allow to start gateway server with a single endpoint: /payment
		break
	default:
		err = errors.New("Invalid database.type param")
		return
	}

	if c.Accounts.AuthorizingSeed != "" {
		_, err = keypair.Parse(c.Accounts.AuthorizingSeed)
		if err != nil {
			err = errors.New("accounts.authorizing_seed is invalid")
			return
		}
	}

	if c.Accounts.BaseSeed != "" {
		_, err = keypair.Parse(c.Accounts.BaseSeed)
		if err != nil {
			err = errors.New("accounts.base_seed is invalid")
			return
		}
	}

	if c.Accounts.IssuingAccountID != "" {
		_, err = keypair.Parse(c.Accounts.IssuingAccountID)
		if err != nil {
			err = errors.New("accounts.issuing_account_id is invalid")
			return
		}
	}

	if c.Accounts.ReceivingAccountID != "" {
		_, err = keypair.Parse(c.Accounts.ReceivingAccountID)
		if err != nil {
			err = errors.New("accounts.receiving_account_id is invalid")
			return
		}
	}

	if c.Callbacks.Receive != "" {
		_, err = url.Parse(c.Callbacks.Receive)
		if err != nil {
			err = errors.New("Cannot parse callbacks.receive param")
			return
		}
	}

	if c.Callbacks.Error != "" {
		_, err = url.Parse(c.Callbacks.Error)
		if err != nil {
			err = errors.New("Cannot parse callbacks.error param")
			return
		}
	}

	return
}
