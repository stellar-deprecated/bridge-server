package ar_SS

import "github.com/theplant/cldr"

var Locale = &cldr.Locale{
	Locale: "ar_SS",
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
