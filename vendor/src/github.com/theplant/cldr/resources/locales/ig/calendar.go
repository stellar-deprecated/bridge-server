package ig

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"},
		Time:     cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Jen", Feb: "Feb", Mar: "Maa", Apr: "Epr", May: "Mee", Jun: "Juu", Jul: "Jul", Aug: "Ọgọ", Sep: "Sep", Oct: "Ọkt", Nov: "Nov", Dec: "Dis"},
			Narrow:      cldr.CalendarMonthFormatNameValue{},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Jenụwarị", Feb: "Febrụwarị", Mar: "Maachị", Apr: "Eprel", May: "Mee", Jun: "Juun", Jul: "Julaị", Aug: "Ọgọọst", Sep: "Septemba", Oct: "Ọktoba", Nov: "Novemba", Dec: "Disemba"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Ụka", Mon: "Mọn", Tue: "Tiu", Wed: "Wen", Thu: "Tọọ", Fri: "Fraị", Sat: "Sat"},
			Narrow:      cldr.CalendarDayFormatNameValue{},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "Mbọsị Ụka", Mon: "Mọnde", Tue: "Tiuzdee", Wed: "Wenezdee", Thu: "Tọọzdee", Fri: "Fraịdee", Sat: "Satọdee"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "A.M.", PM: "P.M."},
		},
	},
}
