package ar

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE، d MMMM، y", Long: "d MMMM، y", Medium: "dd\u200f/MM\u200f/y", Short: "d\u200f/M\u200f/y"},
		Time:     cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
		DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "يناير", Feb: "فبراير", Mar: "مارس", Apr: "أبريل", May: "مايو", Jun: "يونيو", Jul: "يوليو", Aug: "أغسطس", Sep: "سبتمبر", Oct: "أكتوبر", Nov: "نوفمبر", Dec: "ديسمبر"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "ي", Feb: "ف", Mar: "م", Apr: "أ", May: "و", Jun: "ن", Jul: "ل", Aug: "غ", Sep: "س", Oct: "ك", Nov: "ب", Dec: "د"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "يناير", Feb: "فبراير", Mar: "مارس", Apr: "أبريل", May: "مايو", Jun: "يونيو", Jul: "يوليو", Aug: "أغسطس", Sep: "سبتمبر", Oct: "أكتوبر", Nov: "نوفمبر", Dec: "ديسمبر"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "الأحد", Mon: "الاثنين", Tue: "الثلاثاء", Wed: "الأربعاء", Thu: "الخميس", Fri: "الجمعة", Sat: "السبت"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "ح", Mon: "ن", Tue: "ث", Wed: "ر", Thu: "خ", Fri: "ج", Sat: "س"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "الأحد", Mon: "الاثنين", Tue: "الثلاثاء", Wed: "الأربعاء", Thu: "الخميس", Fri: "الجمعة", Sat: "السبت"},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "الأحد", Mon: "الاثنين", Tue: "الثلاثاء", Wed: "الأربعاء", Thu: "الخميس", Fri: "الجمعة", Sat: "السبت"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "ص", PM: "م"},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "ص", PM: "م"},
		},
	},
}
