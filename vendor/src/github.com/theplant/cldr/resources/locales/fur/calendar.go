package fur

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE d 'di' MMMM 'dal' y", Long: "d 'di' MMMM 'dal' y", Medium: "dd/MM/y", Short: "dd/MM/yy"},
		Time:     cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Zen", Feb: "Fev", Mar: "Mar", Apr: "Avr", May: "Mai", Jun: "Jug", Jul: "Lui", Aug: "Avo", Sep: "Set", Oct: "Otu", Nov: "Nov", Dec: "Dic"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "Z", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "L", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Zenâr", Feb: "Fevrâr", Mar: "Març", Apr: "Avrîl", May: "Mai", Jun: "Jugn", Jul: "Lui", Aug: "Avost", Sep: "Setembar", Oct: "Otubar", Nov: "Novembar", Dec: "Dicembar"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "dom", Mon: "lun", Tue: "mar", Wed: "mie", Thu: "joi", Fri: "vin", Sat: "sab"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "D", Mon: "L", Tue: "M", Wed: "M", Thu: "J", Fri: "V", Sat: "S"},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "domenie", Mon: "lunis", Tue: "martars", Wed: "miercus", Thu: "joibe", Fri: "vinars", Sat: "sabide"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "a.", PM: "p."},
		},
	},
}
