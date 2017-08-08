package en_AU

import "github.com/theplant/cldr"

var Locale = &cldr.Locale{
	Locale: "en_AU",
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
