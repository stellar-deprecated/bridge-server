package dav

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"},
		Time:     cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Imb", Feb: "Kaw", Mar: "Kad", Apr: "Kan", May: "Kas", Jun: "Kar", Jul: "Mfu", Aug: "Wun", Sep: "Ike", Oct: "Iku", Nov: "Imw", Dec: "Iwi"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "I", Feb: "K", Mar: "K", Apr: "K", May: "K", Jun: "K", Jul: "M", Aug: "W", Sep: "I", Oct: "I", Nov: "I", Dec: "I"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Mori ghwa imbiri", Feb: "Mori ghwa kawi", Mar: "Mori ghwa kadadu", Apr: "Mori ghwa kana", May: "Mori ghwa kasanu", Jun: "Mori ghwa karandadu", Jul: "Mori ghwa mfungade", Aug: "Mori ghwa wunyanya", Sep: "Mori ghwa ikenda", Oct: "Mori ghwa ikumi", Nov: "Mori ghwa ikumi na imweri", Dec: "Mori ghwa ikumi na iwi"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Jum", Mon: "Jim", Tue: "Kaw", Wed: "Kad", Thu: "Kan", Fri: "Kas", Sat: "Ngu"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "J", Mon: "J", Tue: "K", Wed: "K", Thu: "K", Fri: "K", Sat: "N"},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "Ituku ja jumwa", Mon: "Kuramuka jimweri", Tue: "Kuramuka kawi", Wed: "Kuramuka kadadu", Thu: "Kuramuka kana", Fri: "Kuramuka kasanu", Sat: "Kifula nguwo"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "Luma lwa K", PM: "luma lwa p"},
		},
	},
}
