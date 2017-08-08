package ar_LB

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "كانون الثاني", Feb: "شباط", Mar: "آذار", Apr: "نيسان", May: "أيار", Jun: "حزيران", Jul: "تموز", Aug: "آب", Sep: "أيلول", Oct: "تشرين الأول", Nov: "تشرين الثاني", Dec: "كانون الأول"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "ك", Feb: "ش", Mar: "آ", Apr: "ن", May: "أ", Jun: "ح", Jul: "ت", Aug: "آ", Sep: "أ", Oct: "ت", Nov: "ت", Dec: "ك"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "كانون الثاني", Feb: "شباط", Mar: "آذار", Apr: "نيسان", May: "أيار", Jun: "حزيران", Jul: "تموز", Aug: "آب", Sep: "أيلول", Oct: "تشرين الأول", Nov: "تشرين الثاني", Dec: "كانون الأول"},
		},
		Days:    cldr.CalendarDayFormatNames{},
		Periods: cldr.CalendarPeriodFormatNames{},
	},
}
