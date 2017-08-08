package ewo

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "d/M/y"},
		Time:     cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "ngo", Feb: "ngb", Mar: "ngl", Apr: "ngn", May: "ngt", Jun: "ngs", Jul: "ngz", Aug: "ngm", Sep: "nge", Oct: "nga", Nov: "ngad", Dec: "ngab"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "o", Feb: "b", Mar: "l", Apr: "n", May: "t", Jun: "s", Jul: "z", Aug: "m", Sep: "e", Oct: "a", Nov: "d", Dec: "b"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "ngɔn osú", Feb: "ngɔn bɛ̌", Mar: "ngɔn lála", Apr: "ngɔn nyina", May: "ngɔn tána", Jun: "ngɔn saməna", Jul: "ngɔn zamgbála", Aug: "ngɔn mwom", Sep: "ngɔn ebulú", Oct: "ngɔn awóm", Nov: "ngɔn awóm ai dziá", Dec: "ngɔn awóm ai bɛ̌"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "sɔ́n", Mon: "mɔ́n", Tue: "smb", Wed: "sml", Thu: "smn", Fri: "fúl", Sat: "sér"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "s", Mon: "m", Tue: "s", Wed: "s", Thu: "s", Fri: "f", Sat: "s"},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "sɔ́ndɔ", Mon: "mɔ́ndi", Tue: "sɔ́ndɔ məlú mə́bɛ̌", Wed: "sɔ́ndɔ məlú mə́lɛ́", Thu: "sɔ́ndɔ məlú mə́nyi", Fri: "fúladé", Sat: "séradé"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "kíkíríg", PM: "ngəgógəle"},
		},
	},
}
