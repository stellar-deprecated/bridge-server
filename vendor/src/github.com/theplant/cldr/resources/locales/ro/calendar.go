package ro

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd.MM.y"},
		Time:     cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
		DateTime: cldr.CalendarDateFormat{Full: "{1}, {0}", Long: "{1}, {0}", Medium: "{1}, {0}", Short: "{1}, {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "ian.", Feb: "feb.", Mar: "mar.", Apr: "apr.", May: "mai", Jun: "iun.", Jul: "iul.", Aug: "aug.", Sep: "sept.", Oct: "oct.", Nov: "nov.", Dec: "dec."},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "I", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "I", Jul: "I", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Ianuarie", Feb: "Februarie", Mar: "Martie", Apr: "Aprilie", May: "Mai", Jun: "Iunie", Jul: "Iulie", Aug: "August", Sep: "Septembrie", Oct: "Octombrie", Nov: "Noiembrie", Dec: "Decembrie"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Dum", Mon: "Lun", Tue: "Mar", Wed: "Mie", Thu: "Joi", Fri: "Vin", Sat: "Sâm"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "D", Mon: "L", Tue: "M", Wed: "M", Thu: "J", Fri: "V", Sat: "S"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "Du", Mon: "Lu", Tue: "Ma", Wed: "Mi", Thu: "Jo", Fri: "Vi", Sat: "Sâ"},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "Duminică", Mon: "Luni", Tue: "Marți", Wed: "Miercuri", Thu: "Joi", Fri: "Vineri", Sat: "Sâmbătă"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "a.m.", PM: "p.m."},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "a.m.", PM: "p.m."},
		},
	},
}
