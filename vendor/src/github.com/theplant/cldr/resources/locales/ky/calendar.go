package ky

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, d-MMMM, y-'ж'.", Long: "y MMMM d", Medium: "y MMM d", Short: "dd.MM.yy"},
		Time:     cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
		DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Янв", Feb: "Фев", Mar: "Мар", Apr: "Апр", May: "Май", Jun: "Июн", Jul: "Июл", Aug: "Авг", Sep: "Сен", Oct: "Окт", Nov: "Ноя", Dec: "Дек"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "Я", Feb: "Ф", Mar: "М", Apr: "А", May: "М", Jun: "И", Jul: "И", Aug: "А", Sep: "С", Oct: "О", Nov: "Н", Dec: "Д"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Январь", Feb: "Февраль", Mar: "Март", Apr: "Апрель", May: "Май", Jun: "Июнь", Jul: "Июль", Aug: "Август", Sep: "Сентябрь", Oct: "Октябрь", Nov: "Ноябрь", Dec: "Декабрь"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "жек.", Mon: "дүй.", Tue: "шейш.", Wed: "шарш.", Thu: "бейш.", Fri: "жума", Sat: "ишм."},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "Ж", Mon: "Д", Tue: "Ш", Wed: "Ш", Thu: "Б", Fri: "Ж", Sat: "И"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "жк", Mon: "дш.", Tue: "шш.", Wed: "шр.", Thu: "бш.", Fri: "жм.", Sat: "иш."},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "жекшемби", Mon: "дүйшөмбү", Tue: "шейшемби", Wed: "шаршемби", Thu: "бейшемби", Fri: "жума", Sat: "ишемби"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "тң", PM: "тк"},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "таңкы", PM: "түштөн кийин"},
		},
	},
}
