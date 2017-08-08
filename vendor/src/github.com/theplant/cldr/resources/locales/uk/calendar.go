package uk

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, d MMMM y 'р'.", Long: "d MMMM y 'р'.", Medium: "d MMM y 'р'.", Short: "dd.MM.yy"},
		Time:     cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
		DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Січ", Feb: "Лют", Mar: "Бер", Apr: "Кві", May: "Тра", Jun: "Чер", Jul: "Лип", Aug: "Сер", Sep: "Вер", Oct: "Жов", Nov: "Лис", Dec: "Гру"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "С", Feb: "Л", Mar: "Б", Apr: "К", May: "Т", Jun: "Ч", Jul: "Л", Aug: "С", Sep: "В", Oct: "Ж", Nov: "Л", Dec: "Г"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Січень", Feb: "Лютий", Mar: "Березень", Apr: "Квітень", May: "Травень", Jun: "Червень", Jul: "Липень", Aug: "Серпень", Sep: "Вересень", Oct: "Жовтень", Nov: "Листопад", Dec: "Грудень"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Нд", Mon: "Пн", Tue: "Вт", Wed: "Ср", Thu: "Чт", Fri: "Пт", Sat: "Сб"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "Н", Mon: "П", Tue: "В", Wed: "С", Thu: "Ч", Fri: "П", Sat: "С"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "Нд", Mon: "Пн", Tue: "Вт", Wed: "Ср", Thu: "Чт", Fri: "Пт", Sat: "Сб"},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "Неділя", Mon: "Понеділок", Tue: "Вівторок", Wed: "Середа", Thu: "Четвер", Fri: "Пʼятниця", Sat: "Субота"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "дп", PM: "пп"},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "дп", PM: "пп"},
		},
	},
}
