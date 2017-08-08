package wae

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, d. MMMM y", Long: "d. MMMM y", Medium: "d. MMM y", Short: ""},
		Time:     cldr.CalendarDateFormat{},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Jen", Feb: "Hor", Mar: "Mär", Apr: "Abr", May: "Mei", Jun: "Brá", Jul: "Hei", Aug: "Öig", Sep: "Her", Oct: "Wím", Nov: "Win", Dec: "Chr"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "H", Mar: "M", Apr: "A", May: "M", Jun: "B", Jul: "H", Aug: "Ö", Sep: "H", Oct: "W", Nov: "W", Dec: "C"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Jenner", Feb: "Hornig", Mar: "Märze", Apr: "Abrille", May: "Meije", Jun: "Bráčet", Jul: "Heiwet", Aug: "Öigšte", Sep: "Herbštmánet", Oct: "Wímánet", Nov: "Wintermánet", Dec: "Chrištmánet"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Sun", Mon: "Män", Tue: "Ziš", Wed: "Mit", Thu: "Fró", Fri: "Fri", Sat: "Sam"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "S", Mon: "M", Tue: "Z", Wed: "M", Thu: "F", Fri: "F", Sat: "S"},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "Sunntag", Mon: "Mäntag", Tue: "Zištag", Wed: "Mittwuč", Thu: "Fróntag", Fri: "Fritag", Sat: "Samštag"},
		},
		Periods: cldr.CalendarPeriodFormatNames{},
	},
}
