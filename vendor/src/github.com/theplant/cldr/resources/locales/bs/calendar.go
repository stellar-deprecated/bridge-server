package bs

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, dd. MMMM y.", Long: "dd. MMMM y.", Medium: "dd. MMM. y.", Short: "dd.MM.yy."},
		Time:     cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
		DateTime: cldr.CalendarDateFormat{Full: "{1} 'u' {0}", Long: "{1} 'u' {0}", Medium: "{1} {0}", Short: "{1} {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "jan", Feb: "feb", Mar: "mar", Apr: "apr", May: "maj", Jun: "jun", Jul: "jul", Aug: "aug", Sep: "sep", Oct: "okt", Nov: "nov", Dec: "dec"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "j", Feb: "f", Mar: "m", Apr: "a", May: "m", Jun: "j", Jul: "j", Aug: "a", Sep: "s", Oct: "o", Nov: "n", Dec: "d"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "januar", Feb: "februar", Mar: "mart", Apr: "april", May: "maj", Jun: "juni", Jul: "juli", Aug: "august", Sep: "septembar", Oct: "oktobar", Nov: "novembar", Dec: "decembar"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "ned", Mon: "pon", Tue: "uto", Wed: "sri", Thu: "čet", Fri: "pet", Sat: "sub"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "n", Mon: "p", Tue: "u", Wed: "s", Thu: "č", Fri: "p", Sat: "s"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "ned", Mon: "pon", Tue: "uto", Wed: "sri", Thu: "čet", Fri: "pet", Sat: "sub"},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "nedjelja", Mon: "ponedjeljak", Tue: "utorak", Wed: "srijeda", Thu: "četvrtak", Fri: "petak", Sat: "subota"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "prijepodne", PM: "popodne"},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "prije podne", PM: "popodne"},
		},
	},
}
