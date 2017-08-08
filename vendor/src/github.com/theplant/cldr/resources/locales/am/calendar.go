package am

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"},
		Time:     cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
		DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "ጃንዩ", Feb: "ፌብሩ", Mar: "ማርች", Apr: "ኤፕሪ", May: "ሜይ", Jun: "ጁን", Jul: "ጁላይ", Aug: "ኦገስ", Sep: "ሴፕቴ", Oct: "ኦክቶ", Nov: "ኖቬም", Dec: "ዲሴም"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "ጃ", Feb: "ፌ", Mar: "ማ", Apr: "ኤ", May: "ሜ", Jun: "ጁ", Jul: "ጁ", Aug: "ኦ", Sep: "ሴ", Oct: "ኦ", Nov: "ኖ", Dec: "ዲ"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "ጃንዩወሪ", Feb: "ፌብሩወሪ", Mar: "ማርች", Apr: "ኤፕሪል", May: "ሜይ", Jun: "ጁን", Jul: "ጁላይ", Aug: "ኦገስት", Sep: "ሴፕቴምበር", Oct: "ኦክቶበር", Nov: "ኖቬምበር", Dec: "ዲሴምበር"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "እሑድ", Mon: "ሰኞ", Tue: "ማክሰ", Wed: "ረቡዕ", Thu: "ሐሙስ", Fri: "ዓርብ", Sat: "ቅዳሜ"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "እ", Mon: "ሰ", Tue: "ማ", Wed: "ረ", Thu: "ሐ", Fri: "ዓ", Sat: "ቅ"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "እ", Mon: "ሰ", Tue: "ማ", Wed: "ረ", Thu: "ሐ", Fri: "ዓ", Sat: "ቅ"},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "እሑድ", Mon: "ሰኞ", Tue: "ማክሰኞ", Wed: "ረቡዕ", Thu: "ሐሙስ", Fri: "ዓርብ", Sat: "ቅዳሜ"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "ጠ", PM: "ከ"},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "ጥዋት", PM: "ከሰዓት"},
		},
	},
}
