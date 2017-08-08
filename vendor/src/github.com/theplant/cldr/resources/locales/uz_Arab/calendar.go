package uz_Arab

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "y نچی ییل d نچی MMMM EEEE کونی", Long: "d نچی MMMM y", Medium: "d MMM y", Short: "y/M/d"},
		Time:     cldr.CalendarDateFormat{Full: "H:mm:ss (zzzz)", Long: "H:mm:ss (z)", Medium: "H:mm:ss", Short: "H:mm"},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "جنو", Feb: "فبر", Mar: "مار", Apr: "اپر", May: "مـی", Jun: "جون", Jul: "جول", Aug: "اگس", Sep: "سپت", Oct: "اکت", Nov: "نوم", Dec: "دسم"},
			Narrow:      cldr.CalendarMonthFormatNameValue{},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "جنوری", Feb: "فبروری", Mar: "مارچ", Apr: "اپریل", May: "می", Jun: "جون", Jul: "جولای", Aug: "اگست", Sep: "سپتمبر", Oct: "اکتوبر", Nov: "نومبر", Dec: "دسمبر"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "ی.", Mon: "د.", Tue: "س.", Wed: "چ.", Thu: "پ.", Fri: "ج.", Sat: "ش."},
			Narrow:      cldr.CalendarDayFormatNameValue{},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "یکشنبه", Mon: "دوشنبه", Tue: "سه\u200cشنبه", Wed: "چهارشنبه", Thu: "پنجشنبه", Fri: "جمعه", Sat: "شنبه"},
		},
		Periods: cldr.CalendarPeriodFormatNames{},
	},
}
