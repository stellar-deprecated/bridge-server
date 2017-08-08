package kde

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"},
		Time:     cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Jan", Feb: "Feb", Mar: "Mac", Apr: "Apr", May: "Mei", Jun: "Jun", Jul: "Jul", Aug: "Ago", Sep: "Sep", Oct: "Okt", Nov: "Nov", Dec: "Des"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Mwedi Ntandi", Feb: "Mwedi wa Pili", Mar: "Mwedi wa Tatu", Apr: "Mwedi wa Nchechi", May: "Mwedi wa Nnyano", Jun: "Mwedi wa Nnyano na Umo", Jul: "Mwedi wa Nnyano na Mivili", Aug: "Mwedi wa Nnyano na Mitatu", Sep: "Mwedi wa Nnyano na Nchechi", Oct: "Mwedi wa Nnyano na Nnyano", Nov: "Mwedi wa Nnyano na Nnyano na U", Dec: "Mwedi wa Nnyano na Nnyano na M"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Ll2", Mon: "Ll3", Tue: "Ll4", Wed: "Ll5", Thu: "Ll6", Fri: "Ll7", Sat: "Ll1"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "2", Mon: "3", Tue: "4", Wed: "5", Thu: "6", Fri: "7", Sat: "1"},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "Liduva lyapili", Mon: "Liduva lyatatu", Tue: "Liduva lyanchechi", Wed: "Liduva lyannyano", Thu: "Liduva lyannyano na linji", Fri: "Liduva lyannyano na mavili", Sat: "Liduva litandi"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "Muhi", PM: "Chilo"},
		},
	},
}
