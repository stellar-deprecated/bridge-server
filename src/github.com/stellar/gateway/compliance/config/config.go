package config

import (
	"errors"
	"net/url"
)

type Config struct {
	ExternalPort      *int   `mapstructure:"external_port"`
	InternalPort      *int   `mapstructure:"internal_port"`
	NeedsAuth         bool   `mapstructure:"needs_auth"`
	NetworkPassphrase string `mapstructure:"network_passphrase"`
	Database          struct {
		Type string
		Url  string
	}
	Keys
	Callbacks *Callbacks
}

type Keys struct {
	SigningSeed string `mapstructure:"signing_seed"`
	Encryption  string `mapstructure:"encryption"`
}

type Callbacks struct {
	Sanctions *string
	AskUser   *string `mapstructure:"ask_user"`
	FetchInfo *string `mapstructure:"fetch_info"`
}

func (c *Config) Validate() (err error) {
	if c.ExternalPort == nil {
		err = errors.New("external_port param is required")
		return
	}

	if c.InternalPort == nil {
		err = errors.New("internal_port param is required")
		return
	}

	if c.NetworkPassphrase == "" {
		err = errors.New("network_passphrase param is required")
		return
	}

	if c.Keys.SigningSeed == "" || c.Keys.Encryption == "" {
		err = errors.New("keys.signing_seed and keys.encryption params are required")
		return
	}

	var dbUrl *url.URL
	dbUrl, err = url.Parse(c.Database.Url)
	if err != nil {
		err = errors.New("Cannot parse database.url param")
		return
	}

	switch c.Database.Type {
	case "mysql":
		// Add `parseTime=true` param to mysql url
		query := dbUrl.Query()
		query.Set("parseTime", "true")
		dbUrl.RawQuery = query.Encode()
		c.Database.Url = dbUrl.String()
	case "postgres":
		break
	default:
		err = errors.New("Invalid database.type param")
		return
	}

	return
}
