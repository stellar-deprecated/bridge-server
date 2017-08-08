package fa

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "y/M/d"},
		Time:     cldr.CalendarDateFormat{Full: "H:mm:ss (zzzz)", Long: "H:mm:ss (z)", Medium: "H:mm:ss", Short: "H:mm"},
		DateTime: cldr.CalendarDateFormat{Full: "{1}، ساعت {0}", Long: "{1}، ساعت {0}", Medium: "{1}،\u200f {0}", Short: "{1}،\u200f {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "ژانویه", Feb: "فوریه", Mar: "مارس", Apr: "آوریل", May: "مه", Jun: "ژوئن", Jul: "ژوئیه", Aug: "اوت", Sep: "سپتامبر", Oct: "اکتبر", Nov: "نوامبر", Dec: "دسامبر"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "ژ", Feb: "ف", Mar: "م", Apr: "آ", May: "م", Jun: "ژ", Jul: "ژ", Aug: "ا", Sep: "س", Oct: "ا", Nov: "ن", Dec: "د"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "ژانویه", Feb: "فوریه", Mar: "مارس", Apr: "آوریل", May: "مه", Jun: "ژوئن", Jul: "ژوئیه", Aug: "اوت", Sep: "سپتامبر", Oct: "اکتبر", Nov: "نوامبر", Dec: "دسامبر"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "یکشنبه", Mon: "دوشنبه", Tue: "سه\u200cشنبه", Wed: "چهارشنبه", Thu: "پنجشنبه", Fri: "جمعه", Sat: "شنبه"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "ی", Mon: "د", Tue: "س", Wed: "چ", Thu: "پ", Fri: "ج", Sat: "ش"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "۱ش", Mon: "۲ش", Tue: "۳ش", Wed: "۴ش", Thu: "۵ش", Fri: "ج", Sat: "ش"},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "یکشنبه", Mon: "دوشنبه", Tue: "سه\u200cشنبه", Wed: "چهارشنبه", Thu: "پنجشنبه", Fri: "جمعه", Sat: "شنبه"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "ق", PM: "ب"},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "قبل\u200cازظهر", PM: "بعدازظهر"},
		},
	},
}
