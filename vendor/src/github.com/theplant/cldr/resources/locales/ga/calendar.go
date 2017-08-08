package ga

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"},
		Time:     cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
		DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Ean", Feb: "Feabh", Mar: "Márta", Apr: "Aib", May: "Beal", Jun: "Meith", Jul: "Iúil", Aug: "Lún", Sep: "MFómh", Oct: "DFómh", Nov: "Samh", Dec: "Noll"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "E", Feb: "F", Mar: "M", Apr: "A", May: "B", Jun: "M", Jul: "I", Aug: "L", Sep: "M", Oct: "D", Nov: "S", Dec: "N"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Eanáir", Feb: "Feabhra", Mar: "Márta", Apr: "Aibreán", May: "Bealtaine", Jun: "Meitheamh", Jul: "Iúil", Aug: "Lúnasa", Sep: "Meán Fómhair", Oct: "Deireadh Fómhair", Nov: "Samhain", Dec: "Nollaig"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Domh", Mon: "Luan", Tue: "Máirt", Wed: "Céad", Thu: "Déar", Fri: "Aoine", Sat: "Sath"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "D", Mon: "L", Tue: "M", Wed: "C", Thu: "D", Fri: "A", Sat: "S"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "Do", Mon: "Lu", Tue: "Má", Wed: "Cé", Thu: "Dé", Fri: "Ao", Sat: "Sa"},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "Dé Domhnaigh", Mon: "Dé Luain", Tue: "Dé Máirt", Wed: "Dé Céadaoin", Thu: "Déardaoin", Fri: "Dé hAoine", Sat: "Dé Sathairn"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "a", PM: "p"},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "a.m.", PM: "p.m."},
		},
	},
}
