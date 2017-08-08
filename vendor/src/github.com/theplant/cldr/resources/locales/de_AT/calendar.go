package de_AT

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Jän", Feb: "Feb", Mar: "Mär", Apr: "Apr", May: "Mai", Jun: "Jun", Jul: "Jul", Aug: "Aug", Sep: "Sep", Oct: "Okt", Nov: "Nov", Dec: "Dez"},
			Narrow:      cldr.CalendarMonthFormatNameValue{},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Jänner", Feb: "Februar", Mar: "März", Apr: "April", May: "Mai", Jun: "Juni", Jul: "Juli", Aug: "August", Sep: "September", Oct: "Oktober", Nov: "November", Dec: "Dezember"},
		},
		Days:    cldr.CalendarDayFormatNames{},
		Periods: cldr.CalendarPeriodFormatNames{},
	},
}
