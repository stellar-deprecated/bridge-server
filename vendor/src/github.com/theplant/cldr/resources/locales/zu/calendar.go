package zu

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, MMMM d, y", Long: "MMMM d, y", Medium: "MMM d, y", Short: "M/d/yy"},
		Time:     cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
		DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Jan", Feb: "Feb", Mar: "Mas", Apr: "Apr", May: "Mey", Jun: "Jun", Jul: "Jul", Aug: "Aga", Sep: "Sep", Oct: "Okt", Nov: "Nov", Dec: "Dis"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Januwari", Feb: "Februwari", Mar: "Mashi", Apr: "Apreli", May: "Meyi", Jun: "Juni", Jul: "Julayi", Aug: "Agasti", Sep: "Septhemba", Oct: "Okthoba", Nov: "Novemba", Dec: "Disemba"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Son", Mon: "Mso", Tue: "Bil", Wed: "Tha", Thu: "Sin", Fri: "Hla", Sat: "Mgq"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "S", Mon: "M", Tue: "B", Wed: "T", Thu: "S", Fri: "H", Sat: "M"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "Isonto", Mon: "Umsombuluko", Tue: "Ulwesibili", Wed: "Ulwesithathu", Thu: "Ulwesine", Fri: "Ulwesihlanu", Sat: "Umgqibelo"},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "Sonto", Mon: "Msombuluko", Tue: "Lwesibili", Wed: "Lwesithathu", Thu: "Lwesine", Fri: "Lwesihlanu", Sat: "Mgqibelo"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "a", PM: "p"},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "Ekuseni", PM: "Ntambama"},
		},
	},
}
