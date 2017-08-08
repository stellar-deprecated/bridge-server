package ln_AO

import "github.com/theplant/cldr"

var Locale = &cldr.Locale{
	Locale: "ln_AO",
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
