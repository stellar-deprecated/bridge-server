package ca_FR

import "github.com/theplant/cldr"

var Locale = &cldr.Locale{
	Locale: "ca_FR",
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
