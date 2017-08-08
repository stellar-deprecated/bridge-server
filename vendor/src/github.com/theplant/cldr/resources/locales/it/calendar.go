package it

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "dd MMM y", Short: "dd/MM/yy"},
		Time:     cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
		DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1}, {0}", Short: "{1}, {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "gen", Feb: "feb", Mar: "mar", Apr: "apr", May: "mag", Jun: "giu", Jul: "lug", Aug: "ago", Sep: "set", Oct: "ott", Nov: "nov", Dec: "dic"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "G", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "G", Jul: "L", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Gennaio", Feb: "Febbraio", Mar: "Marzo", Apr: "Aprile", May: "Maggio", Jun: "Giugno", Jul: "Luglio", Aug: "Agosto", Sep: "Settembre", Oct: "Ottobre", Nov: "Novembre", Dec: "Dicembre"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "dom", Mon: "lun", Tue: "mar", Wed: "mer", Thu: "gio", Fri: "ven", Sat: "sab"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "D", Mon: "L", Tue: "M", Wed: "M", Thu: "G", Fri: "V", Sat: "S"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "dom", Mon: "lun", Tue: "mar", Wed: "mer", Thu: "gio", Fri: "ven", Sat: "sab"},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "Domenica", Mon: "Lunedì", Tue: "Martedì", Wed: "Mercoledì", Thu: "Giovedì", Fri: "Venerdì", Sat: "Sabato"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "m.", PM: "p."},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"},
		},
	},
}
