package sq

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "d.M.yy"},
		Time:     cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
		DateTime: cldr.CalendarDateFormat{Full: "{1} 'në' {0}", Long: "{1} 'në' {0}", Medium: "{1}, {0}", Short: "{1}, {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Jan", Feb: "Shk", Mar: "Mar", Apr: "Pri", May: "Maj", Jun: "Qer", Jul: "Kor", Aug: "Gsh", Sep: "Sht", Oct: "Tet", Nov: "Nën", Dec: "Dhj"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "S", Mar: "M", Apr: "P", May: "M", Jun: "Q", Jul: "K", Aug: "G", Sep: "S", Oct: "T", Nov: "N", Dec: "D"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Janar", Feb: "Shkurt", Mar: "Mars", Apr: "Prill", May: "Maj", Jun: "Qershor", Jul: "Korrik", Aug: "Gusht", Sep: "Shtator", Oct: "Tetor", Nov: "Nëntor", Dec: "Dhjetor"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Die", Mon: "Hën", Tue: "Mar", Wed: "Mër", Thu: "Enj", Fri: "Pre", Sat: "Sht"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "D", Mon: "H", Tue: "M", Wed: "M", Thu: "E", Fri: "P", Sat: "S"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "Die", Mon: "Hën", Tue: "Mar", Wed: "Mër", Thu: "Enj", Fri: "Pre", Sat: "Sht"},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "E diel", Mon: "E hënë", Tue: "E martë", Wed: "E mërkurë", Thu: "E enjte", Fri: "E premte", Sat: "E shtunë"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "paradite", PM: "pasdite"},
		},
	},
}
