package en_FK

import "github.com/theplant/cldr"

var Locale = &cldr.Locale{
	Locale: "en_FK",
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
