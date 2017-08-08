package es_AR

import "github.com/theplant/cldr"

var Locale = &cldr.Locale{
	Locale: "es_AR",
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
