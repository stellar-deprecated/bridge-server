package ti_ER

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE፡ dd MMMM መዓልቲ y G", Long: "", Medium: "", Short: ""},
		Time:     cldr.CalendarDateFormat{},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "ጥሪ", Feb: "ለካቲ", Mar: "መጋቢ", Apr: "ሚያዝ", May: "ግንቦ", Jun: "ሰነ", Jul: "ሓምለ", Aug: "ነሓሰ", Sep: "መስከ", Oct: "ጥቅም", Nov: "ሕዳር", Dec: "ታሕሳ"},
			Narrow:      cldr.CalendarMonthFormatNameValue{},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "ጥሪ", Feb: "ለካቲት", Mar: "መጋቢት", Apr: "ሚያዝያ", May: "ግንቦት", Jun: "ሰነ", Jul: "ሓምለ", Aug: "ነሓሰ", Sep: "መስከረም", Oct: "ጥቅምቲ", Nov: "ሕዳር", Dec: "ታሕሳስ"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{},
			Narrow:      cldr.CalendarDayFormatNameValue{},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "ሰንበት", Mon: "ሰኑይ", Tue: "ሰሉስ", Wed: "ረቡዕ", Thu: "ሓሙስ", Fri: "ዓርቢ", Sat: "ቀዳም"},
		},
		Periods: cldr.CalendarPeriodFormatNames{},
	},
}
