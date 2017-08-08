package sg

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM, y", Short: "d/M/y"},
		Time:     cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Nye", Feb: "Ful", Mar: "Mbä", Apr: "Ngu", May: "Bêl", Jun: "Fön", Jul: "Len", Aug: "Kük", Sep: "Mvu", Oct: "Ngb", Nov: "Nab", Dec: "Kak"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "N", Feb: "F", Mar: "M", Apr: "N", May: "B", Jun: "F", Jul: "L", Aug: "K", Sep: "M", Oct: "N", Nov: "N", Dec: "K"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Nyenye", Feb: "Fulundïgi", Mar: "Mbängü", Apr: "Ngubùe", May: "Bêläwü", Jun: "Föndo", Jul: "Lengua", Aug: "Kükürü", Sep: "Mvuka", Oct: "Ngberere", Nov: "Nabändüru", Dec: "Kakauka"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Bk1", Mon: "Bk2", Tue: "Bk3", Wed: "Bk4", Thu: "Bk5", Fri: "Lâp", Sat: "Lây"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "K", Mon: "S", Tue: "T", Wed: "S", Thu: "K", Fri: "P", Sat: "Y"},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "Bikua-ôko", Mon: "Bïkua-ûse", Tue: "Bïkua-ptâ", Wed: "Bïkua-usïö", Thu: "Bïkua-okü", Fri: "Lâpôsö", Sat: "Lâyenga"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "ND", PM: "LK"},
		},
	},
}
