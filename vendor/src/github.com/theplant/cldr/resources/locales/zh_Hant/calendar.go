package zh_Hant

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "y年M月d日 EEEE", Long: "y年M月d日", Medium: "y年M月d日", Short: "y/M/d"},
		Time:     cldr.CalendarDateFormat{Full: "ah:mm:ss [zzzz]", Long: "ah:mm:ss [z]", Medium: "ah:mm:ss", Short: "ah:mm"},
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
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "週日", Mon: "週一", Tue: "週二", Wed: "週三", Thu: "週四", Fri: "週五", Sat: "週六"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "日", Mon: "一", Tue: "二", Wed: "三", Thu: "四", Fri: "五", Sat: "六"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "日", Mon: "一", Tue: "二", Wed: "三", Thu: "四", Fri: "五", Sat: "六"},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "星期日", Mon: "星期一", Tue: "星期二", Wed: "星期三", Thu: "星期四", Fri: "星期五", Sat: "星期六"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "上午", PM: "下午"},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "上午", PM: "下午"},
		},
	},
}
