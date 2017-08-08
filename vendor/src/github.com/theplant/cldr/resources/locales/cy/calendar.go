package cy

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/yy"},
		Time:     cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
		DateTime: cldr.CalendarDateFormat{Full: "{1} 'am' {0}", Long: "{1} 'am' {0}", Medium: "{1} {0}", Short: "{1} {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Ion", Feb: "Chw", Mar: "Maw", Apr: "Ebr", May: "Mai", Jun: "Meh", Jul: "Gor", Aug: "Awst", Sep: "Medi", Oct: "Hyd", Nov: "Tach", Dec: "Rhag"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "I", Feb: "Ch", Mar: "M", Apr: "E", May: "M", Jun: "M", Jul: "G", Aug: "A", Sep: "M", Oct: "H", Nov: "T", Dec: "Rh"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Ionawr", Feb: "Chwefror", Mar: "Mawrth", Apr: "Ebrill", May: "Mai", Jun: "Mehefin", Jul: "Gorffennaf", Aug: "Awst", Sep: "Medi", Oct: "Hydref", Nov: "Tachwedd", Dec: "Rhagfyr"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Sul", Mon: "Llun", Tue: "Maw", Wed: "Mer", Thu: "Iau", Fri: "Gwe", Sat: "Sad"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "S", Mon: "Ll", Tue: "M", Wed: "M", Thu: "I", Fri: "G", Sat: "S"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "Su", Mon: "Ll", Tue: "Ma", Wed: "Me", Thu: "Ia", Fri: "Gw", Sat: "Sa"},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "Dydd Sul", Mon: "Dydd Llun", Tue: "Dydd Mawrth", Wed: "Dydd Mercher", Thu: "Dydd Iau", Fri: "Dydd Gwener", Sat: "Dydd Sadwrn"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"},
		},
	},
}
