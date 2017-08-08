package ha_Latn_GH

import "github.com/theplant/cldr"

var Locale = &cldr.Locale{
	Locale: "ha_Latn_GH",
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
