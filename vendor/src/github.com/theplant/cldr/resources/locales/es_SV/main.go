package es_SV

import "github.com/theplant/cldr"

var Locale = &cldr.Locale{
	Locale: "es_SV",
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
