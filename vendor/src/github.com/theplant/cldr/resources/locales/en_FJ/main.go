package en_FJ

import "github.com/theplant/cldr"

var Locale = &cldr.Locale{
	Locale: "en_FJ",
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
