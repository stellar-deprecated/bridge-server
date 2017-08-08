package saq

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"},
		Time:     cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Obo", Feb: "Waa", Mar: "Oku", Apr: "Ong", May: "Ime", Jun: "Ile", Jul: "Sap", Aug: "Isi", Sep: "Saa", Oct: "Tom", Nov: "Tob", Dec: "Tow"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "O", Feb: "W", Mar: "O", Apr: "O", May: "I", Jun: "I", Jul: "S", Aug: "I", Sep: "S", Oct: "T", Nov: "T", Dec: "T"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Lapa le obo", Feb: "Lapa le waare", Mar: "Lapa le okuni", Apr: "Lapa le ong’wan", May: "Lapa le imet", Jun: "Lapa le ile", Jul: "Lapa le sapa", Aug: "Lapa le isiet", Sep: "Lapa le saal", Oct: "Lapa le tomon", Nov: "Lapa le tomon obo", Dec: "Lapa le tomon waare"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Are", Mon: "Kun", Tue: "Ong", Wed: "Ine", Thu: "Ile", Fri: "Sap", Sat: "Kwe"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "A", Mon: "K", Tue: "O", Wed: "I", Thu: "I", Fri: "S", Sat: "K"},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "Mderot ee are", Mon: "Mderot ee kuni", Tue: "Mderot ee ong’wan", Wed: "Mderot ee inet", Thu: "Mderot ee ile", Fri: "Mderot ee sapa", Sat: "Mderot ee kwe"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "Tesiran", PM: "Teipa"},
		},
	},
}
