package cs

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE d. MMMM y", Long: "d. MMMM y", Medium: "d. M. y", Short: "dd.MM.yy"},
		Time:     cldr.CalendarDateFormat{Full: "H:mm:ss zzzz", Long: "H:mm:ss z", Medium: "H:mm:ss", Short: "H:mm"},
		DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "led", Feb: "úno", Mar: "bře", Apr: "dub", May: "kvě", Jun: "čvn", Jul: "čvc", Aug: "srp", Sep: "zář", Oct: "říj", Nov: "lis", Dec: "pro"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "1", Feb: "2", Mar: "3", Apr: "4", May: "5", Jun: "6", Jul: "7", Aug: "8", Sep: "9", Oct: "10", Nov: "11", Dec: "12"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "leden", Feb: "únor", Mar: "březen", Apr: "duben", May: "květen", Jun: "červen", Jul: "červenec", Aug: "srpen", Sep: "září", Oct: "říjen", Nov: "listopad", Dec: "prosinec"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "ne", Mon: "po", Tue: "út", Wed: "st", Thu: "čt", Fri: "pá", Sat: "so"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "N", Mon: "P", Tue: "Ú", Wed: "S", Thu: "Č", Fri: "P", Sat: "S"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "ne", Mon: "po", Tue: "út", Wed: "st", Thu: "čt", Fri: "pá", Sat: "so"},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "neděle", Mon: "pondělí", Tue: "úterý", Wed: "středa", Thu: "čtvrtek", Fri: "pátek", Sat: "sobota"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "dop.", PM: "odp."},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "dop.", PM: "odp."},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "dopoledne", PM: "odpoledne"},
		},
	},
}
