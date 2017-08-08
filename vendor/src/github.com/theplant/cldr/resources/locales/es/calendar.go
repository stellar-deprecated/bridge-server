package es

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, d 'de' MMMM 'de' y", Long: "d 'de' MMMM 'de' y", Medium: "d MMM y", Short: "d/M/yy"},
		Time:     cldr.CalendarDateFormat{Full: "H:mm:ss (zzzz)", Long: "H:mm:ss z", Medium: "H:mm:ss", Short: "H:mm"},
		DateTime: cldr.CalendarDateFormat{Full: "{1}, {0}", Long: "{1}, {0}", Medium: "{1} {0}", Short: "{1} {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Ene.", Feb: "Feb.", Mar: "Mar.", Apr: "Abr.", May: "May.", Jun: "Jun.", Jul: "Jul.", Aug: "Ago.", Sep: "Sept.", Oct: "Oct.", Nov: "Nov.", Dec: "Dic."},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "E", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Enero", Feb: "Febrero", Mar: "Marzo", Apr: "Abril", May: "Mayo", Jun: "Junio", Jul: "Julio", Aug: "Agosto", Sep: "Septiembre", Oct: "Octubre", Nov: "Noviembre", Dec: "Diciembre"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Dom.", Mon: "Lun.", Tue: "Mar.", Wed: "Mié.", Thu: "Jue.", Fri: "Vie.", Sat: "Sáb."},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "D", Mon: "L", Tue: "M", Wed: "X", Thu: "J", Fri: "V", Sat: "S"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "DO", Mon: "LU", Tue: "MA", Wed: "MI", Thu: "JU", Fri: "VI", Sat: "SA"},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "Domingo", Mon: "Lunes", Tue: "Martes", Wed: "Miércoles", Thu: "Jueves", Fri: "Viernes", Sat: "Sábado"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "a.m.", PM: "p.m."},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "a. m.", PM: "p. m."},
		},
	},
}
