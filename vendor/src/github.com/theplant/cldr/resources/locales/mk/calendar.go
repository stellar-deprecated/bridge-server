package mk

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, dd MMMM y", Long: "dd MMMM y", Medium: "dd.M.y", Short: "dd.M.yy"},
		Time:     cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
		DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "јан.", Feb: "фев.", Mar: "мар.", Apr: "апр.", May: "мај", Jun: "јун.", Jul: "јул.", Aug: "авг.", Sep: "септ.", Oct: "окт.", Nov: "ноем.", Dec: "дек."},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "ј", Feb: "ф", Mar: "м", Apr: "а", May: "м", Jun: "ј", Jul: "ј", Aug: "а", Sep: "с", Oct: "о", Nov: "н", Dec: "д"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "јануари", Feb: "февруари", Mar: "март", Apr: "април", May: "мај", Jun: "јуни", Jul: "јули", Aug: "август", Sep: "септември", Oct: "октомври", Nov: "ноември", Dec: "декември"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "нед.", Mon: "пон.", Tue: "вт.", Wed: "сре.", Thu: "чет.", Fri: "пет.", Sat: "саб."},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "н", Mon: "п", Tue: "в", Wed: "с", Thu: "ч", Fri: "п", Sat: "с"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "нед.", Mon: "пон.", Tue: "вто.", Wed: "сре.", Thu: "чет.", Fri: "пет.", Sat: "саб."},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "недела", Mon: "понеделник", Tue: "вторник", Wed: "среда", Thu: "четврток", Fri: "петок", Sat: "сабота"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "прет.", PM: "попл."},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "претпладне", PM: "попладне"},
		},
	},
}
