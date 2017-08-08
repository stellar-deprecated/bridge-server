package os

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, d MMMM, y 'аз'", Long: "d MMMM, y 'аз'", Medium: "dd MMM y 'аз'", Short: "dd.MM.yy"},
		Time:     cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
		DateTime: cldr.CalendarDateFormat{Full: "{1}, {0}", Long: "{1}, {0}", Medium: "{1}, {0}", Short: "{1}, {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Янв.", Feb: "Февр.", Mar: "Март", Apr: "Апр.", May: "Май", Jun: "Июнь", Jul: "Июль", Aug: "Авг.", Sep: "Сент.", Oct: "Окт.", Nov: "Нояб.", Dec: "Дек."},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "Я", Feb: "Ф", Mar: "М", Apr: "А", May: "М", Jun: "И", Jul: "И", Aug: "А", Sep: "С", Oct: "О", Nov: "Н", Dec: "Д"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Январь", Feb: "Февраль", Mar: "Мартъи", Apr: "Апрель", May: "Май", Jun: "Июнь", Jul: "Июль", Aug: "Август", Sep: "Сентябрь", Oct: "Октябрь", Nov: "Ноябрь", Dec: "Декабрь"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Хцб", Mon: "Крс", Tue: "Дцг", Wed: "Ӕрт", Thu: "Цпр", Fri: "Мрб", Sat: "Сбт"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "Х", Mon: "К", Tue: "Д", Wed: "Ӕ", Thu: "Ц", Fri: "М", Sat: "С"},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "Хуыцаубон", Mon: "Къуырисӕр", Tue: "Дыццӕг", Wed: "Ӕртыццӕг", Thu: "Цыппӕрӕм", Fri: "Майрӕмбон", Sat: "Сабат"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "ӕмбисбоны размӕ", PM: "ӕмбисбоны фӕстӕ"},
		},
	},
}
