package nd

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"},
		Time:     cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Zib", Feb: "Nhlo", Mar: "Mbi", Apr: "Mab", May: "Nkw", Jun: "Nhla", Jul: "Ntu", Aug: "Ncw", Sep: "Mpan", Oct: "Mfu", Nov: "Lwe", Dec: "Mpal"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "Z", Feb: "N", Mar: "M", Apr: "M", May: "N", Jun: "N", Jul: "N", Aug: "N", Sep: "M", Oct: "M", Nov: "L", Dec: "M"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Zibandlela", Feb: "Nhlolanja", Mar: "Mbimbitho", Apr: "Mabasa", May: "Nkwenkwezi", Jun: "Nhlangula", Jul: "Ntulikazi", Aug: "Ncwabakazi", Sep: "Mpandula", Oct: "Mfumfu", Nov: "Lwezi", Dec: "Mpalakazi"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Son", Mon: "Mvu", Tue: "Sib", Wed: "Sit", Thu: "Sin", Fri: "Sih", Sat: "Mgq"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "S", Mon: "M", Tue: "S", Wed: "S", Thu: "S", Fri: "S", Sat: "M"},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "Sonto", Mon: "Mvulo", Tue: "Sibili", Wed: "Sithathu", Thu: "Sine", Fri: "Sihlanu", Sat: "Mgqibelo"},
		},
		Periods: cldr.CalendarPeriodFormatNames{},
	},
}
