package ar_TN

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "جانفي", Feb: "فيفري", Mar: "مارس", Apr: "أفريل", May: "ماي", Jun: "جوان", Jul: "جويلية", Aug: "أوت", Sep: "سبتمبر", Oct: "أكتوبر", Nov: "نوفمبر", Dec: "ديسمبر"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "ج", Feb: "ف", Mar: "م", Apr: "أ", May: "م", Jun: "ج", Jul: "ج", Aug: "أ", Sep: "س", Oct: "أ", Nov: "ن", Dec: "د"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "جانفي", Feb: "فيفري", Mar: "مارس", Apr: "أفريل", May: "ماي", Jun: "جوان", Jul: "جويلية", Aug: "أوت", Sep: "سبتمبر", Oct: "أكتوبر", Nov: "نوفمبر", Dec: "ديسمبر"},
		},
		Days:    cldr.CalendarDayFormatNames{},
		Periods: cldr.CalendarPeriodFormatNames{},
	},
}
