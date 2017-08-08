package ru_KG

import "github.com/theplant/cldr"

var Locale = &cldr.Locale{
	Locale: "ru_KG",
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
