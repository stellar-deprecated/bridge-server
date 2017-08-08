package en_MY

import "github.com/theplant/cldr"

var Locale = &cldr.Locale{
	Locale: "en_MY",
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
