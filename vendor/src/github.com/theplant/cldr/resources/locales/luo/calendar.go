package luo

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"},
		Time:     cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "DAC", Feb: "DAR", Mar: "DAD", Apr: "DAN", May: "DAH", Jun: "DAU", Jul: "DAO", Aug: "DAB", Sep: "DOC", Oct: "DAP", Nov: "DGI", Dec: "DAG"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "C", Feb: "R", Mar: "D", Apr: "N", May: "B", Jun: "U", Jul: "B", Aug: "B", Sep: "C", Oct: "P", Nov: "C", Dec: "P"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Dwe mar Achiel", Feb: "Dwe mar Ariyo", Mar: "Dwe mar Adek", Apr: "Dwe mar Ang’wen", May: "Dwe mar Abich", Jun: "Dwe mar Auchiel", Jul: "Dwe mar Abiriyo", Aug: "Dwe mar Aboro", Sep: "Dwe mar Ochiko", Oct: "Dwe mar Apar", Nov: "Dwe mar gi achiel", Dec: "Dwe mar Apar gi ariyo"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "JMP", Mon: "WUT", Tue: "TAR", Wed: "TAD", Thu: "TAN", Fri: "TAB", Sat: "NGS"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "J", Mon: "W", Tue: "T", Wed: "T", Thu: "T", Fri: "T", Sat: "N"},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "Jumapil", Mon: "Wuok Tich", Tue: "Tich Ariyo", Wed: "Tich Adek", Thu: "Tich Ang’wen", Fri: "Tich Abich", Sat: "Ngeso"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "OD", PM: "OT"},
		},
	},
}
