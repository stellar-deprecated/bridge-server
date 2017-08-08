package lv

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, y. 'gada' d. MMMM", Long: "y. 'gada' d. MMMM", Medium: "y. 'gada' d. MMM", Short: "dd.MM.yy"},
		Time:     cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
		DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Janv.", Feb: "Febr.", Mar: "Marts", Apr: "Apr.", May: "Maijs", Jun: "Jūn.", Jul: "Jūl.", Aug: "Aug.", Sep: "Sept.", Oct: "Okt.", Nov: "Nov.", Dec: "Dec."},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Janvāris", Feb: "Februāris", Mar: "Marts", Apr: "Aprīlis", May: "Maijs", Jun: "Jūnijs", Jul: "Jūlijs", Aug: "Augusts", Sep: "Septembris", Oct: "Oktobris", Nov: "Novembris", Dec: "Decembris"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Sv", Mon: "Pr", Tue: "Ot", Wed: "Tr", Thu: "Ce", Fri: "Pk", Sat: "Se"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "S", Mon: "P", Tue: "O", Wed: "T", Thu: "C", Fri: "P", Sat: "S"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "Sv", Mon: "Pr", Tue: "Ot", Wed: "Tr", Thu: "Ce", Fri: "Pk", Sat: "Se"},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "Svētdiena", Mon: "Pirmdiena", Tue: "Otrdiena", Wed: "Trešdiena", Thu: "Ceturtdiena", Fri: "Piektdiena", Sat: "Sestdiena"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "priekšp.", PM: "pēcp."},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "priekšpusdienā", PM: "pēcpusdienā"},
		},
	},
}
