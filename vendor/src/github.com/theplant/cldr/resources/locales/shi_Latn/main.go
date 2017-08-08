package shi_Latn

import "github.com/theplant/cldr"

var Locale = &cldr.Locale{
	Locale: "shi_Latn",
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
