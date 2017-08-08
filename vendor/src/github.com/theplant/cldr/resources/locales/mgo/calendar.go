package mgo

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, y MMMM dd", Long: "y MMMM d", Medium: "y MMM d", Short: "y-MM-dd"},
		Time:     cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
		DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "mbegtug", Feb: "imeg àbùbì", Mar: "imeg mbəŋchubi", Apr: "iməg ngwə̀t", May: "iməg fog", Jun: "iməg ichiibɔd", Jul: "iməg àdùmbə̀ŋ", Aug: "iməg ichika", Sep: "iməg kud", Oct: "iməg tèsiʼe", Nov: "iməg zò", Dec: "iməg krizmed"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "M1", Feb: "A2", Mar: "M3", Apr: "N4", May: "F5", Jun: "I6", Jul: "A7", Aug: "I8", Sep: "K9", Oct: "10", Nov: "11", Dec: "12"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "iməg mbegtug", Feb: "imeg àbùbì", Mar: "imeg mbəŋchubi", Apr: "iməg ngwə̀t", May: "iməg fog", Jun: "iməg ichiibɔd", Jul: "iməg àdùmbə̀ŋ", Aug: "iməg ichika", Sep: "iməg kud", Oct: "iməg tèsiʼe", Nov: "iməg zò", Dec: "iməg krizmed"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Aneg 1", Mon: "Aneg 2", Tue: "Aneg 3", Wed: "Aneg 4", Thu: "Aneg 5", Fri: "Aneg 6", Sat: "Aneg 7"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "A1", Mon: "A2", Tue: "A3", Wed: "A4", Thu: "A5", Fri: "A6", Sat: "A7"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "1", Mon: "2", Tue: "3", Wed: "4", Thu: "5", Fri: "6", Sat: "7"},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "Aneg 1", Mon: "Aneg 2", Tue: "Aneg 3", Wed: "Aneg 4", Thu: "Aneg 5", Fri: "Aneg 6", Sat: "Aneg 7"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"},
		},
	},
}
