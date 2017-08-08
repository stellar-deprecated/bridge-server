package es_UY

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Ene.", Feb: "Feb.", Mar: "Mar.", Apr: "Abr.", May: "May.", Jun: "Jun.", Jul: "Jul.", Aug: "Ago.", Sep: "Set.", Oct: "Oct.", Nov: "Nov.", Dec: "Dic."},
			Narrow:      cldr.CalendarMonthFormatNameValue{},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Enero", Feb: "Febrero", Mar: "Marzo", Apr: "Abril", May: "Mayo", Jun: "Junio", Jul: "Julio", Aug: "Agosto", Sep: "Setiembre", Oct: "Octubre", Nov: "Noviembre", Dec: "Diciembre"},
		},
		Days:    cldr.CalendarDayFormatNames{},
		Periods: cldr.CalendarPeriodFormatNames{},
	},
}
