package ebu

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"},
		Time:     cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Mbe", Feb: "Kai", Mar: "Kat", Apr: "Kan", May: "Gat", Jun: "Gan", Jul: "Mug", Aug: "Knn", Sep: "Ken", Oct: "Iku", Nov: "Imw", Dec: "Igi"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "M", Feb: "K", Mar: "K", Apr: "K", May: "G", Jun: "G", Jul: "M", Aug: "K", Sep: "K", Oct: "I", Nov: "I", Dec: "I"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Mweri wa mbere", Feb: "Mweri wa kaĩri", Mar: "Mweri wa kathatũ", Apr: "Mweri wa kana", May: "Mweri wa gatano", Jun: "Mweri wa gatantatũ", Jul: "Mweri wa mũgwanja", Aug: "Mweri wa kanana", Sep: "Mweri wa kenda", Oct: "Mweri wa ikũmi", Nov: "Mweri wa ikũmi na ũmwe", Dec: "Mweri wa ikũmi na Kaĩrĩ"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Kma", Mon: "Tat", Tue: "Ine", Wed: "Tan", Thu: "Arm", Fri: "Maa", Sat: "NMM"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "K", Mon: "N", Tue: "N", Wed: "N", Thu: "A", Fri: "M", Sat: "N"},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "Kiumia", Mon: "Njumatatu", Tue: "Njumaine", Wed: "Njumatano", Thu: "Aramithi", Fri: "Njumaa", Sat: "NJumamothii"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "KI", PM: "UT"},
		},
	},
}
