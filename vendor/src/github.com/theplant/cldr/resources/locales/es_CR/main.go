package es_CR

import "github.com/theplant/cldr"

var Locale = &cldr.Locale{
	Locale: "es_CR",
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
