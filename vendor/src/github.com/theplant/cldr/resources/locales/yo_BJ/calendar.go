package yo_BJ

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"},
		Time:     cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Shɛ́rɛ́", Feb: "Èrèlè", Mar: "Ɛrɛ̀nà", Apr: "Ìgbé", May: "Ɛ̀bibi", Jun: "Òkúdu", Jul: "Agɛmɔ", Aug: "Ògún", Sep: "Owewe", Oct: "Ɔ̀wàrà", Nov: "Bélú", Dec: "Ɔ̀pɛ̀"},
			Narrow:      cldr.CalendarMonthFormatNameValue{},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Oshù Shɛ́rɛ́", Feb: "Oshù Èrèlè", Mar: "Oshù Ɛrɛ̀nà", Apr: "Oshù Ìgbé", May: "Oshù Ɛ̀bibi", Jun: "Oshù Òkúdu", Jul: "Oshù Agɛmɔ", Aug: "Oshù Ògún", Sep: "Oshù Owewe", Oct: "Oshù Ɔ̀wàrà", Nov: "Oshù Bélú", Dec: "Oshù Ɔ̀pɛ̀"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Àìkú", Mon: "Ajé", Tue: "Ìsɛ́gun", Wed: "Ɔjɔ́rú", Thu: "Ɔjɔ́bɔ", Fri: "Ɛtì", Sat: "Àbámɛ́ta"},
			Narrow:      cldr.CalendarDayFormatNameValue{},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "Ɔjɔ́ Àìkú", Mon: "Ɔjɔ́ Ajé", Tue: "Ɔjɔ́ Ìsɛ́gun", Wed: "Ɔjɔ́rú", Thu: "Ɔjɔ́bɔ", Fri: "Ɔjɔ́ Ɛtì", Sat: "Ɔjɔ́ Àbámɛ́ta"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "Àárɔ̀", PM: "Ɔ̀sán"},
		},
	},
}
