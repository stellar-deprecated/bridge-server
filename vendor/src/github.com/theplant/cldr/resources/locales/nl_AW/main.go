package nl_AW

import "github.com/theplant/cldr"

var Locale = &cldr.Locale{
	Locale: "nl_AW",
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
