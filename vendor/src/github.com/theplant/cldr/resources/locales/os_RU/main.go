package os_RU

import "github.com/theplant/cldr"

var Locale = &cldr.Locale{
	Locale: "os_RU",
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
