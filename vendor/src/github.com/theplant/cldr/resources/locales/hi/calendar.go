package hi

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "dd/MM/y", Short: "d/M/yy"},
		Time:     cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
		DateTime: cldr.CalendarDateFormat{Full: "{1} को {0}", Long: "{1} को {0}", Medium: "{1}, {0}", Short: "{1}, {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "जन॰", Feb: "फ़र॰", Mar: "मार्च", Apr: "अप्रैल", May: "मई", Jun: "जून", Jul: "जुल॰", Aug: "अग॰", Sep: "सित॰", Oct: "अक्तू॰", Nov: "नव॰", Dec: "दिस॰"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "ज", Feb: "फ़", Mar: "मा", Apr: "अ", May: "म", Jun: "जू", Jul: "जु", Aug: "अ", Sep: "सि", Oct: "अ", Nov: "न", Dec: "दि"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "जनवरी", Feb: "फ़रवरी", Mar: "मार्च", Apr: "अप्रैल", May: "मई", Jun: "जून", Jul: "जुलाई", Aug: "अगस्त", Sep: "सितंबर", Oct: "अक्तूबर", Nov: "नवंबर", Dec: "दिसंबर"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "रवि", Mon: "सोम", Tue: "मंगल", Wed: "बुध", Thu: "गुरु", Fri: "शुक्र", Sat: "शनि"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "र", Mon: "सो", Tue: "मं", Wed: "बु", Thu: "गु", Fri: "शु", Sat: "श"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "र", Mon: "सो", Tue: "मं", Wed: "बु", Thu: "गु", Fri: "शु", Sat: "श"},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "रविवार", Mon: "सोमवार", Tue: "मंगलवार", Wed: "बुधवार", Thu: "गुरुवार", Fri: "शुक्रवार", Sat: "शनिवार"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "पूर्व", PM: "अपर"},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "पू", PM: "अ"},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "पूर्वाह्न", PM: "अपराह्न"},
		},
	},
}
