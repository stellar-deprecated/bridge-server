package vai_Latn

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"},
		Time:     cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{},
			Narrow:      cldr.CalendarMonthFormatNameValue{},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "luukao kemã", Feb: "ɓandaɓu", Mar: "vɔɔ", Apr: "fulu", May: "goo", Jun: "6", Jul: "7", Aug: "kɔnde", Sep: "saah", Oct: "galo", Nov: "kenpkato ɓololɔ", Dec: "luukao lɔma"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{},
			Narrow:      cldr.CalendarDayFormatNameValue{},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "lahadi", Mon: "tɛɛnɛɛ", Tue: "talata", Wed: "alaba", Thu: "aimisa", Fri: "aijima", Sat: "siɓiti"},
		},
		Periods: cldr.CalendarPeriodFormatNames{},
	},
}
