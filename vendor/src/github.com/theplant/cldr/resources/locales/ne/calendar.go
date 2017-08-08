package ne

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "y MMMM d, EEEE", Long: "y MMMM d", Medium: "y MMM d", Short: "y-MM-dd"},
		Time:     cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
		DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1}, {0}", Short: "{1}, {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "जनवरी", Feb: "फेब्रुअरी", Mar: "मार्च", Apr: "अप्रिल", May: "मे", Jun: "जुन", Jul: "जुलाई", Aug: "अगस्ट", Sep: "सेप्टेम्बर", Oct: "अक्टोबर", Nov: "नोभेम्बर", Dec: "डिसेम्बर"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "१", Feb: "२", Mar: "३", Apr: "४", May: "५", Jun: "६", Jul: "७", Aug: "८", Sep: "९", Oct: "१०", Nov: "११", Dec: "१२"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "जनवरी", Feb: "फेब्रुअरी", Mar: "मार्च", Apr: "अप्रिल", May: "मे", Jun: "जुन", Jul: "जुलाई", Aug: "अगस्ट", Sep: "सेप्टेम्बर", Oct: "अक्टोबर", Nov: "नोभेम्बर", Dec: "डिसेम्बर"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "आइत", Mon: "सोम", Tue: "मङ्गल", Wed: "बुध", Thu: "बिही", Fri: "शुक्र", Sat: "शनि"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "आ", Mon: "सो", Tue: "म", Wed: "बु", Thu: "बि", Fri: "शु", Sat: "श"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "आइत", Mon: "सोम", Tue: "मङ्गल", Wed: "बुध", Thu: "बिही", Fri: "शुक्र", Sat: "शनि"},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "आइतबार", Mon: "सोमबार", Tue: "मङ्गलबार", Wed: "बुधबार", Thu: "बिहिबार", Fri: "शुक्रबार", Sat: "शनिबार"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "a", PM: "p"},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "पूर्वाह्न", PM: "अपराह्न"},
		},
	},
}
