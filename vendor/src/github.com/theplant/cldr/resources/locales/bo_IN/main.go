package bo_IN

import "github.com/theplant/cldr"

var Locale = &cldr.Locale{
	Locale: "bo_IN",
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
