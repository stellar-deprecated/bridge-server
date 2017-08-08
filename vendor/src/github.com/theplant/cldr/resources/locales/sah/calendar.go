package sah

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "y 'сыл' MMMM d 'күнэ', EEEE", Long: "y, MMMM d", Medium: "y, MMM d", Short: "yy/M/d"},
		Time:     cldr.CalendarDateFormat{Full: "", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: ""},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Тохс", Feb: "Олун", Mar: "Клн_ттр", Apr: "Мус_уст", May: "Ыам_йн", Jun: "Бэс_йн", Jul: "От_йн", Aug: "Атрдь_йн", Sep: "Блҕн_йн", Oct: "Алт", Nov: "Сэт", Dec: "Ахс"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "Т", Feb: "О", Mar: "К", Apr: "М", May: "Ы", Jun: "Б", Jul: "О", Aug: "А", Sep: "Б", Oct: "А", Nov: "С", Dec: "А"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Тохсунньу", Feb: "Олунньу", Mar: "Кулун тутар", Apr: "Муус устар", May: "Ыам ыйын", Jun: "Бэс ыйын", Jul: "От ыйын", Aug: "Атырдьых ыйын", Sep: "Балаҕан ыйын", Oct: "Алтынньы", Nov: "Сэтинньи", Dec: "Ахсынньы"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Бс", Mon: "Бн", Tue: "Оп", Wed: "Сэ", Thu: "Чп", Fri: "Бэ", Sat: "Сб"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "Б", Mon: "Б", Tue: "О", Wed: "С", Thu: "Ч", Fri: "Б", Sat: "С"},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "Баскыһыанньа", Mon: "Бэнидиэлинньик", Tue: "Оптуорунньук", Wed: "Сэрэдэ", Thu: "Чэппиэр", Fri: "Бээтиҥсэ", Sat: "Субуота"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "ЭИ", PM: "ЭК"},
		},
	},
}
