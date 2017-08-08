package en_US_POSIX

import "github.com/theplant/cldr"

var Locale = &cldr.Locale{
	Locale: "en_US_POSIX",
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
