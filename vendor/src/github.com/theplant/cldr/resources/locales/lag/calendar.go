package lag

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"},
		Time:     cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Fúngatɨ", Feb: "Naanɨ", Mar: "Keenda", Apr: "Ikúmi", May: "Inyambala", Jun: "Idwaata", Jul: "Mʉʉnchɨ", Aug: "Vɨɨrɨ", Sep: "Saatʉ", Oct: "Inyi", Nov: "Saano", Dec: "Sasatʉ"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "F", Feb: "N", Mar: "K", Apr: "I", May: "I", Jun: "I", Jul: "M", Aug: "V", Sep: "S", Oct: "I", Nov: "S", Dec: "S"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Kʉfúngatɨ", Feb: "Kʉnaanɨ", Mar: "Kʉkeenda", Apr: "Kwiikumi", May: "Kwiinyambála", Jun: "Kwiidwaata", Jul: "Kʉmʉʉnchɨ", Aug: "Kʉvɨɨrɨ", Sep: "Kʉsaatʉ", Oct: "Kwiinyi", Nov: "Kʉsaano", Dec: "Kʉsasatʉ"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Píili", Mon: "Táatu", Tue: "Íne", Wed: "Táano", Thu: "Alh", Fri: "Ijm", Sat: "Móosi"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "P", Mon: "T", Tue: "E", Wed: "O", Thu: "A", Fri: "I", Sat: "M"},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "Jumapíiri", Mon: "Jumatátu", Tue: "Jumaíne", Wed: "Jumatáano", Thu: "Alamíisi", Fri: "Ijumáa", Sat: "Jumamóosi"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "TOO", PM: "MUU"},
		},
	},
}
