package kl

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE dd MMMM y", Long: "dd MMMM y", Medium: "MMM dd, y", Short: "y-MM-dd"},
		Time:     cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
		DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "jan", Feb: "feb", Mar: "mar", Apr: "apr", May: "maj", Jun: "jun", Jul: "jul", Aug: "aug", Sep: "sep", Oct: "okt", Nov: "nov", Dec: "dec"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "januari", Feb: "februari", Mar: "martsi", Apr: "aprili", May: "maji", Jun: "juni", Jul: "juli", Aug: "augustusi", Sep: "septemberi", Oct: "oktoberi", Nov: "novemberi", Dec: "decemberi"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "sab", Mon: "ata", Tue: "mar", Wed: "pin", Thu: "sis", Fri: "tal", Sat: "arf"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "S", Mon: "A", Tue: "M", Wed: "P", Thu: "S", Fri: "T", Sat: "A"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "sab", Mon: "ata", Tue: "mar", Wed: "pin", Thu: "sis", Fri: "tal", Sat: "arf"},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "sabaat", Mon: "ataasinngorneq", Tue: "marlunngorneq", Wed: "pingasunngorneq", Thu: "sisamanngorneq", Fri: "tallimanngorneq", Sat: "arfininngorneq"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "u.t.", PM: "u.k."},
			Narrow:      cldr.CalendarPeriodFormatNameValue{},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "ulloqeqqata-tungaa", PM: "ulloqeqqata-kingorna"},
		},
	},
}
