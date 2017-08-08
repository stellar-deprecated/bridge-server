package qu

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, d MMMM, y", Long: "", Medium: "", Short: "dd/MM/y"},
		Time:     cldr.CalendarDateFormat{Full: "", Long: "", Medium: "hh:mm:ss a", Short: "hh:mm a"},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Qul", Feb: "Hat", Mar: "Pau", Apr: "Ayr", May: "Aym", Jun: "Int", Jul: "Ant", Aug: "Qha", Sep: "Uma", Oct: "Kan", Nov: "Aya", Dec: "Kap"},
			Narrow:      cldr.CalendarMonthFormatNameValue{},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Qulla puquy", Feb: "Hatun puquy", Mar: "Pauqar waray", Apr: "Ayriwa", May: "Aymuray", Jun: "Inti raymi", Jul: "Anta Sitwa", Aug: "Qhapaq Sitwa", Sep: "Uma raymi", Oct: "Kantaray", Nov: "Ayamarqʼa", Dec: "Kapaq Raymi"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Dom", Mon: "Lun", Tue: "Mar", Wed: "Mié", Thu: "Jue", Fri: "Vie", Sat: "Sab"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "D", Mon: "L", Tue: "M", Wed: "X", Thu: "J", Fri: "V", Sat: "S"},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "Domingo", Mon: "Lunes", Tue: "Martes", Wed: "Miércoles", Thu: "Jueves", Fri: "Viernes", Sat: "Sábado"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "a.m.", PM: "p.m."},
		},
	},
}
