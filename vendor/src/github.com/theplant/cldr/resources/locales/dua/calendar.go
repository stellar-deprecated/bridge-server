package dua

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "d/M/y"},
		Time:     cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "di", Feb: "ŋgɔn", Mar: "sɔŋ", Apr: "diɓ", May: "emi", Jun: "esɔ", Jul: "mad", Aug: "diŋ", Sep: "nyɛt", Oct: "may", Nov: "tin", Dec: "elá"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "d", Feb: "ŋ", Mar: "s", Apr: "d", May: "e", Jun: "e", Jul: "m", Aug: "d", Sep: "n", Oct: "m", Nov: "t", Dec: "e"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "dimɔ́di", Feb: "ŋgɔndɛ", Mar: "sɔŋɛ", Apr: "diɓáɓá", May: "emiasele", Jun: "esɔpɛsɔpɛ", Jul: "madiɓɛ́díɓɛ́", Aug: "diŋgindi", Sep: "nyɛtɛki", Oct: "mayésɛ́", Nov: "tiníní", Dec: "eláŋgɛ́"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "ét", Mon: "mɔ́s", Tue: "kwa", Wed: "muk", Thu: "ŋgi", Fri: "ɗón", Sat: "esa"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "e", Mon: "m", Tue: "k", Wed: "m", Thu: "ŋ", Fri: "ɗ", Sat: "e"},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "éti", Mon: "mɔ́sú", Tue: "kwasú", Wed: "mukɔ́sú", Thu: "ŋgisú", Fri: "ɗónɛsú", Sat: "esaɓasú"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "idiɓa", PM: "ebyámu"},
		},
	},
}
