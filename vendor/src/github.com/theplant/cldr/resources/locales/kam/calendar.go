package kam

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"},
		Time:     cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Mbe", Feb: "Kel", Mar: "Ktũ", Apr: "Kan", May: "Ktn", Jun: "Tha", Jul: "Moo", Aug: "Nya", Sep: "Knd", Oct: "Ĩku", Nov: "Ĩkm", Dec: "Ĩkl"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "M", Feb: "K", Mar: "K", Apr: "K", May: "K", Jun: "T", Jul: "M", Aug: "N", Sep: "K", Oct: "Ĩ", Nov: "Ĩ", Dec: "Ĩ"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Mwai wa mbee", Feb: "Mwai wa kelĩ", Mar: "Mwai wa katatũ", Apr: "Mwai wa kana", May: "Mwai wa katano", Jun: "Mwai wa thanthatũ", Jul: "Mwai wa muonza", Aug: "Mwai wa nyaanya", Sep: "Mwai wa kenda", Oct: "Mwai wa ĩkumi", Nov: "Mwai wa ĩkumi na ĩmwe", Dec: "Mwai wa ĩkumi na ilĩ"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Wky", Mon: "Wkw", Tue: "Wkl", Wed: "Wtũ", Thu: "Wkn", Fri: "Wtn", Sat: "Wth"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "Y", Mon: "W", Tue: "E", Wed: "A", Thu: "A", Fri: "A", Sat: "A"},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "Wa kyumwa", Mon: "Wa kwambĩlĩlya", Tue: "Wa kelĩ", Wed: "Wa katatũ", Thu: "Wa kana", Fri: "Wa katano", Sat: "Wa thanthatũ"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "Ĩyakwakya", PM: "Ĩyawĩoo"},
		},
	},
}
