package pa

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "d/M/yy"},
		Time:     cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
		DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1}, {0}", Short: "{1}, {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "ਜਨ", Feb: "ਫ਼ਰ", Mar: "ਮਾਰਚ", Apr: "ਅਪ੍ਰੈ", May: "ਮਈ", Jun: "ਜੂਨ", Jul: "ਜੁਲਾ", Aug: "ਅਗ", Sep: "ਸਤੰ", Oct: "ਅਕਤੂ", Nov: "ਨਵੰ", Dec: "ਦਸੰ"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "ਜ", Feb: "ਫ਼", Mar: "ਮਾ", Apr: "ਅ", May: "ਮ", Jun: "ਜੂ", Jul: "ਜੁ", Aug: "ਅ", Sep: "ਸ", Oct: "ਅ", Nov: "ਨ", Dec: "ਦ"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "ਜਨਵਰੀ", Feb: "ਫ਼ਰਵਰੀ", Mar: "ਮਾਰਚ", Apr: "ਅਪ੍ਰੈਲ", May: "ਮਈ", Jun: "ਜੂਨ", Jul: "ਜੁਲਾਈ", Aug: "ਅਗਸਤ", Sep: "ਸਤੰਬਰ", Oct: "ਅਕਤੂਬਰ", Nov: "ਨਵੰਬਰ", Dec: "ਦਸੰਬਰ"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "ਐਤ", Mon: "ਸੋਮ", Tue: "ਮੰਗਲ", Wed: "ਬੁੱਧ", Thu: "ਵੀਰ", Fri: "ਸ਼ੁੱਕਰ", Sat: "ਸ਼ਨਿੱਚਰ"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "ਐ", Mon: "ਸੋ", Tue: "ਮੰ", Wed: "ਬੁੱ", Thu: "ਵੀ", Fri: "ਸ਼ੁੱ", Sat: "ਸ਼"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "ਐਤ", Mon: "ਸੋਮ", Tue: "ਮੰਗ", Wed: "ਬੁੱਧ", Thu: "ਵੀਰ", Fri: "ਸ਼ੁੱਕ", Sat: "ਸ਼ਨਿੱ"},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "ਐਤਵਾਰ", Mon: "ਸੋਮਵਾਰ", Tue: "ਮੰਗਲਵਾਰ", Wed: "ਬੁੱਧਵਾਰ", Thu: "ਵੀਰਵਾਰ", Fri: "ਸ਼ੁੱਕਰਵਾਰ", Sat: "ਸ਼ਨਿੱਚਰਵਾਰ"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "ਸ.", PM: "ਸ਼."},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "ਪੂ.ਦੁ.", PM: "ਬਾ.ਦੁ."},
		},
	},
}
