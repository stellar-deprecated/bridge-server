package ms

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "d/MM/yy"},
		Time:     cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
		DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Jan", Feb: "Feb", Mar: "Mac", Apr: "Apr", May: "Mei", Jun: "Jun", Jul: "Jul", Aug: "Ogo", Sep: "Sep", Oct: "Okt", Nov: "Nov", Dec: "Dis"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "O", Sep: "S", Oct: "O", Nov: "N", Dec: "D"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Januari", Feb: "Februari", Mar: "Mac", Apr: "April", May: "Mei", Jun: "Jun", Jul: "Julai", Aug: "Ogos", Sep: "September", Oct: "Oktober", Nov: "November", Dec: "Disember"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Ahd", Mon: "Isn", Tue: "Sel", Wed: "Rab", Thu: "Kha", Fri: "Jum", Sat: "Sab"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "A", Mon: "I", Tue: "S", Wed: "R", Thu: "K", Fri: "J", Sat: "S"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "Ah", Mon: "Is", Tue: "Se", Wed: "Ra", Thu: "Kh", Fri: "Ju", Sat: "Sa"},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "Ahad", Mon: "Isnin", Tue: "Selasa", Wed: "Rabu", Thu: "Khamis", Fri: "Jumaat", Sat: "Sabtu"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "pg", PM: "ptg"},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "a", PM: "p"},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "PG", PM: "PTG"},
		},
	},
}
