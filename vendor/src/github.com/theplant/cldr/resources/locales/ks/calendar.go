package ks

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, MMMM d, y", Long: "MMMM d, y", Medium: "MMM d, y", Short: "M/d/yy"},
		Time:     cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
		DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "ج", Feb: "ف", Mar: "م", Apr: "ا", May: "م", Jun: "ج", Jul: "ج", Aug: "ا", Sep: "س", Oct: "س", Nov: "ا", Dec: "ن"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "جنؤری", Feb: "فرؤری", Mar: "مارٕچ", Apr: "اپریل", May: "میٔ", Jun: "جوٗن", Jul: "جوٗلایی", Aug: "اگست", Sep: "ستمبر", Oct: "اکتوٗبر", Nov: "نومبر", Dec: "دسمبر"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "آتھوار", Mon: "ژٔنٛدٕروار", Tue: "بوٚموار", Wed: "بودوار", Thu: "برٛٮ۪سوار", Fri: "جُمہ", Sat: "بٹوار"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "ا", Mon: "ژ", Tue: "ب", Wed: "ب", Thu: "ب", Fri: "ج", Sat: "ب"},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "اَتھوار", Mon: "ژٔنٛدرٕروار", Tue: "بوٚموار", Wed: "بودوار", Thu: "برٛٮ۪سوار", Fri: "جُمہ", Sat: "بٹوار"},
		},
		Periods: cldr.CalendarPeriodFormatNames{},
	},
}
