package ur

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE، d MMMM، y", Long: "d MMMM، y", Medium: "d MMM، y", Short: "d/M/yy"},
		Time:     cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
		DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "جنوری", Feb: "فروری", Mar: "مارچ", Apr: "اپریل", May: "مئی", Jun: "جون", Jul: "جولائی", Aug: "اگست", Sep: "ستمبر", Oct: "اکتوبر", Nov: "نومبر", Dec: "دسمبر"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "جنوری", Feb: "فروری", Mar: "مارچ", Apr: "اپریل", May: "مئی", Jun: "جون", Jul: "جولائی", Aug: "اگست", Sep: "ستمبر", Oct: "اکتوبر", Nov: "نومبر", Dec: "دسمبر"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "اتوار", Mon: "سوموار", Tue: "منگل", Wed: "بدھ", Thu: "جمعرات", Fri: "جمعہ", Sat: "ہفتہ"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "S", Mon: "M", Tue: "T", Wed: "W", Thu: "T", Fri: "F", Sat: "S"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "اتوار", Mon: "سوموار", Tue: "منگل", Wed: "بدھ", Thu: "جمعرات", Fri: "جمعہ", Sat: "ہفتہ"},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "اتوار", Mon: "سوموار", Tue: "منگل", Wed: "بدھ", Thu: "جمعرات", Fri: "جمعہ", Sat: "ہفتہ"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "ق.د.", PM: "ب.د."},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "a", PM: "p"},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "قبل دوپہر", PM: "بعد دوپہر"},
		},
	},
}
