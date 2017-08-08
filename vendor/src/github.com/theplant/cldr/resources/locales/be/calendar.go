package be

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d.M.y", Short: "d.M.yy"},
		Time:     cldr.CalendarDateFormat{Full: "HH.mm.ss zzzz", Long: "HH.mm.ss z", Medium: "HH.mm.ss", Short: "HH.mm"},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "сту", Feb: "лют", Mar: "сак", Apr: "кра", May: "май", Jun: "чэр", Jul: "ліп", Aug: "жні", Sep: "вер", Oct: "кас", Nov: "ліс", Dec: "сне"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "с", Feb: "л", Mar: "с", Apr: "к", May: "м", Jun: "ч", Jul: "л", Aug: "ж", Sep: "в", Oct: "к", Nov: "л", Dec: "с"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "студзень", Feb: "люты", Mar: "сакавік", Apr: "красавік", May: "май", Jun: "чэрвень", Jul: "ліпень", Aug: "жнівень", Sep: "верасень", Oct: "кастрычнік", Nov: "лістапад", Dec: "снежань"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "нд", Mon: "пн", Tue: "аў", Wed: "ср", Thu: "чц", Fri: "пт", Sat: "сб"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "н", Mon: "п", Tue: "а", Wed: "с", Thu: "ч", Fri: "п", Sat: "с"},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "нядзеля", Mon: "панядзелак", Tue: "аўторак", Wed: "серада", Thu: "чацвер", Fri: "пятніца", Sat: "субота"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "раніцы", PM: "вечара"},
			Narrow:      cldr.CalendarPeriodFormatNameValue{},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "да палудня", PM: "пасля палудня"},
		},
	},
}
