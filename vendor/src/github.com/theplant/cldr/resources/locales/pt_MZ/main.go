package pt_MZ

import "github.com/theplant/cldr"

var Locale = &cldr.Locale{
	Locale: "pt_MZ",
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
