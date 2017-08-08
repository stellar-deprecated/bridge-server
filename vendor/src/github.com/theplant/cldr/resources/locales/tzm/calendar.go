package tzm

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"},
		Time:     cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Yen", Feb: "Yeb", Mar: "Mar", Apr: "Ibr", May: "May", Jun: "Yun", Jul: "Yul", Aug: "Ɣuc", Sep: "Cut", Oct: "Kṭu", Nov: "Nwa", Dec: "Duj"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "Y", Feb: "Y", Mar: "M", Apr: "I", May: "M", Jun: "Y", Jul: "Y", Aug: "Ɣ", Sep: "C", Oct: "K", Nov: "N", Dec: "D"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Yennayer", Feb: "Yebrayer", Mar: "Mars", Apr: "Ibrir", May: "Mayyu", Jun: "Yunyu", Jul: "Yulyuz", Aug: "Ɣuct", Sep: "Cutanbir", Oct: "Kṭuber", Nov: "Nwanbir", Dec: "Dujanbir"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Asa", Mon: "Ayn", Tue: "Asn", Wed: "Akr", Thu: "Akw", Fri: "Asm", Sat: "Asḍ"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "A", Mon: "A", Tue: "A", Wed: "A", Thu: "A", Fri: "A", Sat: "A"},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "Asamas", Mon: "Aynas", Tue: "Asinas", Wed: "Akras", Thu: "Akwas", Fri: "Asimwas", Sat: "Asiḍyas"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "Zdat azal", PM: "Ḍeffir aza"},
		},
	},
}
