package mer

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"},
		Time:     cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "JAN", Feb: "FEB", Mar: "MAC", Apr: "ĨPU", May: "MĨĨ", Jun: "NJU", Jul: "NJR", Aug: "AGA", Sep: "SPT", Oct: "OKT", Nov: "NOV", Dec: "DEC"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "Ĩ", May: "M", Jun: "N", Jul: "N", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Januarĩ", Feb: "Feburuarĩ", Mar: "Machi", Apr: "Ĩpurũ", May: "Mĩĩ", Jun: "Njuni", Jul: "Njuraĩ", Aug: "Agasti", Sep: "Septemba", Oct: "Oktũba", Nov: "Novemba", Dec: "Dicemba"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "KIU", Mon: "MRA", Tue: "WAI", Wed: "WET", Thu: "WEN", Fri: "WTN", Sat: "JUM"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "K", Mon: "M", Tue: "W", Wed: "W", Thu: "W", Fri: "W", Sat: "J"},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "Kiumia", Mon: "Muramuko", Tue: "Wairi", Wed: "Wethatu", Thu: "Wena", Fri: "Wetano", Sat: "Jumamosi"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "RŨ", PM: "ŨG"},
		},
	},
}
