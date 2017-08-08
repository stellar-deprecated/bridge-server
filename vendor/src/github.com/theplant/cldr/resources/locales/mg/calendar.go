package mg

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM, y", Short: "d/M/y"},
		Time:     cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Jan", Feb: "Feb", Mar: "Mar", Apr: "Apr", May: "Mey", Jun: "Jon", Jul: "Jol", Aug: "Aog", Sep: "Sep", Oct: "Okt", Nov: "Nov", Dec: "Des"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Janoary", Feb: "Febroary", Mar: "Martsa", Apr: "Aprily", May: "Mey", Jun: "Jona", Jul: "Jolay", Aug: "Aogositra", Sep: "Septambra", Oct: "Oktobra", Nov: "Novambra", Dec: "Desambra"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Alah", Mon: "Alats", Tue: "Tal", Wed: "Alar", Thu: "Alak", Fri: "Zom", Sat: "Asab"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "A", Mon: "A", Tue: "T", Wed: "A", Thu: "A", Fri: "Z", Sat: "A"},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "Alahady", Mon: "Alatsinainy", Tue: "Talata", Wed: "Alarobia", Thu: "Alakamisy", Fri: "Zoma", Sat: "Asabotsy"},
		},
		Periods: cldr.CalendarPeriodFormatNames{},
	},
}
