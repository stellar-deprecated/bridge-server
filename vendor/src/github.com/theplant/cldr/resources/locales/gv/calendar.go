package gv

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE dd MMMM y", Long: "dd MMMM y", Medium: "MMM dd, y", Short: "dd/MM/yy"},
		Time:     cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "J-guer", Feb: "T-arree", Mar: "Mayrnt", Apr: "Avrril", May: "Boaldyn", Jun: "M-souree", Jul: "J-souree", Aug: "Luanistyn", Sep: "M-fouyir", Oct: "J-fouyir", Nov: "M.Houney", Dec: "M.Nollick"},
			Narrow:      cldr.CalendarMonthFormatNameValue{},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Jerrey-geuree", Feb: "Toshiaght-arree", Mar: "Mayrnt", Apr: "Averil", May: "Boaldyn", Jun: "Mean-souree", Jul: "Jerrey-souree", Aug: "Luanistyn", Sep: "Mean-fouyir", Oct: "Jerrey-fouyir", Nov: "Mee Houney", Dec: "Mee ny Nollick"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Jed", Mon: "Jel", Tue: "Jem", Wed: "Jerc", Thu: "Jerd", Fri: "Jeh", Sat: "Jes"},
			Narrow:      cldr.CalendarDayFormatNameValue{},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "Jedoonee", Mon: "Jelhein", Tue: "Jemayrt", Wed: "Jercean", Thu: "Jerdein", Fri: "Jeheiney", Sat: "Jesarn"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "a.m.", PM: "p.m."},
		},
	},
}
