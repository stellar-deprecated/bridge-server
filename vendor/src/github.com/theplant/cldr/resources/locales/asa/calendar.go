package asa

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"},
		Time:     cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Jan", Feb: "Feb", Mar: "Mac", Apr: "Apr", May: "Mei", Jun: "Jun", Jul: "Jul", Aug: "Ago", Sep: "Sep", Oct: "Okt", Nov: "Nov", Dec: "Dec"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Januari", Feb: "Februari", Mar: "Machi", Apr: "Aprili", May: "Mei", Jun: "Juni", Jul: "Julai", Aug: "Agosti", Sep: "Septemba", Oct: "Oktoba", Nov: "Novemba", Dec: "Desemba"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Jpi", Mon: "Jtt", Tue: "Jnn", Wed: "Jtn", Thu: "Alh", Fri: "Ijm", Sat: "Jmo"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "J", Mon: "J", Tue: "J", Wed: "J", Thu: "A", Fri: "I", Sat: "J"},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "Jumapili", Mon: "Jumatatu", Tue: "Jumanne", Wed: "Jumatano", Thu: "Alhamisi", Fri: "Ijumaa", Sat: "Jumamosi"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "icheheavo", PM: "ichamthi"},
		},
	},
}
