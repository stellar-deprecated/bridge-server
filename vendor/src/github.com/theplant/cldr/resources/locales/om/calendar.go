package om

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, MMMM d, y", Long: "dd MMMM y", Medium: "dd-MMM-y", Short: "dd/MM/yy"},
		Time:     cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Ama", Feb: "Gur", Mar: "Bit", Apr: "Elb", May: "Cam", Jun: "Wax", Jul: "Ado", Aug: "Hag", Sep: "Ful", Oct: "Onk", Nov: "Sad", Dec: "Mud"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Amajjii", Feb: "Guraandhala", Mar: "Bitooteessa", Apr: "Elba", May: "Caamsa", Jun: "Waxabajjii", Jul: "Adooleessa", Aug: "Hagayya", Sep: "Fuulbana", Oct: "Onkololeessa", Nov: "Sadaasa", Dec: "Muddee"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Dil", Mon: "Wix", Tue: "Qib", Wed: "Rob", Thu: "Kam", Fri: "Jim", Sat: "San"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "S", Mon: "M", Tue: "T", Wed: "W", Thu: "T", Fri: "F", Sat: "S"},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "Dilbata", Mon: "Wiixata", Tue: "Qibxata", Wed: "Roobii", Thu: "Kamiisa", Fri: "Jimaata", Sat: "Sanbata"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "WD", PM: "WB"},
		},
	},
}
