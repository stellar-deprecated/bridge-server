package af_NA

import "github.com/theplant/cldr"

var Locale = &cldr.Locale{
	Locale: "af_NA",
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
