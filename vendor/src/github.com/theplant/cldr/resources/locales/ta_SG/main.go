package ta_SG

import "github.com/theplant/cldr"

var Locale = &cldr.Locale{
	Locale: "ta_SG",
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
