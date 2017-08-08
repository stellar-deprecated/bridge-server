package en_SZ

import "github.com/theplant/cldr"

var Locale = &cldr.Locale{
	Locale: "en_SZ",
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
