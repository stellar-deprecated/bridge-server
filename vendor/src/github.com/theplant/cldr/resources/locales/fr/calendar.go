package fr

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"},
		Time:     cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
		DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Janv.", Feb: "Févr.", Mar: "Mars", Apr: "Avr.", May: "Mai", Jun: "Juin", Jul: "Juil.", Aug: "Août", Sep: "Sept.", Oct: "Oct.", Nov: "Nov.", Dec: "Déc."},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Janvier", Feb: "Février", Mar: "Mars", Apr: "Avril", May: "Mai", Jun: "Juin", Jul: "Juillet", Aug: "Août", Sep: "Septembre", Oct: "Octobre", Nov: "Novembre", Dec: "Décembre"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Dim.", Mon: "Lun.", Tue: "Mar.", Wed: "Mer.", Thu: "Jeu.", Fri: "Ven.", Sat: "Sam."},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "D", Mon: "L", Tue: "M", Wed: "M", Thu: "J", Fri: "V", Sat: "S"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "Di", Mon: "Lu", Tue: "Ma", Wed: "Me", Thu: "Je", Fri: "Ve", Sat: "Sa"},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "Dimanche", Mon: "Lundi", Tue: "Mardi", Wed: "Mercredi", Thu: "Jeudi", Fri: "Vendredi", Sat: "Samedi"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "av.m.", PM: "ap.m."},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "a", PM: "p"},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "avant-midi", PM: "après-midi"},
		},
	},
}
