package ak

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, y MMMM dd", Long: "y MMMM d", Medium: "y MMM d", Short: "yy/MM/dd"},
		Time:     cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "S-Ɔ", Feb: "K-Ɔ", Mar: "E-Ɔ", Apr: "E-O", May: "E-K", Jun: "O-A", Jul: "A-K", Aug: "D-Ɔ", Sep: "F-Ɛ", Oct: "Ɔ-A", Nov: "Ɔ-O", Dec: "M-Ɔ"},
			Narrow:      cldr.CalendarMonthFormatNameValue{},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Sanda-Ɔpɛpɔn", Feb: "Kwakwar-Ɔgyefuo", Mar: "Ebɔw-Ɔbenem", Apr: "Ebɔbira-Oforisuo", May: "Esusow Aketseaba-Kɔtɔnimba", Jun: "Obirade-Ayɛwohomumu", Jul: "Ayɛwoho-Kitawonsa", Aug: "Difuu-Ɔsandaa", Sep: "Fankwa-Ɛbɔ", Oct: "Ɔbɛsɛ-Ahinime", Nov: "Ɔberɛfɛw-Obubuo", Dec: "Mumu-Ɔpɛnimba"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Kwe", Mon: "Dwo", Tue: "Ben", Wed: "Wuk", Thu: "Yaw", Fri: "Fia", Sat: "Mem"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "K", Mon: "D", Tue: "B", Wed: "W", Thu: "Y", Fri: "F", Sat: "M"},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "Kwesida", Mon: "Dwowda", Tue: "Benada", Wed: "Wukuda", Thu: "Yawda", Fri: "Fida", Sat: "Memeneda"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "AN", PM: "EW"},
		},
	},
}
