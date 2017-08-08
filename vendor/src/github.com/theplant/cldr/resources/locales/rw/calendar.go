package rw

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, y MMMM dd", Long: "y MMMM d", Medium: "y MMM d", Short: "yy/MM/dd"},
		Time:     cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "mut.", Feb: "gas.", Mar: "wer.", Apr: "mat.", May: "gic.", Jun: "kam.", Jul: "nya.", Aug: "kan.", Sep: "nze.", Oct: "ukw.", Nov: "ugu.", Dec: "uku."},
			Narrow:      cldr.CalendarMonthFormatNameValue{},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Mutarama", Feb: "Gashyantare", Mar: "Werurwe", Apr: "Mata", May: "Gicuransi", Jun: "Kamena", Jul: "Nyakanga", Aug: "Kanama", Sep: "Nzeli", Oct: "Ukwakira", Nov: "Ugushyingo", Dec: "Ukuboza"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "cyu.", Mon: "mbe.", Tue: "kab.", Wed: "gtu.", Thu: "kan.", Fri: "gnu.", Sat: "gnd."},
			Narrow:      cldr.CalendarDayFormatNameValue{},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "Ku cyumweru", Mon: "Kuwa mbere", Tue: "Kuwa kabiri", Wed: "Kuwa gatatu", Thu: "Kuwa kane", Fri: "Kuwa gatanu", Sat: "Kuwa gatandatu"},
		},
		Periods: cldr.CalendarPeriodFormatNames{},
	},
}
