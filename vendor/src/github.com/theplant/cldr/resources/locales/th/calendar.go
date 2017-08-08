package th

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEEที่ d MMMM G y", Long: "d MMMM G y", Medium: "d MMM y", Short: "d/M/yy"},
		Time:     cldr.CalendarDateFormat{Full: "H นาฬิกา mm นาที ss วินาที zzzz", Long: "H นาฬิกา mm นาที ss วินาที z", Medium: "HH:mm:ss", Short: "HH:mm"},
		DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "ม.ค.", Feb: "ก.พ.", Mar: "มี.ค.", Apr: "เม.ย.", May: "พ.ค.", Jun: "มิ.ย.", Jul: "ก.ค.", Aug: "ส.ค.", Sep: "ก.ย.", Oct: "ต.ค.", Nov: "พ.ย.", Dec: "ธ.ค."},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "ม.ค.", Feb: "ก.พ.", Mar: "มี.ค.", Apr: "เม.ย.", May: "พ.ค.", Jun: "มิ.ย.", Jul: "ก.ค.", Aug: "ส.ค.", Sep: "ก.ย.", Oct: "ต.ค.", Nov: "พ.ย.", Dec: "ธ.ค."},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "มกราคม", Feb: "กุมภาพันธ์", Mar: "มีนาคม", Apr: "เมษายน", May: "พฤษภาคม", Jun: "มิถุนายน", Jul: "กรกฎาคม", Aug: "สิงหาคม", Sep: "กันยายน", Oct: "ตุลาคม", Nov: "พฤศจิกายน", Dec: "ธันวาคม"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "อา.", Mon: "จ.", Tue: "อ.", Wed: "พ.", Thu: "พฤ.", Fri: "ศ.", Sat: "ส."},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "อา", Mon: "จ", Tue: "อ", Wed: "พ", Thu: "พฤ", Fri: "ศ", Sat: "ส"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "อา.", Mon: "จ.", Tue: "อ.", Wed: "พ.", Thu: "พฤ.", Fri: "ศ.", Sat: "ส."},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "วันอาทิตย์", Mon: "วันจันทร์", Tue: "วันอังคาร", Wed: "วันพุธ", Thu: "วันพฤหัสบดี", Fri: "วันศุกร์", Sat: "วันเสาร์"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "a", PM: "p"},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "ก่อนเที่ยง", PM: "หลังเที่ยง"},
		},
	},
}
