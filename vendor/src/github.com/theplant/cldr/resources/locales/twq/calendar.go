package twq

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "d/M/y"},
		Time:     cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Žan", Feb: "Fee", Mar: "Mar", Apr: "Awi", May: "Me", Jun: "Žuw", Jul: "Žuy", Aug: "Ut", Sep: "Sek", Oct: "Okt", Nov: "Noo", Dec: "Dee"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "Ž", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "Ž", Jul: "Ž", Aug: "U", Sep: "S", Oct: "O", Nov: "N", Dec: "D"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Žanwiye", Feb: "Feewiriye", Mar: "Marsi", Apr: "Awiril", May: "Me", Jun: "Žuweŋ", Jul: "Žuyye", Aug: "Ut", Sep: "Sektanbur", Oct: "Oktoobur", Nov: "Noowanbur", Dec: "Deesanbur"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Alh", Mon: "Ati", Tue: "Ata", Wed: "Ala", Thu: "Alm", Fri: "Alz", Sat: "Asi"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "H", Mon: "T", Tue: "T", Wed: "L", Thu: "L", Fri: "L", Sat: "S"},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "Alhadi", Mon: "Atinni", Tue: "Atalaata", Wed: "Alarba", Thu: "Alhamiisa", Fri: "Alzuma", Sat: "Asibti"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "Subbaahi", PM: "Zaarikay b"},
		},
	},
}
