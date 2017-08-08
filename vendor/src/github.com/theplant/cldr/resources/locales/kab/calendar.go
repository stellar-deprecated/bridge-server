package kab

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM, y", Short: "d/M/y"},
		Time:     cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Yen", Feb: "Fur", Mar: "Meɣ", Apr: "Yeb", May: "May", Jun: "Yun", Jul: "Yul", Aug: "Ɣuc", Sep: "Cte", Oct: "Tub", Nov: "Nun", Dec: "Duǧ"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "Y", Feb: "F", Mar: "M", Apr: "Y", May: "M", Jun: "Y", Jul: "Y", Aug: "Ɣ", Sep: "C", Oct: "T", Nov: "N", Dec: "D"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Yennayer", Feb: "Fuṛar", Mar: "Meɣres", Apr: "Yebrir", May: "Mayyu", Jun: "Yunyu", Jul: "Yulyu", Aug: "Ɣuct", Sep: "Ctembeṛ", Oct: "Tubeṛ", Nov: "Nunembeṛ", Dec: "Duǧembeṛ"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Yan", Mon: "San", Tue: "Kraḍ", Wed: "Kuẓ", Thu: "Sam", Fri: "Sḍis", Sat: "Say"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "Y", Mon: "S", Tue: "K", Wed: "K", Thu: "S", Fri: "S", Sat: "S"},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "Yanass", Mon: "Sanass", Tue: "Kraḍass", Wed: "Kuẓass", Thu: "Samass", Fri: "Sḍisass", Sat: "Sayass"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "n tufat", PM: "n tmeddit"},
		},
	},
}
