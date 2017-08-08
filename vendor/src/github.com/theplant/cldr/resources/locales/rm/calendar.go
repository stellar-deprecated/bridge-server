package rm

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, 'ils' d 'da' MMMM y", Long: "d 'da' MMMM y", Medium: "dd-MM-y", Short: "dd-MM-yy"},
		Time:     cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "schan.", Feb: "favr.", Mar: "mars", Apr: "avr.", May: "matg", Jun: "zercl.", Jul: "fan.", Aug: "avust", Sep: "sett.", Oct: "oct.", Nov: "nov.", Dec: "dec."},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "S", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "Z", Jul: "F", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "schaner", Feb: "favrer", Mar: "mars", Apr: "avrigl", May: "matg", Jun: "zercladur", Jul: "fanadur", Aug: "avust", Sep: "settember", Oct: "october", Nov: "november", Dec: "december"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "du", Mon: "gli", Tue: "ma", Wed: "me", Thu: "gie", Fri: "ve", Sat: "so"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "D", Mon: "G", Tue: "M", Wed: "M", Thu: "G", Fri: "V", Sat: "S"},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "dumengia", Mon: "glindesdi", Tue: "mardi", Wed: "mesemna", Thu: "gievgia", Fri: "venderdi", Sat: "sonda"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "am", PM: "sm"},
		},
	},
}
