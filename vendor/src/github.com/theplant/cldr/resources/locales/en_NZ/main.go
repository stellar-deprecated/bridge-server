package en_NZ

import "github.com/theplant/cldr"

var Locale = &cldr.Locale{
	Locale: "en_NZ",
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
