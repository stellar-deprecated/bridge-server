package ta_LK

import "github.com/theplant/cldr"

var Locale = &cldr.Locale{
	Locale: "ta_LK",
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
