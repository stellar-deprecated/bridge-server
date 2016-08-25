package config

import (
	"testing"

	"github.com/asaskevich/govalidator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccountIDValidator(t *testing.T) {
	var val struct {
		Empty        string `valid:"stellar_accountid"`
		NotSTRKey    string `valid:"stellar_accountid"`
		NotAccountID string `valid:"stellar_accountid"`
		Valid        string `valid:"stellar_accountid"`
		WrongType    int    `valid:"stellar_accountid"`
	}

	val.NotSTRKey = "hello"
	val.NotAccountID = "SA5MATAU4RNJDKCTIC6VVSYSGB7MFFBVU3OKWOA5K67S62EYB5ESKLTV"
	val.Valid = "GBXS6WTZNRS7LOGHM3SCMAJD6M6JCXB3GATXECCZ3C5NJ3PVSZ23PEWX"
	val.WrongType = 100

	// run the validation
	ok, err := govalidator.ValidateStruct(val)
	require.False(t, ok)
	require.Error(t, err)

	fields := govalidator.ErrorsByField(err)

	// ensure valid is not in the invalid map
	_, ok = fields["Valid"]
	assert.False(t, ok)

	_, ok = fields["Empty"]
	assert.True(t, ok, "Empty is not an invalid field")

	_, ok = fields["NotSTRKey"]
	assert.True(t, ok, "NotSTRKey is not an invalid field")

	_, ok = fields["NotAccountID"]
	assert.True(t, ok, "NotAccountID is not an invalid field")

	_, ok = fields["WrongType"]
	assert.True(t, ok, "WrongType is not an invalid field")
}

func TestSeedValidator(t *testing.T) {
	var val struct {
		Empty     string `valid:"stellar_seed"`
		NotSTRKey string `valid:"stellar_seed"`
		NotSeed   string `valid:"stellar_seed"`
		Valid     string `valid:"stellar_seed"`
		WrongType int    `valid:"stellar_seed"`
	}

	val.NotSTRKey = "hello"
	val.NotSeed = "GBXS6WTZNRS7LOGHM3SCMAJD6M6JCXB3GATXECCZ3C5NJ3PVSZ23PEWX"
	val.Valid = "SA5MATAU4RNJDKCTIC6VVSYSGB7MFFBVU3OKWOA5K67S62EYB5ESKLTV"
	val.WrongType = 100

	// run the validation
	ok, err := govalidator.ValidateStruct(val)
	require.False(t, ok)
	require.Error(t, err)

	fields := govalidator.ErrorsByField(err)

	// ensure valid is not in the invalid map
	_, ok = fields["Valid"]
	assert.False(t, ok)

	_, ok = fields["Empty"]
	assert.True(t, ok, "Empty is not an invalid field")

	_, ok = fields["NotSTRKey"]
	assert.True(t, ok, "NotSTRKey is not an invalid field")

	_, ok = fields["NotSeed"]
	assert.True(t, ok, "NotSeed is not an invalid field")

	_, ok = fields["WrongType"]
	assert.True(t, ok, "WrongType is not an invalid field")
}
