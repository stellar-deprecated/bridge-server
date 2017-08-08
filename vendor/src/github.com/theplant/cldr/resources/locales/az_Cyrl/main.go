package az_Cyrl

import "github.com/theplant/cldr"

var Locale = &cldr.Locale{
	Locale: "az_Cyrl",
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
