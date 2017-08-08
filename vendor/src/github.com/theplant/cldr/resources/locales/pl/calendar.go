package pl

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "dd.MM.y", Short: "dd.MM.y"},
		Time:     cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
		DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1}, {0}", Short: "{1}, {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "sty", Feb: "lut", Mar: "mar", Apr: "kwi", May: "maj", Jun: "cze", Jul: "lip", Aug: "sie", Sep: "wrz", Oct: "paź", Nov: "lis", Dec: "gru"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "s", Feb: "l", Mar: "m", Apr: "k", May: "m", Jun: "c", Jul: "l", Aug: "s", Sep: "w", Oct: "p", Nov: "l", Dec: "g"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "styczeń", Feb: "luty", Mar: "marzec", Apr: "kwiecień", May: "maj", Jun: "czerwiec", Jul: "lipiec", Aug: "sierpień", Sep: "wrzesień", Oct: "październik", Nov: "listopad", Dec: "grudzień"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "niedz.", Mon: "pon.", Tue: "wt.", Wed: "śr.", Thu: "czw.", Fri: "pt.", Sat: "sob."},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "N", Mon: "P", Tue: "W", Wed: "Ś", Thu: "C", Fri: "P", Sat: "S"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "niedz.", Mon: "pon.", Tue: "wt.", Wed: "śr.", Thu: "czw.", Fri: "pt.", Sat: "sob."},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "niedziela", Mon: "poniedziałek", Tue: "wtorek", Wed: "środa", Thu: "czwartek", Fri: "piątek", Sat: "sobota"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "a", PM: "p"},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"},
		},
	},
}
