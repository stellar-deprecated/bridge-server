package ee

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, MMMM d 'lia' y", Long: "MMMM d 'lia' y", Medium: "MMM d 'lia', y", Short: "M/d/yy"},
		Time:     cldr.CalendarDateFormat{Full: "a h:mm:ss zzzz", Long: "a 'ga' h:mm:ss z", Medium: "a 'ga' h:mm:ss", Short: "a 'ga' h:mm"},
		DateTime: cldr.CalendarDateFormat{Full: "{0} {1}", Long: "{0} {1}", Medium: "{0} {1}", Short: "{0} {1}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "dzv", Feb: "dzd", Mar: "ted", Apr: "afɔ", May: "dam", Jun: "mas", Jul: "sia", Aug: "dea", Sep: "any", Oct: "kel", Nov: "ade", Dec: "dzm"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "d", Feb: "d", Mar: "t", Apr: "a", May: "d", Jun: "m", Jul: "s", Aug: "d", Sep: "a", Oct: "k", Nov: "a", Dec: "d"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "dzove", Feb: "dzodze", Mar: "tedoxe", Apr: "afɔfĩe", May: "dama", Jun: "masa", Jul: "siamlɔm", Aug: "deasiamime", Sep: "anyɔnyɔ", Oct: "kele", Nov: "adeɛmekpɔxe", Dec: "dzome"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "kɔs", Mon: "dzo", Tue: "bla", Wed: "kuɖ", Thu: "yaw", Fri: "fiɖ", Sat: "mem"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "k", Mon: "d", Tue: "b", Wed: "k", Thu: "y", Fri: "f", Sat: "m"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "kɔs", Mon: "dzo", Tue: "bla", Wed: "kuɖ", Thu: "yaw", Fri: "fiɖ", Sat: "mem"},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "kɔsiɖa", Mon: "dzoɖa", Tue: "blaɖa", Wed: "kuɖa", Thu: "yawoɖa", Fri: "fiɖa", Sat: "memleɖa"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "ŋ", PM: "ɣ"},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "ŋdi", PM: "ɣetrɔ"},
		},
	},
}
