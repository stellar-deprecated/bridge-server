package rof

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"},
		Time:     cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "M1", Feb: "M2", Mar: "M3", Apr: "M4", May: "M5", Jun: "M6", Jul: "M7", Aug: "M8", Sep: "M9", Oct: "M10", Nov: "M11", Dec: "M12"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "K", Feb: "K", Mar: "K", Apr: "K", May: "T", Jun: "S", Jul: "S", Aug: "N", Sep: "T", Oct: "I", Nov: "I", Dec: "I"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Mweri wa kwanza", Feb: "Mweri wa kaili", Mar: "Mweri wa katatu", Apr: "Mweri wa kaana", May: "Mweri wa tanu", Jun: "Mweri wa sita", Jul: "Mweri wa saba", Aug: "Mweri wa nane", Sep: "Mweri wa tisa", Oct: "Mweri wa ikumi", Nov: "Mweri wa ikumi na moja", Dec: "Mweri wa ikumi na mbili"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Ijp", Mon: "Ijt", Tue: "Ijn", Wed: "Ijtn", Thu: "Alh", Fri: "Iju", Sat: "Ijm"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "2", Mon: "3", Tue: "4", Wed: "5", Thu: "6", Fri: "7", Sat: "1"},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "Ijumapili", Mon: "Ijumatatu", Tue: "Ijumanne", Wed: "Ijumatano", Thu: "Alhamisi", Fri: "Ijumaa", Sat: "Ijumamosi"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "kang’ama", PM: "kingoto"},
		},
	},
}
