package is

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, d. MMMM y", Long: "d. MMMM y", Medium: "d. MMM y", Short: "d.M.y"},
		Time:     cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
		DateTime: cldr.CalendarDateFormat{Full: "{1} 'kl.' {0}", Long: "{1} 'kl.' {0}", Medium: "{1}, {0}", Short: "{1}, {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "jan.", Feb: "feb.", Mar: "mar.", Apr: "apr.", May: "maí", Jun: "jún.", Jul: "júl.", Aug: "ágú.", Sep: "sep.", Oct: "okt.", Nov: "nóv.", Dec: "des."},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "Á", Sep: "S", Oct: "O", Nov: "N", Dec: "D"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "janúar", Feb: "febrúar", Mar: "mars", Apr: "apríl", May: "maí", Jun: "júní", Jul: "júlí", Aug: "ágúst", Sep: "september", Oct: "október", Nov: "nóvember", Dec: "desember"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "sun.", Mon: "mán.", Tue: "þri.", Wed: "mið.", Thu: "fim.", Fri: "fös.", Sat: "lau."},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "S", Mon: "M", Tue: "Þ", Wed: "M", Thu: "F", Fri: "F", Sat: "L"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "su.", Mon: "má.", Tue: "þr.", Wed: "mi.", Thu: "fi.", Fri: "fö.", Sat: "la."},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "sunnudagur", Mon: "mánudagur", Tue: "þriðjudagur", Wed: "miðvikudagur", Thu: "fimmtudagur", Fri: "föstudagur", Sat: "laugardagur"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "f.", PM: "e."},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "árdegi", PM: "síðdegi"},
		},
	},
}
