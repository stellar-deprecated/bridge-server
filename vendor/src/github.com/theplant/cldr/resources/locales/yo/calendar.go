package yo

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"},
		Time:     cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Ṣẹ́rẹ́", Feb: "Èrèlè", Mar: "Ẹrẹ̀nà", Apr: "Ìgbé", May: "Ẹ̀bibi", Jun: "Òkúdu", Jul: "Agẹmọ", Aug: "Ògún", Sep: "Owewe", Oct: "Ọ̀wàrà", Nov: "Bélú", Dec: "Ọ̀pẹ̀"},
			Narrow:      cldr.CalendarMonthFormatNameValue{},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Oṣù Ṣẹ́rẹ́", Feb: "Oṣù Èrèlè", Mar: "Oṣù Ẹrẹ̀nà", Apr: "Oṣù Ìgbé", May: "Oṣù Ẹ̀bibi", Jun: "Oṣù Òkúdu", Jul: "Oṣù Agẹmọ", Aug: "Oṣù Ògún", Sep: "Oṣù Owewe", Oct: "Oṣù Ọ̀wàrà", Nov: "Oṣù Bélú", Dec: "Oṣù Ọ̀pẹ̀"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Àìkú", Mon: "Ajé", Tue: "Ìsẹ́gun", Wed: "Ọjọ́rú", Thu: "Ọjọ́bọ", Fri: "Ẹtì", Sat: "Àbámẹ́ta"},
			Narrow:      cldr.CalendarDayFormatNameValue{},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "Ọjọ́ Àìkú", Mon: "Ọjọ́ Ajé", Tue: "Ọjọ́ Ìsẹ́gun", Wed: "Ọjọ́rú", Thu: "Ọjọ́bọ", Fri: "Ọjọ́ Ẹtì", Sat: "Ọjọ́ Àbámẹ́ta"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "Àárọ̀", PM: "Ọ̀sán"},
		},
	},
}
