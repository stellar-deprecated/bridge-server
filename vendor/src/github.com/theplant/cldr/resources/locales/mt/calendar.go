package mt

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, d 'ta'’ MMMM y", Long: "d 'ta'’ MMMM y", Medium: "dd MMM y", Short: "dd/MM/y"},
		Time:     cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Jan", Feb: "Fra", Mar: "Mar", Apr: "Apr", May: "Mej", Jun: "Ġun", Jul: "Lul", Aug: "Aww", Sep: "Set", Oct: "Ott", Nov: "Nov", Dec: "Diċ"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "Jn", Feb: "Fr", Mar: "Mz", Apr: "Ap", May: "Mj", Jun: "Ġn", Jul: "Lj", Aug: "Aw", Sep: "St", Oct: "Ob", Nov: "Nv", Dec: "Dċ"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Jannar", Feb: "Frar", Mar: "Marzu", Apr: "April", May: "Mejju", Jun: "Ġunju", Jul: "Lulju", Aug: "Awwissu", Sep: "Settembru", Oct: "Ottubru", Nov: "Novembru", Dec: "Diċembru"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Ħad", Mon: "Tne", Tue: "Tli", Wed: "Erb", Thu: "Ħam", Fri: "Ġim", Sat: "Sib"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "Ħd", Mon: "Tn", Tue: "Tl", Wed: "Er", Thu: "Ħm", Fri: "Ġm", Sat: "Sb"},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "Il-Ħadd", Mon: "It-Tnejn", Tue: "It-Tlieta", Wed: "L-Erbgħa", Thu: "Il-Ħamis", Fri: "Il-Ġimgħa", Sat: "Is-Sibt"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"},
		},
	},
}
