package ru

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, d MMMM y 'г'.", Long: "d MMMM y 'г'.", Medium: "d MMM y 'г'.", Short: "dd.MM.yy"},
		Time:     cldr.CalendarDateFormat{Full: "H:mm:ss zzzz", Long: "H:mm:ss z", Medium: "H:mm:ss", Short: "H:mm"},
		DateTime: cldr.CalendarDateFormat{Full: "{1}, {0}", Long: "{1}, {0}", Medium: "{1}, {0}", Short: "{1}, {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "янв.", Feb: "февр.", Mar: "март", Apr: "апр.", May: "май", Jun: "июнь", Jul: "июль", Aug: "авг.", Sep: "сент.", Oct: "окт.", Nov: "нояб.", Dec: "дек."},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "Я", Feb: "Ф", Mar: "М", Apr: "А", May: "М", Jun: "И", Jul: "И", Aug: "А", Sep: "С", Oct: "О", Nov: "Н", Dec: "Д"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "январь", Feb: "февраль", Mar: "март", Apr: "апрель", May: "май", Jun: "июнь", Jul: "июль", Aug: "август", Sep: "сентябрь", Oct: "октябрь", Nov: "ноябрь", Dec: "декабрь"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Вс", Mon: "Пн", Tue: "Вт", Wed: "Ср", Thu: "Чт", Fri: "Пт", Sat: "Сб"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "В", Mon: "П", Tue: "В", Wed: "С", Thu: "Ч", Fri: "П", Sat: "С"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "вс", Mon: "пн", Tue: "вт", Wed: "ср", Thu: "чт", Fri: "пт", Sat: "сб"},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "Воскресенье", Mon: "Понедельник", Tue: "Вторник", Wed: "Среда", Thu: "Четверг", Fri: "Пятница", Sat: "Суббота"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"},
		},
	},
}
