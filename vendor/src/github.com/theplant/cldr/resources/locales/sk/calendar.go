package sk

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, d. MMMM y", Long: "d. MMMM y", Medium: "d. M. y", Short: "dd.MM.yy"},
		Time:     cldr.CalendarDateFormat{Full: "H:mm:ss zzzz", Long: "H:mm:ss z", Medium: "H:mm:ss", Short: "H:mm"},
		DateTime: cldr.CalendarDateFormat{Full: "{1}, {0}", Long: "{1}, {0}", Medium: "{1}, {0}", Short: "{1} {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "jan", Feb: "feb", Mar: "mar", Apr: "apr", May: "máj", Jun: "jún", Jul: "júl", Aug: "aug", Sep: "sep", Oct: "okt", Nov: "nov", Dec: "dec"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "j", Feb: "f", Mar: "m", Apr: "a", May: "m", Jun: "j", Jul: "j", Aug: "a", Sep: "s", Oct: "o", Nov: "n", Dec: "d"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "január", Feb: "február", Mar: "marec", Apr: "apríl", May: "máj", Jun: "jún", Jul: "júl", Aug: "august", Sep: "september", Oct: "október", Nov: "november", Dec: "december"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "ne", Mon: "po", Tue: "ut", Wed: "st", Thu: "št", Fri: "pi", Sat: "so"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "N", Mon: "P", Tue: "U", Wed: "S", Thu: "Š", Fri: "P", Sat: "S"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "Ne", Mon: "Po", Tue: "Ut", Wed: "St", Thu: "Št", Fri: "Pi", Sat: "So"},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "nedeľa", Mon: "pondelok", Tue: "utorok", Wed: "streda", Thu: "štvrtok", Fri: "piatok", Sat: "sobota"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "dop.", PM: "odp."},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "dopoludnia", PM: "odpoludnia"},
		},
	},
}
