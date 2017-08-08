package bez

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"},
		Time:     cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Hut", Feb: "Vil", Mar: "Dat", Apr: "Tai", May: "Han", Jun: "Sit", Jul: "Sab", Aug: "Nan", Sep: "Tis", Oct: "Kum", Nov: "Kmj", Dec: "Kmb"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "H", Feb: "V", Mar: "D", Apr: "T", May: "H", Jun: "S", Jul: "S", Aug: "N", Sep: "T", Oct: "K", Nov: "K", Dec: "K"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "pa mwedzi gwa hutala", Feb: "pa mwedzi gwa wuvili", Mar: "pa mwedzi gwa wudatu", Apr: "pa mwedzi gwa wutai", May: "pa mwedzi gwa wuhanu", Jun: "pa mwedzi gwa sita", Jul: "pa mwedzi gwa saba", Aug: "pa mwedzi gwa nane", Sep: "pa mwedzi gwa tisa", Oct: "pa mwedzi gwa kumi", Nov: "pa mwedzi gwa kumi na moja", Dec: "pa mwedzi gwa kumi na mbili"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Mul", Mon: "Vil", Tue: "Hiv", Wed: "Hid", Thu: "Hit", Fri: "Hih", Sat: "Lem"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "M", Mon: "J", Tue: "H", Wed: "H", Thu: "H", Fri: "W", Sat: "J"},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "pa mulungu", Mon: "pa shahuviluha", Tue: "pa hivili", Wed: "pa hidatu", Thu: "pa hitayi", Fri: "pa hihanu", Sat: "pa shahulembela"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "pamilau", PM: "pamunyi"},
		},
	},
}
