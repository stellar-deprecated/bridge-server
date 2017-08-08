package de_LU

import "github.com/theplant/cldr"

var Locale = &cldr.Locale{
	Locale: "de_LU",
	Number: cldr.Number{
		Symbols:    symbols,
		Formats:    formats,
		Currencies: currencies,
	},
	Calendar:   calendar,
	PluralRule: pluralRule,
}

func init() {
	cldr.RegisterLocale(Locale)
}
