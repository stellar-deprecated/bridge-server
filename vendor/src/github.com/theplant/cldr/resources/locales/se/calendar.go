package se

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "ođđj", Feb: "guov", Mar: "njuk", Apr: "cuo", May: "mies", Jun: "geas", Jul: "suoi", Aug: "borg", Sep: "čakč", Oct: "golg", Nov: "skáb", Dec: "juov"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "O", Feb: "G", Mar: "N", Apr: "C", May: "M", Jun: "G", Jul: "S", Aug: "B", Sep: "Č", Oct: "G", Nov: "S", Dec: "J"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "ođđajagemánnu", Feb: "guovvamánnu", Mar: "njukčamánnu", Apr: "cuoŋománnu", May: "miessemánnu", Jun: "geassemánnu", Jul: "suoidnemánnu", Aug: "borgemánnu", Sep: "čakčamánnu", Oct: "golggotmánnu", Nov: "skábmamánnu", Dec: "juovlamánnu"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "sotn", Mon: "vuos", Tue: "maŋ", Wed: "gask", Thu: "duor", Fri: "bear", Sat: "láv"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "S", Mon: "V", Tue: "M", Wed: "G", Thu: "D", Fri: "B", Sat: "L"},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "sotnabeaivi", Mon: "vuossárga", Tue: "maŋŋebárga", Wed: "gaskavahkku", Thu: "duorasdat", Fri: "bearjadat", Sat: "lávvardat"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "i.b.", PM: "e.b."},
			Narrow:      cldr.CalendarPeriodFormatNameValue{},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "iđitbeaivi", PM: "eahketbeaivi"},
		},
	},
}
