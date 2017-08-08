package sbp

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"},
		Time:     cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Mup", Feb: "Mwi", Mar: "Msh", Apr: "Mun", May: "Mag", Jun: "Muj", Jul: "Msp", Aug: "Mpg", Sep: "Mye", Oct: "Mok", Nov: "Mus", Dec: "Muh"},
			Narrow:      cldr.CalendarMonthFormatNameValue{},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Mupalangulwa", Feb: "Mwitope", Mar: "Mushende", Apr: "Munyi", May: "Mushende Magali", Jun: "Mujimbi", Jul: "Mushipepo", Aug: "Mupuguto", Sep: "Munyense", Oct: "Mokhu", Nov: "Musongandembwe", Dec: "Muhaano"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Mul", Mon: "Jtt", Tue: "Jnn", Wed: "Jtn", Thu: "Alh", Fri: "Iju", Sat: "Jmo"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "M", Mon: "J", Tue: "J", Wed: "J", Thu: "A", Fri: "I", Sat: "J"},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "Mulungu", Mon: "Jumatatu", Tue: "Jumanne", Wed: "Jumatano", Thu: "Alahamisi", Fri: "Ijumaa", Sat: "Jumamosi"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "Lwamilawu", PM: "Pashamihe"},
		},
	},
}
