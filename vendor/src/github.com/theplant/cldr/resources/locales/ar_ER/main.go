package ar_ER

import "github.com/theplant/cldr"

var Locale = &cldr.Locale{
	Locale: "ar_ER",
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
