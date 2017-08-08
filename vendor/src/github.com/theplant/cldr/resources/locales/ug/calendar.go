package ug

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE، MMMM d، y", Long: "MMMM d، y", Medium: "MMM d، y", Short: "M/d/yy"},
		Time:     cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
		DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1}، {0}", Short: "{1}، {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "يانۋار", Feb: "فېۋرال", Mar: "مارت", Apr: "ئاپرېل", May: "ماي", Jun: "ئىيۇن", Jul: "ئىيۇل", Aug: "ئاۋغۇست", Sep: "سېنتەبىر", Oct: "ئۆكتەبىر", Nov: "نويابىر", Dec: "دېكابىر"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "1", Feb: "2", Mar: "3", Apr: "4", May: "5", Jun: "6", Jul: "7", Aug: "8", Sep: "9", Oct: "10", Nov: "11", Dec: "12"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "يانۋار", Feb: "فېۋرال", Mar: "مارت", Apr: "ئاپرېل", May: "ماي", Jun: "ئىيۇن", Jul: "ئىيۇل", Aug: "ئاۋغۇست", Sep: "سېنتەبىر", Oct: "ئۆكتەبىر", Nov: "بويابىر", Dec: "دېكابىر"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "يە", Mon: "دۈ", Tue: "سە", Wed: "چا", Thu: "پە", Fri: "چۈ", Sat: "شە"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "ي", Mon: "د", Tue: "س", Wed: "چ", Thu: "پ", Fri: "ج", Sat: "ش"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "ي", Mon: "د", Tue: "س", Wed: "چ", Thu: "پ", Fri: "ج", Sat: "ش"},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "يەكشەنبە", Mon: "دۈشەنبە", Tue: "سەيشەنبە", Wed: "چارشەنبە", Thu: "پەيشەنبە", Fri: "جۈمە", Sat: "شەنبە"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "چۈشتىن بۇرۇن", PM: "چۈشتىن كېيىن"},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "چۈشتىن بۇرۇن", PM: "چۈشتىن كېيىن"},
		},
	},
}
