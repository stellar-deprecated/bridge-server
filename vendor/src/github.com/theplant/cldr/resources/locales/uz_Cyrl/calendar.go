package uz_Cyrl

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, y MMMM dd", Long: "y MMMM d", Medium: "y MMM d", Short: "yy/MM/dd"},
		Time:     cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
		DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Янв", Feb: "Фев", Mar: "Мар", Apr: "Апр", May: "Май", Jun: "Июн", Jul: "Июл", Aug: "Авг", Sep: "Сен", Oct: "Окт", Nov: "Ноя", Dec: "Дек"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "Я", Feb: "Ф", Mar: "М", Apr: "А", May: "М", Jun: "И", Jul: "И", Aug: "А", Sep: "С", Oct: "О", Nov: "Н", Dec: "Д"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Январ", Feb: "Феврал", Mar: "Март", Apr: "Апрел", May: "Май", Jun: "Июн", Jul: "Июл", Aug: "Август", Sep: "Сентябр", Oct: "Октябр", Nov: "Ноябр", Dec: "Декабр"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Якш", Mon: "Душ", Tue: "Сеш", Wed: "Чор", Thu: "Пай", Fri: "Жум", Sat: "Шан"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "Я", Mon: "Д", Tue: "С", Wed: "Ч", Thu: "П", Fri: "Ж", Sat: "Ш"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "Якш", Mon: "Душ", Tue: "Сеш", Wed: "Чор", Thu: "Пай", Fri: "Жум", Sat: "Шан"},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "якшанба", Mon: "душанба", Tue: "сешанба", Wed: "чоршанба", Thu: "пайшанба", Fri: "жума", Sat: "шанба"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"},
		},
	},
}
