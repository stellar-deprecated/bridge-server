package zh_Hant

import "github.com/theplant/cldr"

var Locale = &cldr.Locale{
	Locale: "zh_Hant",
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
