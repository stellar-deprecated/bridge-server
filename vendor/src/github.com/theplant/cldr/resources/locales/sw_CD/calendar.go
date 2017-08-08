package sw_CD

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "", Medium: "", Short: "d/M/y"},
		Time:     cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "mkw", Feb: "mpi", Mar: "mtu", Apr: "min", May: "mtn", Jun: "mst", Jul: "msb", Aug: "mun", Sep: "mts", Oct: "mku", Nov: "mkm", Dec: "mkb"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "k", Feb: "p", Mar: "t", Apr: "i", May: "t", Jun: "s", Jul: "s", Aug: "m", Sep: "t", Oct: "k", Nov: "m", Dec: "m"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "mwezi ya kwanja", Feb: "mwezi ya pili", Mar: "mwezi ya tatu", Apr: "mwezi ya ine", May: "mwezi ya tanu", Jun: "mwezi ya sita", Jul: "mwezi ya saba", Aug: "mwezi ya munane", Sep: "mwezi ya tisa", Oct: "mwezi ya kumi", Nov: "mwezi ya kumi na moya", Dec: "mwezi ya kumi ya mbili"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "yen", Mon: "kwa", Tue: "pil", Wed: "tat", Thu: "ine", Fri: "tan", Sat: "sit"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "y", Mon: "k", Tue: "p", Wed: "t", Thu: "i", Fri: "t", Sat: "s"},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "siku ya yenga", Mon: "siku ya kwanza", Tue: "siku ya pili", Wed: "siku ya tatu", Thu: "siku ya ine", Fri: "siku ya tanu", Sat: "siku ya sita"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "ya asubuyi", PM: "ya muchana"},
		},
	},
}
