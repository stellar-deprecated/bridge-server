package seh

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, d 'de' MMMM 'de' y", Long: "d 'de' MMMM 'de' y", Medium: "d 'de' MMM 'de' y", Short: "d/M/y"},
		Time:     cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Jan", Feb: "Fev", Mar: "Mar", Apr: "Abr", May: "Mai", Jun: "Jun", Jul: "Jul", Aug: "Aug", Sep: "Set", Oct: "Otu", Nov: "Nov", Dec: "Dec"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Janeiro", Feb: "Fevreiro", Mar: "Marco", Apr: "Abril", May: "Maio", Jun: "Junho", Jul: "Julho", Aug: "Augusto", Sep: "Setembro", Oct: "Otubro", Nov: "Novembro", Dec: "Decembro"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Dim", Mon: "Pos", Tue: "Pir", Wed: "Tat", Thu: "Nai", Fri: "Sha", Sat: "Sab"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "D", Mon: "P", Tue: "C", Wed: "T", Thu: "N", Fri: "S", Sat: "S"},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "Dimingu", Mon: "Chiposi", Tue: "Chipiri", Wed: "Chitatu", Thu: "Chinai", Fri: "Chishanu", Sat: "Sabudu"},
		},
		Periods: cldr.CalendarPeriodFormatNames{},
	},
}
