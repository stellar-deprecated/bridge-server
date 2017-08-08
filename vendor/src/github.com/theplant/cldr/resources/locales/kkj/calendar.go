package kkj

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE dd MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM y"},
		Time:     cldr.CalendarDateFormat{Full: "", Long: "", Medium: "HH:mm:ss", Short: "HH:mm"},
		DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{},
			Narrow:      cldr.CalendarMonthFormatNameValue{},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "pamba", Feb: "wanja", Mar: "mbiyɔ mɛndoŋgɔ", Apr: "Nyɔlɔmbɔŋgɔ", May: "Mɔnɔ ŋgbanja", Jun: "Nyaŋgwɛ ŋgbanja", Jul: "kuŋgwɛ", Aug: "fɛ", Sep: "njapi", Oct: "nyukul", Nov: "11", Dec: "ɓulɓusɛ"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "sɔndi", Mon: "lundi", Tue: "mardi", Wed: "mɛrkɛrɛdi", Thu: "yedi", Fri: "vaŋdɛrɛdi", Sat: "mɔnɔ sɔndi"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "so", Mon: "lu", Tue: "ma", Wed: "mɛ", Thu: "ye", Fri: "va", Sat: "ms"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "so", Mon: "lu", Tue: "ma", Wed: "mɛ", Thu: "ye", Fri: "va", Sat: "ms"},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "sɔndi", Mon: "lundi", Tue: "mardi", Wed: "mɛrkɛrɛdi", Thu: "yedi", Fri: "vaŋdɛrɛdi", Sat: "mɔnɔ sɔndi"},
		},
		Periods: cldr.CalendarPeriodFormatNames{},
	},
}
