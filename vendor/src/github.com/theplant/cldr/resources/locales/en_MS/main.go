package en_MS

import "github.com/theplant/cldr"

var Locale = &cldr.Locale{
	Locale: "en_MS",
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
