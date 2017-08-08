package ar_MR

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "يناير", Feb: "فبراير", Mar: "مارس", Apr: "إبريل", May: "مايو", Jun: "يونيو", Jul: "يوليو", Aug: "أغشت", Sep: "شتمبر", Oct: "أكتوبر", Nov: "نوفمبر", Dec: "دجمبر"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "إ", May: "", Jun: "", Jul: "", Aug: "", Sep: "ش", Oct: "", Nov: "", Dec: ""},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "يناير", Feb: "فبراير", Mar: "مارس", Apr: "إبريل", May: "مايو", Jun: "يونيو", Jul: "يوليو", Aug: "أغشت", Sep: "شتمبر", Oct: "أكتوبر", Nov: "نوفمبر", Dec: "دجمبر"},
		},
		Days:    cldr.CalendarDayFormatNames{},
		Periods: cldr.CalendarPeriodFormatNames{},
	},
}
