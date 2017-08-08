package kea

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, d 'di' MMMM 'di' y", Long: "d 'di' MMMM 'di' y", Medium: "d MMM y", Short: "d/M/y"},
		Time:     cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
		DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Jan", Feb: "Feb", Mar: "Mar", Apr: "Abr", May: "Mai", Jun: "Jun", Jul: "Jul", Aug: "Ago", Sep: "Set", Oct: "Otu", Nov: "Nuv", Dec: "Diz"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Janeru", Feb: "Febreru", Mar: "Marsu", Apr: "Abril", May: "Maiu", Jun: "Junhu", Jul: "Julhu", Aug: "Agostu", Sep: "Setenbru", Oct: "Otubru", Nov: "Nuvenbru", Dec: "Dizenbru"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "dum", Mon: "sig", Tue: "ter", Wed: "kua", Thu: "kin", Fri: "ses", Sat: "sab"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "d", Mon: "s", Tue: "t", Wed: "k", Thu: "k", Fri: "s", Sat: "s"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "du", Mon: "si", Tue: "te", Wed: "ku", Thu: "ki", Fri: "se", Sat: "sa"},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "dumingu", Mon: "sigunda-fera", Tue: "tersa-fera", Wed: "kuarta-fera", Thu: "kinta-fera", Fri: "sesta-fera", Sat: "sabadu"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "a", PM: "p"},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "am", PM: "pm"},
		},
	},
}
