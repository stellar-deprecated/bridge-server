package hsb

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, d. MMMM y", Long: "d. MMMM y", Medium: "d.M.y", Short: "d.M.yy"},
		Time:     cldr.CalendarDateFormat{Full: "H:mm:ss zzzz", Long: "H:mm:ss z", Medium: "H:mm:ss", Short: "H:mm 'hodź'."},
		DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "jan", Feb: "feb", Mar: "měr", Apr: "apr", May: "mej", Jun: "jun", Jul: "jul", Aug: "awg", Sep: "sep", Oct: "okt", Nov: "now", Dec: "dec"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "j", Feb: "f", Mar: "m", Apr: "a", May: "m", Jun: "j", Jul: "j", Aug: "a", Sep: "s", Oct: "o", Nov: "n", Dec: "d"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "januar", Feb: "februar", Mar: "měrc", Apr: "apryl", May: "meja", Jun: "junij", Jul: "julij", Aug: "awgust", Sep: "september", Oct: "oktober", Nov: "nowember", Dec: "december"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "nje", Mon: "pón", Tue: "wut", Wed: "srj", Thu: "štw", Fri: "pja", Sat: "sob"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "n", Mon: "p", Tue: "w", Wed: "s", Thu: "š", Fri: "p", Sat: "s"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "nj", Mon: "pó", Tue: "wu", Wed: "sr", Thu: "št", Fri: "pj", Sat: "so"},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "njedźela", Mon: "póndźela", Tue: "wutora", Wed: "srjeda", Thu: "štwórtk", Fri: "pjatk", Sat: "sobota"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "dop.", PM: "pop."},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "dopołdnja", PM: "popołdnju"},
		},
	},
}
