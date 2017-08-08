package gl

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE dd MMMM y", Long: "dd MMMM y", Medium: "d MMM, y", Short: "dd/MM/yy"},
		Time:     cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
		DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Xan", Feb: "Feb", Mar: "Mar", Apr: "Abr", May: "Mai", Jun: "Xuñ", Jul: "Xul", Aug: "Ago", Sep: "Set", Oct: "Out", Nov: "Nov", Dec: "Dec"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "X", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "X", Jul: "X", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Xaneiro", Feb: "Febreiro", Mar: "Marzo", Apr: "Abril", May: "Maio", Jun: "Xuño", Jul: "Xullo", Aug: "Agosto", Sep: "Setembro", Oct: "Outubro", Nov: "Novembro", Dec: "Decembro"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Dom", Mon: "Lun", Tue: "Mar", Wed: "Mér", Thu: "Xov", Fri: "Ven", Sat: "Sáb"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "D", Mon: "L", Tue: "M", Wed: "M", Thu: "X", Fri: "V", Sat: "S"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "Dom", Mon: "Luns", Tue: "Mt", Wed: "Mc", Thu: "Xv", Fri: "Ven", Sat: "Sáb"},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "Domingo", Mon: "Luns", Tue: "Martes", Wed: "Mércores", Thu: "Xoves", Fri: "Venres", Sat: "Sábado"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "a.m.", PM: "p.m."},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "a", PM: "p"},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "a.m.", PM: "p.m."},
		},
	},
}
