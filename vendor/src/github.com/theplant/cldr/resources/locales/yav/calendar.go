package yav

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "d/M/y"},
		Time:     cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "o.1", Feb: "o.2", Mar: "o.3", Apr: "o.4", May: "o.5", Jun: "o.6", Jul: "o.7", Aug: "o.8", Sep: "o.9", Oct: "o.10", Nov: "o.11", Dec: "o.12"},
			Narrow:      cldr.CalendarMonthFormatNameValue{},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "pikítíkítie, oólí ú kutúan", Feb: "siɛyɛ́, oóli ú kándíɛ", Mar: "ɔnsúmbɔl, oóli ú kátátúɛ", Apr: "mesiŋ, oóli ú kénie", May: "ensil, oóli ú kátánuɛ", Jun: "ɔsɔn", Jul: "efute", Aug: "pisuyú", Sep: "imɛŋ i puɔs", Oct: "imɛŋ i putúk,oóli ú kátíɛ", Nov: "makandikɛ", Dec: "pilɔndɔ́"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "sd", Mon: "md", Tue: "mw", Wed: "et", Thu: "kl", Fri: "fl", Sat: "ss"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "s", Mon: "m", Tue: "m", Wed: "e", Thu: "k", Fri: "f", Sat: "s"},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "sɔ́ndiɛ", Mon: "móndie", Tue: "muányáŋmóndie", Wed: "metúkpíápɛ", Thu: "kúpélimetúkpiapɛ", Fri: "feléte", Sat: "séselé"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "kiɛmɛ́ɛm", PM: "kisɛ́ndɛ"},
		},
	},
}
