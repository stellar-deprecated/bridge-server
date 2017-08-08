package brx

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, MMMM d, y", Long: "MMMM d, y", Medium: "MMM d, y", Short: "M/d/yy"},
		Time:     cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "ज", Feb: "फे", Mar: "मा", Apr: "ए", May: "मे", Jun: "जु", Jul: "जु", Aug: "आ", Sep: "से", Oct: "अ", Nov: "न", Dec: "दि"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "जानुवारी", Feb: "फेब्रुवारी", Mar: "मार्स", Apr: "एफ्रिल", May: "मे", Jun: "जुन", Jul: "जुलाइ", Aug: "आगस्थ", Sep: "सेबथेज्ब़र", Oct: "अखथबर", Nov: "नबेज्ब़र", Dec: "दिसेज्ब़र"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "रबि", Mon: "सम", Tue: "मंगल", Wed: "बुद", Thu: "बिसथि", Fri: "सुखुर", Sat: "सुनि"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "र", Mon: "स", Tue: "मं", Wed: "बु", Thu: "बि", Fri: "सु", Sat: "सु"},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "रबिबार", Mon: "समबार", Tue: "मंगलबार", Wed: "बुदबार", Thu: "बिसथिबार", Fri: "सुखुरबार", Sat: "सुनिबार"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "फुं", PM: "बेलासे"},
		},
	},
}
