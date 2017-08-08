package dyo

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "d/M/y"},
		Time:     cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Sa", Feb: "Fe", Mar: "Ma", Apr: "Ab", May: "Me", Jun: "Su", Jul: "Sú", Aug: "Ut", Sep: "Se", Oct: "Ok", Nov: "No", Dec: "De"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "S", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "S", Jul: "S", Aug: "U", Sep: "S", Oct: "O", Nov: "N", Dec: "D"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Sanvie", Feb: "Fébirie", Mar: "Mars", Apr: "Aburil", May: "Mee", Jun: "Sueŋ", Jul: "Súuyee", Aug: "Ut", Sep: "Settembar", Oct: "Oktobar", Nov: "Novembar", Dec: "Disambar"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Dim", Mon: "Ten", Tue: "Tal", Wed: "Ala", Thu: "Ara", Fri: "Arj", Sat: "Sib"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "D", Mon: "T", Tue: "T", Wed: "A", Thu: "A", Fri: "A", Sat: "S"},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "Dimas", Mon: "Teneŋ", Tue: "Talata", Wed: "Alarbay", Thu: "Aramisay", Fri: "Arjuma", Sat: "Sibiti"},
		},
		Periods: cldr.CalendarPeriodFormatNames{},
	},
}
