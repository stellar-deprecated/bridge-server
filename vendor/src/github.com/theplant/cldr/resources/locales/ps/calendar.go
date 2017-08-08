package ps

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE د y د MMMM d", Long: "د y د MMMM d", Medium: "d MMM y", Short: "y/M/d"},
		Time:     cldr.CalendarDateFormat{Full: "H:mm:ss (zzzz)", Long: "H:mm:ss (z)", Medium: "H:mm:ss", Short: "H:mm"},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{},
			Narrow:      cldr.CalendarMonthFormatNameValue{},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "جنوري", Feb: "فبروري", Mar: "مارچ", Apr: "اپریل", May: "می", Jun: "جون", Jul: "جولای", Aug: "اګست", Sep: "سپتمبر", Oct: "اکتوبر", Nov: "نومبر", Dec: "دسمبر"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{},
			Narrow:      cldr.CalendarDayFormatNameValue{},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "یکشنبه", Mon: "دوشنبه", Tue: "سه\u200cشنبه", Wed: "چهارشنبه", Thu: "پنجشنبه", Fri: "جمعه", Sat: "شنبه"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "غ.م.", PM: "غ.و."},
		},
	},
}
