package pt

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, d 'de' MMMM 'de' y", Long: "d 'de' MMMM 'de' y", Medium: "d 'de' MMM 'de' y", Short: "dd/MM/yy"},
		Time:     cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
		DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "jan", Feb: "fev", Mar: "mar", Apr: "abr", May: "mai", Jun: "jun", Jul: "jul", Aug: "ago", Sep: "set", Oct: "out", Nov: "nov", Dec: "dez"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "janeiro", Feb: "fevereiro", Mar: "março", Apr: "abril", May: "maio", Jun: "junho", Jul: "julho", Aug: "agosto", Sep: "setembro", Oct: "outubro", Nov: "novembro", Dec: "dezembro"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "dom", Mon: "seg", Tue: "ter", Wed: "qua", Thu: "qui", Fri: "sex", Sat: "sáb"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "D", Mon: "S", Tue: "T", Wed: "Q", Thu: "Q", Fri: "S", Sat: "S"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "dom", Mon: "seg", Tue: "ter", Wed: "qua", Thu: "qui", Fri: "sex", Sat: "sáb"},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "domingo", Mon: "segunda-feira", Tue: "terça-feira", Wed: "quarta-feira", Thu: "quinta-feira", Fri: "sexta-feira", Sat: "sábado"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "a", PM: "p"},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"},
		},
	},
}
