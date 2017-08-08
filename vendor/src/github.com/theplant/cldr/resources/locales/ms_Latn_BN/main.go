package ms_Latn_BN

import "github.com/theplant/cldr"

var Locale = &cldr.Locale{
	Locale: "ms_Latn_BN",
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
