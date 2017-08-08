package mr

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, d MMMM, y", Long: "d MMMM, y", Medium: "d MMM, y", Short: "d/M/yy"},
		Time:     cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
		DateTime: cldr.CalendarDateFormat{Full: "{1} 'रोजी' {0}", Long: "{1} 'रोजी' {0}", Medium: "{1}, {0}", Short: "{1}, {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "जाने", Feb: "फेब्रु", Mar: "मार्च", Apr: "एप्रि", May: "मे", Jun: "जून", Jul: "जुलै", Aug: "ऑग", Sep: "सप्टें", Oct: "ऑक्टो", Nov: "नोव्हें", Dec: "डिसें"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "जा", Feb: "फे", Mar: "मा", Apr: "ए", May: "मे", Jun: "जू", Jul: "जु", Aug: "ऑ", Sep: "स", Oct: "ऑ", Nov: "नो", Dec: "डि"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "जानेवारी", Feb: "फेब्रुवारी", Mar: "मार्च", Apr: "एप्रिल", May: "मे", Jun: "जून", Jul: "जुलै", Aug: "ऑगस्ट", Sep: "सप्टेंबर", Oct: "ऑक्टोबर", Nov: "नोव्हेंबर", Dec: "डिसेंबर"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "रवि", Mon: "सोम", Tue: "मंगळ", Wed: "बुध", Thu: "गुरु", Fri: "शुक्र", Sat: "शनि"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "र", Mon: "सो", Tue: "मं", Wed: "बु", Thu: "गु", Fri: "शु", Sat: "श"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "र", Mon: "सो", Tue: "मं", Wed: "बु", Thu: "गु", Fri: "शु", Sat: "श"},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "रविवार", Mon: "सोमवार", Tue: "मंगळवार", Wed: "बुधवार", Thu: "गुरुवार", Fri: "शुक्रवार", Sat: "शनिवार"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "स", PM: "सं"},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "म.पू.", PM: "म.उ."},
		},
	},
}
