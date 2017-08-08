package gd

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, d'mh' MMMM y", Long: "d'mh' MMMM y", Medium: "d MMM y", Short: "dd/MM/y"},
		Time:     cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
		DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Faoi", Feb: "Gearr", Mar: "Màrt", Apr: "Gibl", May: "Cèit", Jun: "Ògmh", Jul: "Iuch", Aug: "Lùna", Sep: "Sult", Oct: "Dàmh", Nov: "Samh", Dec: "Dùbh"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "F", Feb: "G", Mar: "M", Apr: "G", May: "C", Jun: "Ò", Jul: "I", Aug: "L", Sep: "S", Oct: "D", Nov: "S", Dec: "D"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Am Faoilleach", Feb: "An Gearran", Mar: "Am Màrt", Apr: "An Giblean", May: "An Cèitean", Jun: "An t-Ògmhios", Jul: "An t-Iuchar", Aug: "An Lùnastal", Sep: "An t-Sultain", Oct: "An Dàmhair", Nov: "An t-Samhain", Dec: "An Dùbhlachd"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "DiD", Mon: "DiL", Tue: "DiM", Wed: "DiC", Thu: "Dia", Fri: "Dih", Sat: "DiS"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "D", Mon: "L", Tue: "M", Wed: "C", Thu: "A", Fri: "H", Sat: "S"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "Dò", Mon: "Lu", Tue: "Mà", Wed: "Ci", Thu: "Da", Fri: "hA", Sat: "Sa"},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "DiDòmhnaich", Mon: "DiLuain", Tue: "DiMàirt", Wed: "DiCiadain", Thu: "DiarDaoin", Fri: "DihAoine", Sat: "DiSathairne"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "m", PM: "f"},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "m", PM: "f"},
		},
	},
}
