package ar_LY

import "github.com/theplant/cldr"

var Locale = &cldr.Locale{
	Locale: "ar_LY",
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
