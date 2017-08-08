package kok

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "dd-MM-y", Short: "d-M-yy"},
		Time:     cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{},
			Narrow:      cldr.CalendarMonthFormatNameValue{},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "जानेवारी", Feb: "फेब्रुवारी", Mar: "मार्च", Apr: "एप्रिल", May: "मे", Jun: "जून", Jul: "जुलै", Aug: "ओगस्ट", Sep: "सेप्टेंबर", Oct: "ओक्टोबर", Nov: "नोव्हेंबर", Dec: "डिसेंबर"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "रवि", Mon: "सोम", Tue: "मंगळ", Wed: "बुध", Thu: "गुरु", Fri: "शुक्र", Sat: "शनि"},
			Narrow:      cldr.CalendarDayFormatNameValue{},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "आदित्यवार", Mon: "सोमवार", Tue: "मंगळार", Wed: "बुधवार", Thu: "गुरुवार", Fri: "शुक्रवार", Sat: "शनिवार"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "म.पू.", PM: "म.नं."},
		},
	},
}
