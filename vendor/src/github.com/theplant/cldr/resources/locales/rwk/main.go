package rwk

import "github.com/theplant/cldr"

var Locale = &cldr.Locale{
	Locale: "rwk",
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
