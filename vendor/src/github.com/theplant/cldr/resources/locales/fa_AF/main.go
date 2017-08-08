package fa_AF

import "github.com/theplant/cldr"

var Locale = &cldr.Locale{
	Locale: "fa_AF",
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
