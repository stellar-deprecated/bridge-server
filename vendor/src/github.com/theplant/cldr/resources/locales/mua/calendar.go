package mua

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "d/M/y"},
		Time:     cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "FLO", Feb: "CLA", Mar: "CKI", Apr: "FMF", May: "MAD", Jun: "MBI", Jul: "MLI", Aug: "MAM", Sep: "FDE", Oct: "FMU", Nov: "FGW", Dec: "FYU"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "O", Feb: "A", Mar: "I", Apr: "F", May: "D", Jun: "B", Jul: "L", Aug: "M", Sep: "E", Oct: "U", Nov: "W", Dec: "Y"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Fĩi Loo", Feb: "Cokcwaklaŋne", Mar: "Cokcwaklii", Apr: "Fĩi Marfoo", May: "Madǝǝuutǝbijaŋ", Jun: "Mamǝŋgwãafahbii", Jul: "Mamǝŋgwãalii", Aug: "Madǝmbii", Sep: "Fĩi Dǝɓlii", Oct: "Fĩi Mundaŋ", Nov: "Fĩi Gwahlle", Dec: "Fĩi Yuru"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Cya", Mon: "Cla", Tue: "Czi", Wed: "Cko", Thu: "Cka", Fri: "Cga", Sat: "Cze"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "Y", Mon: "L", Tue: "Z", Wed: "O", Thu: "A", Fri: "G", Sat: "E"},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "Com’yakke", Mon: "Comlaaɗii", Tue: "Comzyiiɗii", Wed: "Comkolle", Thu: "Comkaldǝɓlii", Fri: "Comgaisuu", Sat: "Comzyeɓsuu"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "comme", PM: "lilli"},
		},
	},
}
