package ha

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, d MMMM, y", Long: "d MMMM, y", Medium: "d MMM, y", Short: "d/M/yy"},
		Time:     cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Jan", Feb: "Fab", Mar: "Mar", Apr: "Afi", May: "May", Jun: "Yun", Jul: "Yul", Aug: "Agu", Sep: "Sat", Oct: "Okt", Nov: "Nuw", Dec: "Dis"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "Y", Jul: "Y", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Janairu", Feb: "Faburairu", Mar: "Maris", Apr: "Afirilu", May: "Mayu", Jun: "Yuni", Jul: "Yuli", Aug: "Agusta", Sep: "Satumba", Oct: "Oktoba", Nov: "Nuwamba", Dec: "Disamba"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Lh", Mon: "Li", Tue: "Ta", Wed: "Lr", Thu: "Al", Fri: "Ju", Sat: "As"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "L", Mon: "L", Tue: "T", Wed: "L", Thu: "A", Fri: "J", Sat: "A"},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "Lahadi", Mon: "Litinin", Tue: "Talata", Wed: "Laraba", Thu: "Alhamis", Fri: "Jummaʼa", Sat: "Asabar"},
		},
		Periods: cldr.CalendarPeriodFormatNames{},
	},
}
