package fa_AF

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "جنو", Feb: "", Mar: "", Apr: "", May: "مـی", Jun: "", Jul: "جول", Aug: "", Sep: "", Oct: "", Nov: "", Dec: "دسم"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "ج", Feb: "ف", Mar: "م", Apr: "ا", May: "م", Jun: "ج", Jul: "ج", Aug: "ا", Sep: "س", Oct: "ا", Nov: "ن", Dec: "د"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "جنوری", Feb: "فبروری", Mar: "مارچ", Apr: "اپریل", May: "می", Jun: "جون", Jul: "جولای", Aug: "اگست", Sep: "سپتمبر", Oct: "اکتوبر", Nov: "نومبر", Dec: "دسمبر"},
		},
		Days:    cldr.CalendarDayFormatNames{},
		Periods: cldr.CalendarPeriodFormatNames{},
	},
}
