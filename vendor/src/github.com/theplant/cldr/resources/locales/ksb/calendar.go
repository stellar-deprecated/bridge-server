package ksb

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"},
		Time:     cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Jan", Feb: "Feb", Mar: "Mac", Apr: "Apr", May: "Mei", Jun: "Jun", Jul: "Jul", Aug: "Ago", Sep: "Sep", Oct: "Okt", Nov: "Nov", Dec: "Des"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Januali", Feb: "Febluali", Mar: "Machi", Apr: "Aplili", May: "Mei", Jun: "Juni", Jul: "Julai", Aug: "Agosti", Sep: "Septemba", Oct: "Oktoba", Nov: "Novemba", Dec: "Desemba"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Jpi", Mon: "Jtt", Tue: "Jmn", Wed: "Jtn", Thu: "Alh", Fri: "Iju", Sat: "Jmo"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "2", Mon: "3", Tue: "4", Wed: "5", Thu: "A", Fri: "I", Sat: "1"},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "Jumaapii", Mon: "Jumaatatu", Tue: "Jumaane", Wed: "Jumaatano", Thu: "Alhamisi", Fri: "Ijumaa", Sat: "Jumaamosi"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "makeo", PM: "nyiaghuo"},
		},
	},
}
