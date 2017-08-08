package az_Cyrl

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, d, MMMM, y", Long: "d MMMM, y", Medium: "d MMM, y", Short: "dd.MM.yy"},
		Time:     cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{},
			Narrow:      cldr.CalendarMonthFormatNameValue{},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "јанвар", Feb: "феврал", Mar: "март", Apr: "апрел", May: "май", Jun: "ијун", Jul: "ијул", Aug: "август", Sep: "сентјабр", Oct: "октјабр", Nov: "нојабр", Dec: "декабр"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "7", Mon: "1", Tue: "2", Wed: "3", Thu: "4", Fri: "5", Sat: "6"},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "базар", Mon: "базар ертәси", Tue: "чәршәнбә ахшамы", Wed: "чәршәнбә", Thu: "ҹүмә ахшамы", Fri: "ҹүмә", Sat: "шәнбә"},
		},
		Periods: cldr.CalendarPeriodFormatNames{},
	},
}
