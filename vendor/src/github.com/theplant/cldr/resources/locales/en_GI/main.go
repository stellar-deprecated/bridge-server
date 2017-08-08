package en_GI

import "github.com/theplant/cldr"

var Locale = &cldr.Locale{
	Locale: "en_GI",
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
