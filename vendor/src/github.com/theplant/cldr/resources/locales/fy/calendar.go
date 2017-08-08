package fy

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd-MM-yy"},
		Time:     cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
		DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "jan", Feb: "feb", Mar: "mrt", Apr: "apr", May: "mai", Jun: "jun", Jul: "jul", Aug: "aug", Sep: "sep", Oct: "okt", Nov: "nov", Dec: "des"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "jannewaris", Feb: "febrewaris", Mar: "maart", Apr: "april", May: "maaie", Jun: "juny", Jul: "july", Aug: "augustus", Sep: "septimber", Oct: "oktober", Nov: "novimber", Dec: "desimber"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "si", Mon: "mo", Tue: "ti", Wed: "wo", Thu: "to", Fri: "fr", Sat: "so"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "Z", Mon: "M", Tue: "D", Wed: "W", Thu: "D", Fri: "V", Sat: "Z"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "si", Mon: "mo", Tue: "ti", Wed: "wo", Thu: "to", Fri: "fr", Sat: "so"},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "snein", Mon: "moandei", Tue: "tiisdei", Wed: "woansdei", Thu: "tongersdei", Fri: "freed", Sat: "sneon"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "a.m.", PM: "p.m."},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "foarmiddei", PM: "p.m."},
		},
	},
}
