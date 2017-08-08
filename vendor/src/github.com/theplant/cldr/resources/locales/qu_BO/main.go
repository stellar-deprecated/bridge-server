package qu_BO

import "github.com/theplant/cldr"

var Locale = &cldr.Locale{
	Locale: "qu_BO",
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
