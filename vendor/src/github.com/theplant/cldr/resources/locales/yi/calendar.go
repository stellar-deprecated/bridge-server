package yi

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, dטן MMMM y", Long: "dטן MMMM y", Medium: "dטן MMM y", Short: "dd/MM/yy"},
		Time:     cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
		DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1}, {0}", Short: "{1}, {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "יאַנ", Feb: "פֿעב", Mar: "מערץ", Apr: "אַפּר", May: "מיי", Jun: "יוני", Jul: "יולי", Aug: "אויגוסט", Sep: "סעפּ", Oct: "אקט", Nov: "נאוו", Dec: "דעצ"},
			Narrow:      cldr.CalendarMonthFormatNameValue{},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "יאַנואַר", Feb: "פֿעברואַר", Mar: "מערץ", Apr: "אַפּריל", May: "מיי", Jun: "יוני", Jul: "יולי", Aug: "אויגוסט", Sep: "סעפּטעמבער", Oct: "אקטאבער", Nov: "נאוועמבער", Dec: "דעצעמבער"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "זונטיק", Mon: "מאָנטיק", Tue: "דינסטיק", Wed: "מיטוואך", Thu: "דאנערשטיק", Fri: "פֿרײַטיק", Sat: "שבת"},
			Narrow:      cldr.CalendarDayFormatNameValue{},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "זונטיק", Mon: "מאָנטיק", Tue: "דינסטיק", Wed: "מיטוואך", Thu: "דאנערשטיק", Fri: "פֿרײַטיק", Sat: "שבת"},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "זונטיק", Mon: "מאָנטיק", Tue: "דינסטיק", Wed: "מיטוואך", Thu: "דאנערשטיק", Fri: "פֿרײַטיק", Sat: "שבת"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "פארמיטאג", PM: "נאכמיטאג"},
		},
	},
}
