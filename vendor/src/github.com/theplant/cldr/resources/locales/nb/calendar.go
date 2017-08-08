package nb

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE d. MMMM y", Long: "d. MMMM y", Medium: "d. MMM y", Short: "dd.MM.y"},
		Time:     cldr.CalendarDateFormat{Full: "HH.mm.ss zzzz", Long: "HH.mm.ss z", Medium: "HH.mm.ss", Short: "HH.mm"},
		DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} 'kl.' {0}", Medium: "{1}, {0}", Short: "{1}, {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "jan", Feb: "feb", Mar: "mar", Apr: "apr", May: "mai", Jun: "jun", Jul: "jul", Aug: "aug", Sep: "sep", Oct: "okt", Nov: "nov", Dec: "des"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "januar", Feb: "februar", Mar: "mars", Apr: "april", May: "mai", Jun: "juni", Jul: "juli", Aug: "august", Sep: "september", Oct: "oktober", Nov: "november", Dec: "desember"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "sø.", Mon: "ma.", Tue: "ti.", Wed: "on.", Thu: "to.", Fri: "fr.", Sat: "lø."},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "S", Mon: "M", Tue: "T", Wed: "O", Thu: "T", Fri: "F", Sat: "L"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "sø.", Mon: "ma.", Tue: "ti.", Wed: "on.", Thu: "to.", Fri: "fr.", Sat: "lø."},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "søndag", Mon: "mandag", Tue: "tirsdag", Wed: "onsdag", Thu: "torsdag", Fri: "fredag", Sat: "lørdag"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "a.m.", PM: "p.m."},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "a", PM: "p"},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"},
		},
	},
}
