package guz

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"},
		Time:     cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Can", Feb: "Feb", Mar: "Mac", Apr: "Apr", May: "Mei", Jun: "Jun", Jul: "Cul", Aug: "Agt", Sep: "Sep", Oct: "Okt", Nov: "Nob", Dec: "Dis"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "C", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "C", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Chanuari", Feb: "Feburari", Mar: "Machi", Apr: "Apiriri", May: "Mei", Jun: "Juni", Jul: "Chulai", Aug: "Agosti", Sep: "Septemba", Oct: "Okitoba", Nov: "Nobemba", Dec: "Disemba"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Cpr", Mon: "Ctt", Tue: "Cmn", Wed: "Cmt", Thu: "Ars", Fri: "Icm", Sat: "Est"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "C", Mon: "C", Tue: "C", Wed: "C", Thu: "A", Fri: "I", Sat: "E"},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "Chumapiri", Mon: "Chumatato", Tue: "Chumaine", Wed: "Chumatano", Thu: "Aramisi", Fri: "Ichuma", Sat: "Esabato"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "Ma/Mo", PM: "Mambia/Mog"},
		},
	},
}
