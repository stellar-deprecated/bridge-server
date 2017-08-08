package et

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, d. MMMM y", Long: "d. MMMM y", Medium: "d. MMM y", Short: "dd.MM.yy"},
		Time:     cldr.CalendarDateFormat{Full: "H:mm.ss zzzz", Long: "H:mm.ss z", Medium: "H:mm.ss", Short: "H:mm"},
		DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "jaan", Feb: "veebr", Mar: "märts", Apr: "apr", May: "mai", Jun: "juuni", Jul: "juuli", Aug: "aug", Sep: "sept", Oct: "okt", Nov: "nov", Dec: "dets"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "V", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "jaanuar", Feb: "veebruar", Mar: "märts", Apr: "aprill", May: "mai", Jun: "juuni", Jul: "juuli", Aug: "august", Sep: "september", Oct: "oktoober", Nov: "november", Dec: "detsember"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "P", Mon: "E", Tue: "T", Wed: "K", Thu: "N", Fri: "R", Sat: "L"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "P", Mon: "E", Tue: "T", Wed: "K", Thu: "N", Fri: "R", Sat: "L"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "P", Mon: "E", Tue: "T", Wed: "K", Thu: "N", Fri: "R", Sat: "L"},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "pühapäev", Mon: "esmaspäev", Tue: "teisipäev", Wed: "kolmapäev", Thu: "neljapäev", Fri: "reede", Sat: "laupäev"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "e.k.", PM: "p.k."},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "a", PM: "p"},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"},
		},
	},
}
