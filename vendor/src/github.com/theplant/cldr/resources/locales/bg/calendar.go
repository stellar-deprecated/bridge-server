package bg

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, d MMMM y 'г'.", Long: "d MMMM y 'г'.", Medium: "d.MM.y 'г'.", Short: "d.MM.yy 'г'."},
		Time:     cldr.CalendarDateFormat{Full: "H:mm:ss zzzz", Long: "H:mm:ss z", Medium: "H:mm:ss", Short: "H:mm"},
		DateTime: cldr.CalendarDateFormat{Full: "{1}, {0}", Long: "{1}, {0}", Medium: "{1}, {0}", Short: "{1}, {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "ян.", Feb: "февр.", Mar: "март", Apr: "апр.", May: "май", Jun: "юни", Jul: "юли", Aug: "авг.", Sep: "септ.", Oct: "окт.", Nov: "ноем.", Dec: "дек."},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "я", Feb: "ф", Mar: "м", Apr: "а", May: "м", Jun: "ю", Jul: "ю", Aug: "а", Sep: "с", Oct: "о", Nov: "н", Dec: "д"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "януари", Feb: "февруари", Mar: "март", Apr: "април", May: "май", Jun: "юни", Jul: "юли", Aug: "август", Sep: "септември", Oct: "октомври", Nov: "ноември", Dec: "декември"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "нд", Mon: "пн", Tue: "вт", Wed: "ср", Thu: "чт", Fri: "пт", Sat: "сб"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "н", Mon: "п", Tue: "в", Wed: "с", Thu: "ч", Fri: "п", Sat: "с"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "нд", Mon: "пн", Tue: "вт", Wed: "ср", Thu: "чт", Fri: "пт", Sat: "сб"},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "неделя", Mon: "понеделник", Tue: "вторник", Wed: "сряда", Thu: "четвъртък", Fri: "петък", Sat: "събота"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "пр.об.", PM: "сл.об."},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "пр.об.", PM: "сл.об."},
		},
	},
}
