package lu

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "d/M/y"},
		Time:     cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Cio", Feb: "Lui", Mar: "Lus", Apr: "Muu", May: "Lum", Jun: "Luf", Jul: "Kab", Aug: "Lush", Sep: "Lut", Oct: "Lun", Nov: "Kas", Dec: "Cis"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "C", Feb: "L", Mar: "L", Apr: "M", May: "L", Jun: "L", Jul: "K", Aug: "L", Sep: "L", Oct: "L", Nov: "K", Dec: "C"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Ciongo", Feb: "Lùishi", Mar: "Lusòlo", Apr: "Mùuyà", May: "Lumùngùlù", Jun: "Lufuimi", Jul: "Kabàlàshìpù", Aug: "Lùshìkà", Sep: "Lutongolo", Oct: "Lungùdi", Nov: "Kaswèkèsè", Dec: "Ciswà"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Lum", Mon: "Nko", Tue: "Ndy", Wed: "Ndg", Thu: "Njw", Fri: "Ngv", Sat: "Lub"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "L", Mon: "N", Tue: "N", Wed: "N", Thu: "N", Fri: "N", Sat: "L"},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "Lumingu", Mon: "Nkodya", Tue: "Ndàayà", Wed: "Ndangù", Thu: "Njòwa", Fri: "Ngòvya", Sat: "Lubingu"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "Dinda", PM: "Dilolo"},
		},
	},
}
