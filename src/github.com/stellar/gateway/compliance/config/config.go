package config

import (
	"errors"
	"net/url"

	"github.com/stellar/go-stellar-base/keypair"
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
	Callbacks
}

type Keys struct {
	SigningSeed   string `mapstructure:"signing_seed"`
	EncryptionKey string `mapstructure:"encryption_key"`
}

type Callbacks struct {
	Sanctions string
	AskUser   string `mapstructure:"ask_user"`
	FetchInfo string `mapstructure:"fetch_info"`
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

	if c.Keys.SigningSeed == "" || c.Keys.EncryptionKey == "" {
		err = errors.New("keys.signing_seed and keys.encryption_key params are required")
		return
	}

	if c.Keys.SigningSeed != "" {
		_, err = keypair.Parse(c.Keys.SigningSeed)
		if err != nil {
			err = errors.New("keys.signing_seed is invalid")
			return
		}
	}

	if c.Keys.EncryptionKey != "" {
		_, err = keypair.Parse(c.Keys.EncryptionKey)
		if err != nil {
			err = errors.New("keys.encryption_key is invalid")
			return
		}
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

	if c.Callbacks.Sanctions != "" {
		_, err = url.Parse(c.Callbacks.Sanctions)
		if err != nil {
			err = errors.New("Cannot parse callbacks.sanctions param")
			return
		}
	}

	if c.Callbacks.Sanctions != "" {
		_, err = url.Parse(c.Callbacks.Sanctions)
		if err != nil {
			err = errors.New("Cannot parse callbacks.sanctions param")
			return
		}
	}

	if c.Callbacks.AskUser != "" {
		_, err = url.Parse(c.Callbacks.AskUser)
		if err != nil {
			err = errors.New("Cannot parse callbacks.ask_user param")
			return
		}
	}

	if c.Callbacks.FetchInfo != "" {
		_, err = url.Parse(c.Callbacks.FetchInfo)
		if err != nil {
			err = errors.New("Cannot parse callbacks.fetch_info param")
			return
		}
	}

	return
}
