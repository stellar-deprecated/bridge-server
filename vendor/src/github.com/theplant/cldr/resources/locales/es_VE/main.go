package es_VE

import "github.com/theplant/cldr"

var Locale = &cldr.Locale{
	Locale: "es_VE",
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
