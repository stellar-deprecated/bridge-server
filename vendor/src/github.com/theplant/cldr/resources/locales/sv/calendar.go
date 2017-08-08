package sv

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "y-MM-dd"},
		Time:     cldr.CalendarDateFormat{Full: "'kl'. HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
		DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Jan.", Feb: "Feb.", Mar: "Mars", Apr: "Apr.", May: "Maj", Jun: "Juni", Jul: "Juli", Aug: "Aug.", Sep: "Sep.", Oct: "Okt.", Nov: "Nov.", Dec: "Dec."},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Januari", Feb: "Februari", Mar: "Mars", Apr: "April", May: "Maj", Jun: "Juni", Jul: "Juli", Aug: "Augusti", Sep: "September", Oct: "Oktober", Nov: "November", Dec: "December"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Sön", Mon: "Mån", Tue: "Tis", Wed: "Ons", Thu: "Tor", Fri: "Fre", Sat: "Lör"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "S", Mon: "M", Tue: "T", Wed: "O", Thu: "T", Fri: "F", Sat: "L"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "Sö", Mon: "Må", Tue: "Ti", Wed: "On", Thu: "To", Fri: "Fr", Sat: "Lö"},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "Söndag", Mon: "Måndag", Tue: "Tisdag", Wed: "Onsdag", Thu: "Torsdag", Fri: "Fredag", Sat: "Lördag"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "fm", PM: "em"},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "f.m.", PM: "e.m."},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "förmiddag", PM: "eftermiddag"},
		},
	},
}
