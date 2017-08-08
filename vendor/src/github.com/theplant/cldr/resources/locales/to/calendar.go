package to

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "d/M/yy"},
		Time:     cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
		DateTime: cldr.CalendarDateFormat{Full: "{1}, {0}", Long: "{1}, {0}", Medium: "{1}, {0}", Short: "{1} {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Sān", Feb: "Fēp", Mar: "Maʻa", Apr: "ʻEpe", May: "Mē", Jun: "Sun", Jul: "Siu", Aug: "ʻAok", Sep: "Sep", Oct: "ʻOka", Nov: "Nōv", Dec: "Tīs"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "S", Feb: "F", Mar: "M", Apr: "E", May: "M", Jun: "S", Jul: "S", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "T"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Sānuali", Feb: "Fēpueli", Mar: "Maʻasi", Apr: "ʻEpeleli", May: "Mē", Jun: "Sune", Jul: "Siulai", Aug: "ʻAokosi", Sep: "Sepitema", Oct: "ʻOkatopa", Nov: "Nōvema", Dec: "Tīsema"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Sāp", Mon: "Mōn", Tue: "Tūs", Wed: "Pul", Thu: "Tuʻa", Fri: "Fal", Sat: "Tok"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "S", Mon: "M", Tue: "T", Wed: "P", Thu: "T", Fri: "F", Sat: "T"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "Sāp", Mon: "Mōn", Tue: "Tūs", Wed: "Pul", Thu: "Tuʻa", Fri: "Fal", Sat: "Tok"},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "Sāpate", Mon: "Mōnite", Tue: "Tūsite", Wed: "Pulelulu", Thu: "Tuʻapulelulu", Fri: "Falaite", Sat: "Tokonaki"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"},
		},
	},
}
