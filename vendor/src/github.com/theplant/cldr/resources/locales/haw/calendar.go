package haw

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "d/M/yy"},
		Time:     cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Ian.", Feb: "Pep.", Mar: "Mal.", Apr: "ʻAp.", May: "Mei", Jun: "Iun.", Jul: "Iul.", Aug: "ʻAu.", Sep: "Kep.", Oct: "ʻOk.", Nov: "Now.", Dec: "Kek."},
			Narrow:      cldr.CalendarMonthFormatNameValue{},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Ianuali", Feb: "Pepeluali", Mar: "Malaki", Apr: "ʻApelila", May: "Mei", Jun: "Iune", Jul: "Iulai", Aug: "ʻAukake", Sep: "Kepakemapa", Oct: "ʻOkakopa", Nov: "Nowemapa", Dec: "Kekemapa"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "LP", Mon: "P1", Tue: "P2", Wed: "P3", Thu: "P4", Fri: "P5", Sat: "P6"},
			Narrow:      cldr.CalendarDayFormatNameValue{},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "Lāpule", Mon: "Poʻakahi", Tue: "Poʻalua", Wed: "Poʻakolu", Thu: "Poʻahā", Fri: "Poʻalima", Sat: "Poʻaono"},
		},
		Periods: cldr.CalendarPeriodFormatNames{},
	},
}
