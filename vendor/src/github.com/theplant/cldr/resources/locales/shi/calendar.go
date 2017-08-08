package shi

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM, y", Short: "d/M/y"},
		Time:     cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "ⵉⵏⵏ", Feb: "ⴱⵕⴰ", Mar: "ⵎⴰⵕ", Apr: "ⵉⴱⵔ", May: "ⵎⴰⵢ", Jun: "ⵢⵓⵏ", Jul: "ⵢⵓⵍ", Aug: "ⵖⵓⵛ", Sep: "ⵛⵓⵜ", Oct: "ⴽⵜⵓ", Nov: "ⵏⵓⵡ", Dec: "ⴷⵓⵊ"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "ⵉ", Feb: "ⴱ", Mar: "ⵎ", Apr: "ⵉ", May: "ⵎ", Jun: "ⵢ", Jul: "ⵢ", Aug: "ⵖ", Sep: "ⵛ", Oct: "ⴽ", Nov: "ⵏ", Dec: "ⴷ"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "ⵉⵏⵏⴰⵢⵔ", Feb: "ⴱⵕⴰⵢⵕ", Mar: "ⵎⴰⵕⵚ", Apr: "ⵉⴱⵔⵉⵔ", May: "ⵎⴰⵢⵢⵓ", Jun: "ⵢⵓⵏⵢⵓ", Jul: "ⵢⵓⵍⵢⵓⵣ", Aug: "ⵖⵓⵛⵜ", Sep: "ⵛⵓⵜⴰⵏⴱⵉⵔ", Oct: "ⴽⵜⵓⴱⵔ", Nov: "ⵏⵓⵡⴰⵏⴱⵉⵔ", Dec: "ⴷⵓⵊⴰⵏⴱⵉⵔ"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "ⴰⵙⴰ", Mon: "ⴰⵢⵏ", Tue: "ⴰⵙⵉ", Wed: "ⴰⴽⵕ", Thu: "ⴰⴽⵡ", Fri: "ⴰⵙⵉⵎ", Sat: "ⴰⵙⵉⴹ"},
			Narrow:      cldr.CalendarDayFormatNameValue{},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "ⴰⵙⴰⵎⴰⵙ", Mon: "ⴰⵢⵏⴰⵙ", Tue: "ⴰⵙⵉⵏⴰⵙ", Wed: "ⴰⴽⵕⴰⵙ", Thu: "ⴰⴽⵡⴰⵙ", Fri: "ⵙⵉⵎⵡⴰⵙ", Sat: "ⴰⵙⵉⴹⵢⴰⵙ"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "ⵜⵉⴼⴰⵡⵜ", PM: "ⵜⴰⴷⴳⴳⵯⴰⵜ"},
		},
	},
}
