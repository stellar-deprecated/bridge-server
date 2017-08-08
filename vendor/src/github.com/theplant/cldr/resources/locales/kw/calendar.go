package kw

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"},
		Time:     cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Gen", Feb: "Hwe", Mar: "Meu", Apr: "Ebr", May: "Me", Jun: "Met", Jul: "Gor", Aug: "Est", Sep: "Gwn", Oct: "Hed", Nov: "Du", Dec: "Kev"},
			Narrow:      cldr.CalendarMonthFormatNameValue{},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "mis Genver", Feb: "mis Hwevrer", Mar: "mis Meurth", Apr: "mis Ebrel", May: "mis Me", Jun: "mis Metheven", Jul: "mis Gortheren", Aug: "mis Est", Sep: "mis Gwynngala", Oct: "mis Hedra", Nov: "mis Du", Dec: "mis Kevardhu"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Sul", Mon: "Lun", Tue: "Mth", Wed: "Mhr", Thu: "Yow", Fri: "Gwe", Sat: "Sad"},
			Narrow:      cldr.CalendarDayFormatNameValue{},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "dy Sul", Mon: "dy Lun", Tue: "dy Meurth", Wed: "dy Merher", Thu: "dy Yow", Fri: "dy Gwener", Sat: "dy Sadorn"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "a.m.", PM: "p.m."},
		},
	},
}
