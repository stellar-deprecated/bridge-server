package naq

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"},
		Time:     cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Jan", Feb: "Feb", Mar: "Mar", Apr: "Apr", May: "May", Jun: "Jun", Jul: "Jul", Aug: "Aug", Sep: "Sep", Oct: "Oct", Nov: "Nov", Dec: "Dec"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "ǃKhanni", Feb: "ǃKhanǀgôab", Mar: "ǀKhuuǁkhâb", Apr: "ǃHôaǂkhaib", May: "ǃKhaitsâb", Jun: "Gamaǀaeb", Jul: "ǂKhoesaob", Aug: "Aoǁkhuumûǁkhâb", Sep: "Taraǀkhuumûǁkhâb", Oct: "ǂNûǁnâiseb", Nov: "ǀHooǂgaeb", Dec: "Hôasoreǁkhâb"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Son", Mon: "Ma", Tue: "De", Wed: "Wu", Thu: "Do", Fri: "Fr", Sat: "Sat"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "S", Mon: "M", Tue: "E", Wed: "W", Thu: "D", Fri: "F", Sat: "A"},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "Sontaxtsees", Mon: "Mantaxtsees", Tue: "Denstaxtsees", Wed: "Wunstaxtsees", Thu: "Dondertaxtsees", Fri: "Fraitaxtsees", Sat: "Satertaxtsees"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "ǁgoagas", PM: "ǃuias"},
		},
	},
}
