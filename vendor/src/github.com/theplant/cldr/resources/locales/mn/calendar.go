package mn

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, y 'оны' MM 'сарын' d", Long: "y 'оны' MM 'сарын' d", Medium: "y MMM d", Short: "y-MM-dd"},
		Time:     cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
		DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1}, {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "1-р сар", Feb: "2-р сар", Mar: "3-р сар", Apr: "4-р сар", May: "5-р сар", Jun: "6-р сар", Jul: "7-р сар", Aug: "8-р сар", Sep: "9-р сар", Oct: "10-р сар", Nov: "11-р сар", Dec: "12-р сар"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "1", Feb: "2", Mar: "3", Apr: "4", May: "5", Jun: "6", Jul: "7", Aug: "8", Sep: "9", Oct: "10", Nov: "11", Dec: "12"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Нэгдүгээр сар", Feb: "Хоёрдугаар сар", Mar: "Гуравдугаар сар", Apr: "Дөрөвдүгээр сар", May: "Тавдугаар сар", Jun: "Зургадугаар сар", Jul: "Долдугаар сар", Aug: "Наймдугаар сар", Sep: "Есдүгээр сар", Oct: "Аравдугаар сар", Nov: "Арван нэгдүгээр сар", Dec: "Арван хоёрдугаар сар"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Ня", Mon: "Да", Tue: "Мя", Wed: "Лх", Thu: "Пү", Fri: "Ба", Sat: "Бя"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "1", Mon: "2", Tue: "3", Wed: "4", Thu: "5", Fri: "6", Sat: "7"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "Ня", Mon: "Да", Tue: "Мя", Wed: "Лх", Thu: "Пү", Fri: "Ба", Sat: "Бя"},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "ням", Mon: "даваа", Tue: "мягмар", Wed: "лхагва", Thu: "пүрэв", Fri: "баасан", Sat: "бямба"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "үө", PM: "үх"},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "ҮӨ", PM: "ҮХ"},
		},
	},
}
