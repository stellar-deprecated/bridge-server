package ti

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE፣ dd MMMM መዓልቲ y G", Long: "dd MMMM y", Medium: "dd-MMM-y", Short: "dd/MM/yy"},
		Time:     cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "ጃንዩ", Feb: "ፌብሩ", Mar: "ማርች", Apr: "ኤፕረ", May: "ሜይ", Jun: "ጁን", Jul: "ጁላይ", Aug: "ኦገስ", Sep: "ሴፕቴ", Oct: "ኦክተ", Nov: "ኖቬም", Dec: "ዲሴም"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "ጃ", Feb: "ፌ", Mar: "ማ", Apr: "ኤ", May: "ሜ", Jun: "ጁ", Jul: "ጁ", Aug: "ኦ", Sep: "ሴ", Oct: "ኦ", Nov: "ኖ", Dec: "ዲ"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "ጃንዩወሪ", Feb: "ፌብሩወሪ", Mar: "ማርች", Apr: "ኤፕረል", May: "ሜይ", Jun: "ጁን", Jul: "ጁላይ", Aug: "ኦገስት", Sep: "ሴፕቴምበር", Oct: "ኦክተውበር", Nov: "ኖቬምበር", Dec: "ዲሴምበር"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "ሰ", Mon: "ሰ", Tue: "ሠ", Wed: "ረ", Thu: "ኃ", Fri: "ዓ", Sat: "ቀ"},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "ሰንበት", Mon: "ሰኑይ", Tue: "ሠሉስ", Wed: "ረቡዕ", Thu: "ኃሙስ", Fri: "ዓርቢ", Sat: "ቀዳም"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "ንጉሆ ሰዓተ", PM: "ድሕር ሰዓት"},
		},
	},
}
