package de

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, d. MMMM y", Long: "d. MMMM y", Medium: "dd.MM.y", Short: "dd.MM.yy"},
		Time:     cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
		DateTime: cldr.CalendarDateFormat{Full: "{1} 'um' {0}", Long: "{1} 'um' {0}", Medium: "{1}, {0}", Short: "{1}, {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Jan", Feb: "Feb", Mar: "Mär", Apr: "Apr", May: "Mai", Jun: "Jun", Jul: "Jul", Aug: "Aug", Sep: "Sep", Oct: "Okt", Nov: "Nov", Dec: "Dez"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Januar", Feb: "Februar", Mar: "März", Apr: "April", May: "Mai", Jun: "Juni", Jul: "Juli", Aug: "August", Sep: "September", Oct: "Oktober", Nov: "November", Dec: "Dezember"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "So", Mon: "Mo", Tue: "Di", Wed: "Mi", Thu: "Do", Fri: "Fr", Sat: "Sa"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "S", Mon: "M", Tue: "D", Wed: "M", Thu: "D", Fri: "F", Sat: "S"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "So.", Mon: "Mo.", Tue: "Di.", Wed: "Mi.", Thu: "Do.", Fri: "Fr.", Sat: "Sa."},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "Sonntag", Mon: "Montag", Tue: "Dienstag", Wed: "Mittwoch", Thu: "Donnerstag", Fri: "Freitag", Sat: "Samstag"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "vm.", PM: "nm."},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{},
		},
	},
}
