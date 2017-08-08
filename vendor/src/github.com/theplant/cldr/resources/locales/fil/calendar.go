package fil

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, MMMM d, y", Long: "MMMM d, y", Medium: "MMM d, y", Short: "M/d/yy"},
		Time:     cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
		DateTime: cldr.CalendarDateFormat{Full: "{1} 'nang' {0}", Long: "{1} 'nang' {0}", Medium: "{1}, {0}", Short: "{1}, {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Ene", Feb: "Peb", Mar: "Mar", Apr: "Abr", May: "May", Jun: "Hun", Jul: "Hul", Aug: "Ago", Sep: "Set", Oct: "Okt", Nov: "Nob", Dec: "Dis"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "E", Feb: "P", Mar: "M", Apr: "A", May: "M", Jun: "H", Jul: "H", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Enero", Feb: "Pebrero", Mar: "Marso", Apr: "Abril", May: "Mayo", Jun: "Hunyo", Jul: "Hulyo", Aug: "Agosto", Sep: "Setyembre", Oct: "Oktubre", Nov: "Nobyembre", Dec: "Disyembre"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Lin", Mon: "Lun", Tue: "Mar", Wed: "Miy", Thu: "Huw", Fri: "Biy", Sat: "Sab"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "L", Mon: "L", Tue: "M", Wed: "M", Thu: "H", Fri: "B", Sat: "S"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "Li", Mon: "Lu", Tue: "Ma", Wed: "Mi", Thu: "Hu", Fri: "Bi", Sat: "Sa"},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "Linggo", Mon: "Lunes", Tue: "Martes", Wed: "Miyerkules", Thu: "Huwebes", Fri: "Biyernes", Sat: "Sabado"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "a", PM: "p"},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"},
		},
	},
}
