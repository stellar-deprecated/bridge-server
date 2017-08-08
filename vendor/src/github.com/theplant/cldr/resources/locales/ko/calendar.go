package ko

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "y년 M월 d일 EEEE", Long: "y년 M월 d일", Medium: "y. M. d.", Short: "yy. M. d."},
		Time:     cldr.CalendarDateFormat{Full: "a h시 m분 s초 zzzz", Long: "a h시 m분 s초 z", Medium: "a h:mm:ss", Short: "a h:mm"},
		DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "1월", Feb: "2월", Mar: "3월", Apr: "4월", May: "5월", Jun: "6월", Jul: "7월", Aug: "8월", Sep: "9월", Oct: "10월", Nov: "11월", Dec: "12월"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "1월", Feb: "2월", Mar: "3월", Apr: "4월", May: "5월", Jun: "6월", Jul: "7월", Aug: "8월", Sep: "9월", Oct: "10월", Nov: "11월", Dec: "12월"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "1월", Feb: "2월", Mar: "3월", Apr: "4월", May: "5월", Jun: "6월", Jul: "7월", Aug: "8월", Sep: "9월", Oct: "10월", Nov: "11월", Dec: "12월"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "일", Mon: "월", Tue: "화", Wed: "수", Thu: "목", Fri: "금", Sat: "토"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "일", Mon: "월", Tue: "화", Wed: "수", Thu: "목", Fri: "금", Sat: "토"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "일", Mon: "월", Tue: "화", Wed: "수", Thu: "목", Fri: "금", Sat: "토"},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "일요일", Mon: "월요일", Tue: "화요일", Wed: "수요일", Thu: "목요일", Fri: "금요일", Sat: "토요일"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "오전", PM: "오후"},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "오전", PM: "오후"},
		},
	},
}
