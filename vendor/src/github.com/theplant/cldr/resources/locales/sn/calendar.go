package sn

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"},
		Time:     cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Ndi", Feb: "Kuk", Mar: "Kur", Apr: "Kub", May: "Chv", Jun: "Chk", Jul: "Chg", Aug: "Nya", Sep: "Gun", Oct: "Gum", Nov: "Mb", Dec: "Zvi"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "N", Feb: "K", Mar: "K", Apr: "K", May: "C", Jun: "C", Jul: "C", Aug: "N", Sep: "G", Oct: "G", Nov: "M", Dec: "Z"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Ndira", Feb: "Kukadzi", Mar: "Kurume", Apr: "Kubvumbi", May: "Chivabvu", Jun: "Chikumi", Jul: "Chikunguru", Aug: "Nyamavhuvhu", Sep: "Gunyana", Oct: "Gumiguru", Nov: "Mbudzi", Dec: "Zvita"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Svo", Mon: "Muv", Tue: "Chip", Wed: "Chit", Thu: "Chin", Fri: "Chis", Sat: "Mug"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "S", Mon: "M", Tue: "C", Wed: "C", Thu: "C", Fri: "C", Sat: "M"},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "Svondo", Mon: "Muvhuro", Tue: "Chipiri", Wed: "Chitatu", Thu: "China", Fri: "Chishanu", Sat: "Mugovera"},
		},
		Periods: cldr.CalendarPeriodFormatNames{},
	},
}
