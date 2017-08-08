package sl

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, dd. MMMM y", Long: "dd. MMMM y", Medium: "d. MMM y", Short: "d. MM. yy"},
		Time:     cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
		DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "jan", Feb: "feb", Mar: "mar", Apr: "apr", May: "maj", Jun: "jun", Jul: "jul", Aug: "avg", Sep: "sep", Oct: "okt", Nov: "nov", Dec: "dec"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "j", Feb: "f", Mar: "m", Apr: "a", May: "m", Jun: "j", Jul: "j", Aug: "a", Sep: "s", Oct: "o", Nov: "n", Dec: "d"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "januar", Feb: "februar", Mar: "marec", Apr: "april", May: "maj", Jun: "junij", Jul: "julij", Aug: "avgust", Sep: "september", Oct: "oktober", Nov: "november", Dec: "december"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "ned", Mon: "pon", Tue: "tor", Wed: "sre", Thu: "čet", Fri: "pet", Sat: "sob"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "n", Mon: "p", Tue: "t", Wed: "s", Thu: "č", Fri: "p", Sat: "s"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "ned.", Mon: "pon.", Tue: "tor.", Wed: "sre.", Thu: "čet.", Fri: "pet.", Sat: "sob."},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "nedelja", Mon: "ponedeljek", Tue: "torek", Wed: "sreda", Thu: "četrtek", Fri: "petek", Sat: "sobota"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "d", PM: "p"},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "dop.", PM: "pop."},
		},
	},
}
