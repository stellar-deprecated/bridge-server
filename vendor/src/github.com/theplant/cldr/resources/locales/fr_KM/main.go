package fr_KM

import "github.com/theplant/cldr"

var Locale = &cldr.Locale{
	Locale: "fr_KM",
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
