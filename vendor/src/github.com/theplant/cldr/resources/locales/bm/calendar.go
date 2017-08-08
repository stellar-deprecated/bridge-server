package bm

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM, y", Short: "d/M/y"},
		Time:     cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "zan", Feb: "feb", Mar: "mar", Apr: "awi", May: "mɛ", Jun: "zuw", Jul: "zul", Aug: "uti", Sep: "sɛt", Oct: "ɔku", Nov: "now", Dec: "des"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "Z", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "Z", Jul: "Z", Aug: "U", Sep: "S", Oct: "Ɔ", Nov: "N", Dec: "D"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "zanwuye", Feb: "feburuye", Mar: "marisi", Apr: "awirili", May: "mɛ", Jun: "zuwɛn", Jul: "zuluye", Aug: "uti", Sep: "sɛtanburu", Oct: "ɔkutɔburu", Nov: "nowanburu", Dec: "desanburu"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "kar", Mon: "ntɛ", Tue: "tar", Wed: "ara", Thu: "ala", Fri: "jum", Sat: "sib"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "K", Mon: "N", Tue: "T", Wed: "A", Thu: "A", Fri: "J", Sat: "S"},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "kari", Mon: "ntɛnɛ", Tue: "tarata", Wed: "araba", Thu: "alamisa", Fri: "juma", Sat: "sibiri"},
		},
		Periods: cldr.CalendarPeriodFormatNames{},
	},
}
