package ast

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, d MMMM 'de' y", Long: "d MMMM 'de' y", Medium: "d MMM y", Short: "d/M/yy"},
		Time:     cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
		DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Xin", Feb: "Feb", Mar: "Mar", Apr: "Abr", May: "May", Jun: "Xun", Jul: "Xnt", Aug: "Ago", Sep: "Set", Oct: "Och", Nov: "Pay", Dec: "Avi"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "X", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "X", Jul: "X", Aug: "A", Sep: "S", Oct: "O", Nov: "P", Dec: "A"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "xineru", Feb: "febreru", Mar: "marzu", Apr: "abril", May: "mayu", Jun: "xunu", Jul: "xunetu", Aug: "agostu", Sep: "setiembre", Oct: "ochobre", Nov: "payares", Dec: "avientu"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "dom", Mon: "llu", Tue: "mar", Wed: "mie", Thu: "xue", Fri: "vie", Sat: "sab"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "D", Mon: "L", Tue: "M", Wed: "M", Thu: "X", Fri: "V", Sat: "S"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "do", Mon: "ll", Tue: "ma", Wed: "mi", Thu: "xu", Fri: "vi", Sat: "sa"},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "domingu", Mon: "llunes", Tue: "martes", Wed: "miércoles", Thu: "xueves", Fri: "vienres", Sat: "sábadu"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "a", PM: "p"},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"},
		},
	},
}
