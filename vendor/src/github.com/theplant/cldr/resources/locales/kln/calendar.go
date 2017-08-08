package kln

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"},
		Time:     cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Mul", Feb: "Ngat", Mar: "Taa", Apr: "Iwo", May: "Mam", Jun: "Paa", Jul: "Nge", Aug: "Roo", Sep: "Bur", Oct: "Epe", Nov: "Kpt", Dec: "Kpa"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "M", Feb: "N", Mar: "T", Apr: "I", May: "M", Jun: "P", Jul: "N", Aug: "R", Sep: "B", Oct: "E", Nov: "K", Dec: "K"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Mulgul", Feb: "Ng’atyaato", Mar: "Kiptaamo", Apr: "Iwootkuut", May: "Mamuut", Jun: "Paagi", Jul: "Ng’eiyeet", Aug: "Rooptui", Sep: "Bureet", Oct: "Epeeso", Nov: "Kipsuunde ne taai", Dec: "Kipsuunde nebo aeng’"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Kts", Mon: "Kot", Tue: "Koo", Wed: "Kos", Thu: "Koa", Fri: "Kom", Sat: "Kol"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "T", Mon: "T", Tue: "O", Wed: "S", Thu: "A", Fri: "M", Sat: "L"},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "Kotisap", Mon: "Kotaai", Tue: "Koaeng’", Wed: "Kosomok", Thu: "Koang’wan", Fri: "Komuut", Sat: "Kolo"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "krn", PM: "koosk"},
			Narrow:      cldr.CalendarPeriodFormatNameValue{},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "karoon", PM: "kooskoliny"},
		},
	},
}
