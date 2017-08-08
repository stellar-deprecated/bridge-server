package hu

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "y. MMMM d., EEEE", Long: "y. MMMM d.", Medium: "y. MMM d.", Short: "y. MM. dd."},
		Time:     cldr.CalendarDateFormat{Full: "H:mm:ss zzzz", Long: "H:mm:ss z", Medium: "H:mm:ss", Short: "H:mm"},
		DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "jan.", Feb: "febr.", Mar: "márc.", Apr: "ápr.", May: "máj.", Jun: "jún.", Jul: "júl.", Aug: "aug.", Sep: "szept.", Oct: "okt.", Nov: "nov.", Dec: "dec."},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "Á", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "Sz", Oct: "O", Nov: "N", Dec: "D"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "január", Feb: "február", Mar: "március", Apr: "április", May: "május", Jun: "június", Jul: "július", Aug: "augusztus", Sep: "szeptember", Oct: "október", Nov: "november", Dec: "december"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "V", Mon: "H", Tue: "K", Wed: "Sze", Thu: "Cs", Fri: "P", Sat: "Szo"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "V", Mon: "H", Tue: "K", Wed: "Sz", Thu: "Cs", Fri: "P", Sat: "Sz"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "V", Mon: "H", Tue: "K", Wed: "Sze", Thu: "Cs", Fri: "P", Sat: "Szo"},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "vasárnap", Mon: "hétfő", Tue: "kedd", Wed: "szerda", Thu: "csütörtök", Fri: "péntek", Sat: "szombat"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "de.", PM: "du."},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "de.", PM: "du."},
		},
	},
}
