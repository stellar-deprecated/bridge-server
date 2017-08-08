package sr

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, dd. MMMM y.", Long: "dd. MMMM y.", Medium: "dd.MM.y.", Short: "d.M.yy."},
		Time:     cldr.CalendarDateFormat{Full: "HH.mm.ss zzzz", Long: "HH.mm.ss z", Medium: "HH.mm.ss", Short: "HH.mm"},
		DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "јан", Feb: "феб", Mar: "мар", Apr: "апр", May: "мај", Jun: "јун", Jul: "јул", Aug: "авг", Sep: "сеп", Oct: "окт", Nov: "нов", Dec: "дец"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "ј", Feb: "ф", Mar: "м", Apr: "а", May: "м", Jun: "ј", Jul: "ј", Aug: "а", Sep: "с", Oct: "о", Nov: "н", Dec: "д"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "јануар", Feb: "фебруар", Mar: "март", Apr: "април", May: "мај", Jun: "јун", Jul: "јул", Aug: "август", Sep: "септембар", Oct: "октобар", Nov: "новембар", Dec: "децембар"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "нед", Mon: "пон", Tue: "уто", Wed: "сре", Thu: "чет", Fri: "пет", Sat: "суб"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "н", Mon: "п", Tue: "у", Wed: "с", Thu: "ч", Fri: "п", Sat: "с"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "не", Mon: "по", Tue: "ут", Wed: "ср", Thu: "че", Fri: "пе", Sat: "су"},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "недеља", Mon: "понедељак", Tue: "уторак", Wed: "среда", Thu: "четвртак", Fri: "петак", Sat: "субота"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "a", PM: "p"},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "пре подне", PM: "по подне"},
		},
	},
}
