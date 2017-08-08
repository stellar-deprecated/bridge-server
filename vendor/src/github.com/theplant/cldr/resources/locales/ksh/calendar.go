package ksh

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, 'dä' d. MMMM y", Long: "d. MMMM y", Medium: "d. MMM. y", Short: "d. M. y"},
		Time:     cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Jan.", Feb: "Fäb.", Mar: "Mäz.", Apr: "Apr.", May: "Mäi", Jun: "Jun.", Jul: "Jul.", Aug: "Ouj.", Sep: "Säp.", Oct: "Okt.", Nov: "Nov.", Dec: "Dez."},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Jannewa", Feb: "Fäbrowa", Mar: "Määz", Apr: "Aprell", May: "Mäi", Jun: "Juuni", Jul: "Juuli", Aug: "Oujoß", Sep: "Septämber", Oct: "Oktoober", Nov: "Novämber", Dec: "Dezämber"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Su.", Mon: "Mo.", Tue: "Di.", Wed: "Me.", Thu: "Du.", Fri: "Fr.", Sat: "Sa."},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "S", Mon: "M", Tue: "D", Wed: "M", Thu: "D", Fri: "F", Sat: "S"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "Su", Mon: "Mo", Tue: "Di", Wed: "Me", Thu: "Du", Fri: "Fr", Sat: "Sa"},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "Sunndaach", Mon: "Moondaach", Tue: "Dinnsdaach", Wed: "Metwoch", Thu: "Dunnersdaach", Fri: "Friidaach", Sat: "Samsdaach"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "v.m.", PM: "n.m."},
			Narrow:      cldr.CalendarPeriodFormatNameValue{},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "Vormittag", PM: "Nachmittag"},
		},
	},
}
