package uz_Arab

import "github.com/theplant/cldr"

var Locale = &cldr.Locale{
	Locale: "uz_Arab",
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
