package mgh

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"},
		Time:     cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Kwa", Feb: "Una", Mar: "Rar", Apr: "Che", May: "Tha", Jun: "Moc", Jul: "Sab", Aug: "Nan", Sep: "Tis", Oct: "Kum", Nov: "Moj", Dec: "Yel"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "K", Feb: "U", Mar: "R", Apr: "C", May: "T", Jun: "M", Jul: "S", Aug: "N", Sep: "T", Oct: "K", Nov: "M", Dec: "Y"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Mweri wo kwanza", Feb: "Mweri wo unayeli", Mar: "Mweri wo uneraru", Apr: "Mweri wo unecheshe", May: "Mweri wo unethanu", Jun: "Mweri wo thanu na mocha", Jul: "Mweri wo saba", Aug: "Mweri wo nane", Sep: "Mweri wo tisa", Oct: "Mweri wo kumi", Nov: "Mweri wo kumi na moja", Dec: "Mweri wo kumi na yel’li"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Sab", Mon: "Jtt", Tue: "Jnn", Wed: "Jtn", Thu: "Ara", Fri: "Iju", Sat: "Jmo"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "S", Mon: "J", Tue: "J", Wed: "J", Thu: "A", Fri: "I", Sat: "J"},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "Sabato", Mon: "Jumatatu", Tue: "Jumanne", Wed: "Jumatano", Thu: "Arahamisi", Fri: "Ijumaa", Sat: "Jumamosi"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "wichishu", PM: "mchochil’l"},
		},
	},
}
