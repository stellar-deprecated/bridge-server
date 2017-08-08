package pt_PT

import "github.com/theplant/cldr"

var Locale = &cldr.Locale{
	Locale: "pt_PT",
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
