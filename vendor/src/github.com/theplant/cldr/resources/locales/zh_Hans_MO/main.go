package zh_Hans_MO

import "github.com/theplant/cldr"

var Locale = &cldr.Locale{
	Locale: "zh_Hans_MO",
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
