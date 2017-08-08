package fr_CA

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "", Long: "", Medium: "", Short: "yy-MM-dd"},
		Time:     cldr.CalendarDateFormat{},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "janv.", Feb: "févr.", Mar: "mars", Apr: "avr.", May: "mai", Jun: "juin", Jul: "juil.", Aug: "août", Sep: "sept.", Oct: "oct.", Nov: "nov.", Dec: "déc."},
			Narrow:      cldr.CalendarMonthFormatNameValue{},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "janvier", Feb: "février", Mar: "mars", Apr: "avril", May: "mai", Jun: "juin", Jul: "juillet", Aug: "août", Sep: "septembre", Oct: "octobre", Nov: "novembre", Dec: "décembre"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "dim.", Mon: "lun.", Tue: "mar.", Wed: "mer.", Thu: "jeu.", Fri: "ven.", Sat: "sam."},
			Narrow:      cldr.CalendarDayFormatNameValue{},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "dimanche", Mon: "lundi", Tue: "mardi", Wed: "mercredi", Thu: "jeudi", Fri: "vendredi", Sat: "samedi"},
		},
		Periods: cldr.CalendarPeriodFormatNames{},
	},
}
