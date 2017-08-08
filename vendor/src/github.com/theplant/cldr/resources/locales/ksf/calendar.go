package ksf

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "d/M/y"},
		Time:     cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "ŋ1", Feb: "ŋ2", Mar: "ŋ3", Apr: "ŋ4", May: "ŋ5", Jun: "ŋ6", Jul: "ŋ7", Aug: "ŋ8", Sep: "ŋ9", Oct: "ŋ10", Nov: "ŋ11", Dec: "ŋ12"},
			Narrow:      cldr.CalendarMonthFormatNameValue{},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "ŋwíí a ntɔ́ntɔ", Feb: "ŋwíí akǝ bɛ́ɛ", Mar: "ŋwíí akǝ ráá", Apr: "ŋwíí akǝ nin", May: "ŋwíí akǝ táan", Jun: "ŋwíí akǝ táafɔk", Jul: "ŋwíí akǝ táabɛɛ", Aug: "ŋwíí akǝ táaraa", Sep: "ŋwíí akǝ táanin", Oct: "ŋwíí akǝ ntɛk", Nov: "ŋwíí akǝ ntɛk di bɔ́k", Dec: "ŋwíí akǝ ntɛk di bɛ́ɛ"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "sɔ́n", Mon: "lǝn", Tue: "maa", Wed: "mɛk", Thu: "jǝǝ", Fri: "júm", Sat: "sam"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "s", Mon: "l", Tue: "m", Wed: "m", Thu: "j", Fri: "j", Sat: "s"},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "sɔ́ndǝ", Mon: "lǝndí", Tue: "maadí", Wed: "mɛkrɛdí", Thu: "jǝǝdí", Fri: "júmbá", Sat: "samdí"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "sárúwá", PM: "cɛɛ́nko"},
		},
	},
}
