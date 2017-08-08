package vi

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, 'ngày' dd MMMM 'năm' y", Long: "'Ngày' dd 'tháng' MM 'năm' y", Medium: "dd-MM-y", Short: "dd/MM/y"},
		Time:     cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
		DateTime: cldr.CalendarDateFormat{Full: "{0} {1}", Long: "{0} {1}", Medium: "{0} {1}", Short: "{0} {1}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Thg 1", Feb: "Thg 2", Mar: "Thg 3", Apr: "Thg 4", May: "Thg 5", Jun: "Thg 6", Jul: "Thg 7", Aug: "Thg 8", Sep: "Thg 9", Oct: "Thg 10", Nov: "Thg 11", Dec: "Thg 12"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "1", Feb: "2", Mar: "3", Apr: "4", May: "5", Jun: "6", Jul: "7", Aug: "8", Sep: "9", Oct: "10", Nov: "11", Dec: "12"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Tháng 1", Feb: "Tháng 2", Mar: "Tháng 3", Apr: "Tháng 4", May: "Tháng 5", Jun: "Tháng 6", Jul: "Tháng 7", Aug: "Tháng 8", Sep: "Tháng 9", Oct: "Tháng 10", Nov: "Tháng 11", Dec: "Tháng 12"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "CN", Mon: "Th 2", Tue: "Th 3", Wed: "Th 4", Thu: "Th 5", Fri: "Th 6", Sat: "Th 7"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "CN", Mon: "T2", Tue: "T3", Wed: "T4", Thu: "T5", Fri: "T6", Sat: "T7"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "CN", Mon: "T2", Tue: "T3", Wed: "T4", Thu: "T5", Fri: "T6", Sat: "T7"},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "Chủ Nhật", Mon: "Thứ Hai", Tue: "Thứ Ba", Wed: "Thứ Tư", Thu: "Thứ Năm", Fri: "Thứ Sáu", Sat: "Thứ Bảy"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "s", PM: "c"},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "SA", PM: "CH"},
		},
	},
}
