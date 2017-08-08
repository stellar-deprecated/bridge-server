package teo

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"},
		Time:     cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Rar", Feb: "Muk", Mar: "Kwa", Apr: "Dun", May: "Mar", Jun: "Mod", Jul: "Jol", Aug: "Ped", Sep: "Sok", Oct: "Tib", Nov: "Lab", Dec: "Poo"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "R", Feb: "M", Mar: "K", Apr: "D", May: "M", Jun: "M", Jul: "J", Aug: "P", Sep: "S", Oct: "T", Nov: "L", Dec: "P"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Orara", Feb: "Omuk", Mar: "Okwamg’", Apr: "Odung’el", May: "Omaruk", Jun: "Omodok’king’ol", Jul: "Ojola", Aug: "Opedel", Sep: "Osokosokoma", Oct: "Otibar", Nov: "Olabor", Dec: "Opoo"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Jum", Mon: "Bar", Tue: "Aar", Wed: "Uni", Thu: "Ung", Fri: "Kan", Sat: "Sab"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "J", Mon: "B", Tue: "A", Wed: "U", Thu: "U", Fri: "K", Sat: "S"},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "Nakaejuma", Mon: "Nakaebarasa", Tue: "Nakaare", Wed: "Nakauni", Thu: "Nakaung’on", Fri: "Nakakany", Sat: "Nakasabiti"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "Taparachu", PM: "Ebongi"},
		},
	},
}
