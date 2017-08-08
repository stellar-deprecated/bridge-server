package kk

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "y, dd-MMM", Short: "dd/MM/yy"},
		Time:     cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
		DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "қаң.", Feb: "ақп.", Mar: "нау.", Apr: "сәу.", May: "мам.", Jun: "мау.", Jul: "шіл.", Aug: "там.", Sep: "қыр.", Oct: "қаз.", Nov: "қар.", Dec: "желт."},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "Қ", Feb: "А", Mar: "Н", Apr: "С", May: "М", Jun: "М", Jul: "Ш", Aug: "Т", Sep: "Қ", Oct: "Қ", Nov: "Қ", Dec: "Ж"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "қаңтар", Feb: "ақпан", Mar: "наурыз", Apr: "сәуір", May: "мамыр", Jun: "маусым", Jul: "шілде", Aug: "тамыз", Sep: "қыркүйек", Oct: "қазан", Nov: "қараша", Dec: "желтоқсан"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "жек", Mon: "дүй", Tue: "сей", Wed: "сәр", Thu: "бей", Fri: "жұма", Sat: "сб."},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "Ж", Mon: "Д", Tue: "С", Wed: "С", Thu: "Б", Fri: "Ж", Sat: "С"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "жс", Mon: "дс", Tue: "сс", Wed: "ср", Thu: "бс", Fri: "жм", Sat: "сб"},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "жексенбі", Mon: "дүйсенбі", Tue: "сейсенбі", Wed: "сәрсенбі", Thu: "бейсенбі", Fri: "жұма", Sat: "сенбі"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "тт", PM: "кш"},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "таңертеңгі", PM: "түстен кейінгі"},
		},
	},
}
