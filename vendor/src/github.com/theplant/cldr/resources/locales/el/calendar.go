package el

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "d/M/yy"},
		Time:     cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
		DateTime: cldr.CalendarDateFormat{Full: "{1} - {0}", Long: "{1} - {0}", Medium: "{1}, {0}", Short: "{1}, {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Ιαν", Feb: "Φεβ", Mar: "Μάρ", Apr: "Απρ", May: "Μάι", Jun: "Ιούν", Jul: "Ιούλ", Aug: "Αύγ", Sep: "Σεπ", Oct: "Οκτ", Nov: "Νοέ", Dec: "Δεκ"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "Ι", Feb: "Φ", Mar: "Μ", Apr: "Α", May: "Μ", Jun: "Ι", Jul: "Ι", Aug: "Α", Sep: "Σ", Oct: "Ο", Nov: "Ν", Dec: "Δ"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Ιανουάριος", Feb: "Φεβρουάριος", Mar: "Μάρτιος", Apr: "Απρίλιος", May: "Μάιος", Jun: "Ιούνιος", Jul: "Ιούλιος", Aug: "Αύγουστος", Sep: "Σεπτέμβριος", Oct: "Οκτώβριος", Nov: "Νοέμβριος", Dec: "Δεκέμβριος"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Κυρ", Mon: "Δευ", Tue: "Τρί", Wed: "Τετ", Thu: "Πέμ", Fri: "Παρ", Sat: "Σάβ"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "Κ", Mon: "Δ", Tue: "Τ", Wed: "Τ", Thu: "Π", Fri: "Π", Sat: "Σ"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "Κυ", Mon: "Δε", Tue: "Τρ", Wed: "Τε", Thu: "Πέ", Fri: "Πα", Sat: "Σά"},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "Κυριακή", Mon: "Δευτέρα", Tue: "Τρίτη", Wed: "Τετάρτη", Thu: "Πέμπτη", Fri: "Παρασκευή", Sat: "Σάββατο"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "π", PM: "μ"},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "π.μ.", PM: "μ.μ."},
		},
	},
}
