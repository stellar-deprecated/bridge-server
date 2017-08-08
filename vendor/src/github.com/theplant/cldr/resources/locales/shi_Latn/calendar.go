package shi_Latn

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM, y", Short: "d/M/y"},
		Time:     cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "inn", Feb: "bṛa", Mar: "maṛ", Apr: "ibr", May: "may", Jun: "yun", Jul: "yul", Aug: "ɣuc", Sep: "cut", Oct: "ktu", Nov: "nuw", Dec: "duj"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "i", Feb: "b", Mar: "m", Apr: "i", May: "m", Jun: "y", Jul: "y", Aug: "ɣ", Sep: "c", Oct: "k", Nov: "n", Dec: "d"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "innayr", Feb: "bṛayṛ", Mar: "maṛṣ", Apr: "ibrir", May: "mayyu", Jun: "yunyu", Jul: "yulyuz", Aug: "ɣuct", Sep: "cutanbir", Oct: "ktubr", Nov: "nuwanbir", Dec: "dujanbir"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "asa", Mon: "ayn", Tue: "asi", Wed: "akṛ", Thu: "akw", Fri: "asim", Sat: "asiḍ"},
			Narrow:      cldr.CalendarDayFormatNameValue{},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "asamas", Mon: "aynas", Tue: "asinas", Wed: "akṛas", Thu: "akwas", Fri: "asimwas", Sat: "asiḍyas"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "tifawt", PM: "tadggʷat"},
		},
	},
}
