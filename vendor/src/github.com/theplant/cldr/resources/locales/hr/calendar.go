package hr

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, d. MMMM y.", Long: "d. MMMM y.", Medium: "d. MMM y.", Short: "dd.MM.y."},
		Time:     cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
		DateTime: cldr.CalendarDateFormat{Full: "{1} 'u' {0}", Long: "{1} 'u' {0}", Medium: "{1} {0}", Short: "{1} {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "sij", Feb: "velj", Mar: "ožu", Apr: "tra", May: "svi", Jun: "lip", Jul: "srp", Aug: "kol", Sep: "ruj", Oct: "lis", Nov: "stu", Dec: "pro"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "1.", Feb: "2.", Mar: "3.", Apr: "4.", May: "5.", Jun: "6.", Jul: "7.", Aug: "8.", Sep: "9.", Oct: "10.", Nov: "11.", Dec: "12."},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "siječanj", Feb: "veljača", Mar: "ožujak", Apr: "travanj", May: "svibanj", Jun: "lipanj", Jul: "srpanj", Aug: "kolovoz", Sep: "rujan", Oct: "listopad", Nov: "studeni", Dec: "prosinac"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "ned", Mon: "pon", Tue: "uto", Wed: "sri", Thu: "čet", Fri: "pet", Sat: "sub"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "n", Mon: "p", Tue: "u", Wed: "s", Thu: "č", Fri: "p", Sat: "s"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "ned", Mon: "pon", Tue: "uto", Wed: "sri", Thu: "čet", Fri: "pet", Sat: "sub"},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "nedjelja", Mon: "ponedjeljak", Tue: "utorak", Wed: "srijeda", Thu: "četvrtak", Fri: "petak", Sat: "subota"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"},
		},
	},
}
