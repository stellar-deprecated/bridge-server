package en_TT

import "github.com/theplant/cldr"

var Locale = &cldr.Locale{
	Locale: "en_TT",
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
