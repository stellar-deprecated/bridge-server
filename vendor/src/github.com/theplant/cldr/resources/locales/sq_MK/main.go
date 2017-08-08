package sq_MK

import "github.com/theplant/cldr"

var Locale = &cldr.Locale{
	Locale: "sq_MK",
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
