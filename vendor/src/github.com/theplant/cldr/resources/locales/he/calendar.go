package he

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, d בMMMM y", Long: "d בMMMM y", Medium: "d בMMM y", Short: "d.M.y"},
		Time:     cldr.CalendarDateFormat{Full: "H:mm:ss zzzz", Long: "H:mm:ss z", Medium: "H:mm:ss", Short: "H:mm"},
		DateTime: cldr.CalendarDateFormat{Full: "{1} בשעה {0}", Long: "{1} בשעה {0}", Medium: "{1}, {0}", Short: "{1}, {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "ינו׳", Feb: "פבר׳", Mar: "מרץ", Apr: "אפר׳", May: "מאי", Jun: "יוני", Jul: "יולי", Aug: "אוג׳", Sep: "ספט׳", Oct: "אוק׳", Nov: "נוב׳", Dec: "דצמ׳"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "1", Feb: "2", Mar: "3", Apr: "4", May: "5", Jun: "6", Jul: "7", Aug: "8", Sep: "9", Oct: "10", Nov: "11", Dec: "12"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "ינואר", Feb: "פברואר", Mar: "מרץ", Apr: "אפריל", May: "מאי", Jun: "יוני", Jul: "יולי", Aug: "אוגוסט", Sep: "ספטמבר", Oct: "אוקטובר", Nov: "נובמבר", Dec: "דצמבר"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "יום א׳", Mon: "יום ב׳", Tue: "יום ג׳", Wed: "יום ד׳", Thu: "יום ה׳", Fri: "יום ו׳", Sat: "שבת"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "א׳", Mon: "ב׳", Tue: "ג׳", Wed: "ד׳", Thu: "ה׳", Fri: "ו׳", Sat: "ש׳"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "א׳", Mon: "ב׳", Tue: "ג׳", Wed: "ד׳", Thu: "ה׳", Fri: "ו׳", Sat: "ש׳"},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "יום ראשון", Mon: "יום שני", Tue: "יום שלישי", Wed: "יום רביעי", Thu: "יום חמישי", Fri: "יום שישי", Sat: "יום שבת"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "לפנה״צ", PM: "אחה״צ"},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "לפנה״צ", PM: "אחה״צ"},
		},
	},
}
