package so

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, MMMM dd, y", Long: "dd MMMM y", Medium: "dd-MMM-y", Short: "dd/MM/yy"},
		Time:     cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Kob", Feb: "Lab", Mar: "Sad", Apr: "Afr", May: "Sha", Jun: "Lix", Jul: "Tod", Aug: "Sid", Sep: "Sag", Oct: "Tob", Nov: "KIT", Dec: "LIT"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "K", Feb: "L", Mar: "S", Apr: "A", May: "S", Jun: "L", Jul: "T", Aug: "S", Sep: "S", Oct: "T", Nov: "K", Dec: "L"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Bisha Koobaad", Feb: "Bisha Labaad", Mar: "Bisha Saddexaad", Apr: "Bisha Afraad", May: "Bisha Shanaad", Jun: "Bisha Lixaad", Jul: "Bisha Todobaad", Aug: "Bisha Sideedaad", Sep: "Bisha Sagaalaad", Oct: "Bisha Tobnaad", Nov: "Bisha Kow iyo Tobnaad", Dec: "Bisha Laba iyo Tobnaad"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Axd", Mon: "Isn", Tue: "Tal", Wed: "Arb", Thu: "Kha", Fri: "Jim", Sat: "Sab"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "A", Mon: "I", Tue: "T", Wed: "A", Thu: "K", Fri: "J", Sat: "S"},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "Axad", Mon: "Isniin", Tue: "Talaado", Wed: "Arbaco", Thu: "Khamiis", Fri: "Jimco", Sat: "Sabti"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "sn.", PM: "gn."},
		},
	},
}
