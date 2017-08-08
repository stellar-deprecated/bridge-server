package ar_MA

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "يناير", Feb: "فبراير", Mar: "مارس", Apr: "أبريل", May: "ماي", Jun: "يونيو", Jul: "يوليوز", Aug: "غشت", Sep: "شتنبر", Oct: "أكتوبر", Nov: "نونبر", Dec: "دجنبر"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "ي", Feb: "ف", Mar: "م", Apr: "أ", May: "م", Jun: "ن", Jul: "ل", Aug: "غ", Sep: "ش", Oct: "ك", Nov: "ب", Dec: "د"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "يناير", Feb: "فبراير", Mar: "مارس", Apr: "أبريل", May: "ماي", Jun: "يونيو", Jul: "يوليوز", Aug: "غشت", Sep: "شتنبر", Oct: "أكتوبر", Nov: "نونبر", Dec: "دجنبر"},
		},
		Days:    cldr.CalendarDayFormatNames{},
		Periods: cldr.CalendarPeriodFormatNames{},
	},
}
