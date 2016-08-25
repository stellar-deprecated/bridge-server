// Package config provides a common infrastructure for reading configuration
// data stored in local TOML files.
package config

import (
	"github.com/BurntSushi/toml"
	"github.com/asaskevich/govalidator"
	"github.com/stellar/go/strkey"
	"github.com/stellar/go/support/errors"
)

// InvalidConfigError is the error that is returned when an invalid
// configuration is encountered by the `Read` func.
type InvalidConfigError struct {
	InvalidFields map[string]string
}

// Read takes the TOML configuration file at `path`, parses it into `dest` and
// then uses github.com/asaskevich/govalidator to validate the struct.
func Read(path string, dest interface{}) error {
	_, err := toml.DecodeFile(path, dest)
	if err != nil {
		return errors.Wrap(err, "decode-file failed")
	}

	valid, err := govalidator.ValidateStruct(dest)

	if valid {
		return nil
	}

	fields := govalidator.ErrorsByField(err)

	return &InvalidConfigError{
		InvalidFields: fields,
	}
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
	govalidator.CustomTypeTagMap.Set("stellar_accountid", govalidator.CustomTypeValidator(isStellarAccountID))
	govalidator.CustomTypeTagMap.Set("stellar_seed", govalidator.CustomTypeValidator(isStellarSeed))

}

func isStellarAccountID(i interface{}, context interface{}) bool {
	enc, ok := i.(string)

	if !ok {
		return false
	}

	_, err := strkey.Decode(strkey.VersionByteAccountID, enc)

	if err == nil {
		return true
	}

	return false
}

func isStellarSeed(i interface{}, context interface{}) bool {
	enc, ok := i.(string)

	if !ok {
		return false
	}

	_, err := strkey.Decode(strkey.VersionByteSeed, enc)

	if err == nil {
		return true
	}

	return false
}
