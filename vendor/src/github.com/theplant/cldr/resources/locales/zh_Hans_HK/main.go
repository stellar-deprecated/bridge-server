package zh_Hans_HK

import "github.com/theplant/cldr"

var Locale = &cldr.Locale{
	Locale: "zh_Hans_HK",
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
