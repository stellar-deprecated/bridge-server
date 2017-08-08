package de_AT

import "github.com/theplant/cldr"

var Locale = &cldr.Locale{
	Locale: "de_AT",
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
