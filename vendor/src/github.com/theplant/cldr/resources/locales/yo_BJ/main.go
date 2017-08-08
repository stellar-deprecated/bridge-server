package yo_BJ

import "github.com/theplant/cldr"

var Locale = &cldr.Locale{
	Locale: "yo_BJ",
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
