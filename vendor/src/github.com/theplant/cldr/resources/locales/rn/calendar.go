package rn

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "d/M/y"},
		Time:     cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Mut.", Feb: "Gas.", Mar: "Wer.", Apr: "Mat.", May: "Gic.", Jun: "Kam.", Jul: "Nya.", Aug: "Kan.", Sep: "Nze.", Oct: "Ukw.", Nov: "Ugu.", Dec: "Uku."},
			Narrow:      cldr.CalendarMonthFormatNameValue{},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Nzero", Feb: "Ruhuhuma", Mar: "Ntwarante", Apr: "Ndamukiza", May: "Rusama", Jun: "Ruheshi", Jul: "Mukakaro", Aug: "Nyandagaro", Sep: "Nyakanga", Oct: "Gitugutu", Nov: "Munyonyo", Dec: "Kigarama"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "cu.", Mon: "mbe.", Tue: "kab.", Wed: "gtu.", Thu: "kan.", Fri: "gnu.", Sat: "gnd."},
			Narrow:      cldr.CalendarDayFormatNameValue{},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "Ku w’indwi", Mon: "Ku wa mbere", Tue: "Ku wa kabiri", Wed: "Ku wa gatatu", Thu: "Ku wa kane", Fri: "Ku wa gatanu", Sat: "Ku wa gatandatu"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "Z.MU.", PM: "Z.MW."},
		},
	},
}
