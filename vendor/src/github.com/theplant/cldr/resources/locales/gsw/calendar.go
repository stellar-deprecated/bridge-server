package gsw

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, d. MMMM y", Long: "d. MMMM y", Medium: "dd.MM.y", Short: "dd.MM.yy"},
		Time:     cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Jan", Feb: "Feb", Mar: "Mär", Apr: "Apr", May: "Mai", Jun: "Jun", Jul: "Jul", Aug: "Aug", Sep: "Sep", Oct: "Okt", Nov: "Nov", Dec: "Dez"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Januar", Feb: "Februar", Mar: "März", Apr: "April", May: "Mai", Jun: "Juni", Jul: "Juli", Aug: "Auguscht", Sep: "Septämber", Oct: "Oktoober", Nov: "Novämber", Dec: "Dezämber"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Su.", Mon: "Mä.", Tue: "Zi.", Wed: "Mi.", Thu: "Du.", Fri: "Fr.", Sat: "Sa."},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "S", Mon: "M", Tue: "D", Wed: "M", Thu: "D", Fri: "F", Sat: "S"},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "Sunntig", Mon: "Määntig", Tue: "Ziischtig", Wed: "Mittwuch", Thu: "Dunschtig", Fri: "Friitig", Sat: "Samschtig"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "v.m.", PM: "n.m."},
			Narrow:      cldr.CalendarPeriodFormatNameValue{},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "Vormittag", PM: "Namittag"},
		},
	},
}
