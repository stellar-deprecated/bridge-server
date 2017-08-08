package ja

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "y年M月d日EEEE", Long: "y年M月d日", Medium: "y/MM/dd", Short: "y/MM/dd"},
		Time:     cldr.CalendarDateFormat{Full: "H時mm分ss秒 zzzz", Long: "H:mm:ss z", Medium: "H:mm:ss", Short: "H:mm"},
		DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "1月", Feb: "2月", Mar: "3月", Apr: "4月", May: "5月", Jun: "6月", Jul: "7月", Aug: "8月", Sep: "9月", Oct: "10月", Nov: "11月", Dec: "12月"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "1", Feb: "2", Mar: "3", Apr: "4", May: "5", Jun: "6", Jul: "7", Aug: "8", Sep: "9", Oct: "10", Nov: "11", Dec: "12"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "1月", Feb: "2月", Mar: "3月", Apr: "4月", May: "5月", Jun: "6月", Jul: "7月", Aug: "8月", Sep: "9月", Oct: "10月", Nov: "11月", Dec: "12月"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "日", Mon: "月", Tue: "火", Wed: "水", Thu: "木", Fri: "金", Sat: "土"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "日", Mon: "月", Tue: "火", Wed: "水", Thu: "木", Fri: "金", Sat: "土"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "日", Mon: "月", Tue: "火", Wed: "水", Thu: "木", Fri: "金", Sat: "土"},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "日曜日", Mon: "月曜日", Tue: "火曜日", Wed: "水曜日", Thu: "木曜日", Fri: "金曜日", Sat: "土曜日"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "午前", PM: "午後"},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "午前", PM: "午後"},
		},
	},
}
