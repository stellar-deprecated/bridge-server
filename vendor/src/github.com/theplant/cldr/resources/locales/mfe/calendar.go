package mfe

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM, y", Short: "d/M/y"},
		Time:     cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "zan", Feb: "fev", Mar: "mar", Apr: "avr", May: "me", Jun: "zin", Jul: "zil", Aug: "out", Sep: "sep", Oct: "okt", Nov: "nov", Dec: "des"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "z", Feb: "f", Mar: "m", Apr: "a", May: "m", Jun: "z", Jul: "z", Aug: "o", Sep: "s", Oct: "o", Nov: "n", Dec: "d"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "zanvie", Feb: "fevriye", Mar: "mars", Apr: "avril", May: "me", Jun: "zin", Jul: "zilye", Aug: "out", Sep: "septam", Oct: "oktob", Nov: "novam", Dec: "desam"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "dim", Mon: "lin", Tue: "mar", Wed: "mer", Thu: "ze", Fri: "van", Sat: "sam"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "d", Mon: "l", Tue: "m", Wed: "m", Thu: "z", Fri: "v", Sat: "s"},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "dimans", Mon: "lindi", Tue: "mardi", Wed: "merkredi", Thu: "zedi", Fri: "vandredi", Sat: "samdi"},
		},
		Periods: cldr.CalendarPeriodFormatNames{},
	},
}
