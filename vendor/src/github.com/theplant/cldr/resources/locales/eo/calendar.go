package eo

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, d-'a' 'de' MMMM y", Long: "y-MMMM-dd", Medium: "y-MMM-dd", Short: "yy-MM-dd"},
		Time:     cldr.CalendarDateFormat{Full: "H-'a' 'horo' 'kaj' m:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "jan", Feb: "feb", Mar: "mar", Apr: "apr", May: "maj", Jun: "jun", Jul: "jul", Aug: "aŭg", Sep: "sep", Oct: "okt", Nov: "nov", Dec: "dec"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "januaro", Feb: "februaro", Mar: "marto", Apr: "aprilo", May: "majo", Jun: "junio", Jul: "julio", Aug: "aŭgusto", Sep: "septembro", Oct: "oktobro", Nov: "novembro", Dec: "decembro"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "di", Mon: "lu", Tue: "ma", Wed: "me", Thu: "ĵa", Fri: "ve", Sat: "sa"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "D", Mon: "L", Tue: "M", Wed: "M", Thu: "Ĵ", Fri: "V", Sat: "S"},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "dimanĉo", Mon: "lundo", Tue: "mardo", Wed: "merkredo", Thu: "ĵaŭdo", Fri: "vendredo", Sat: "sabato"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "a", PM: "p"},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "atm", PM: "ptm"},
		},
	},
}
