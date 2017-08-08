package nl

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd-MM-yy"},
		Time:     cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
		DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Jan.", Feb: "Feb.", Mar: "Mrt.", Apr: "Apr.", May: "Mei", Jun: "Jun.", Jul: "Jul.", Aug: "Aug.", Sep: "Sep.", Oct: "Okt.", Nov: "Nov.", Dec: "Dec."},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Januari", Feb: "Februari", Mar: "Maart", Apr: "April", May: "Mei", Jun: "Juni", Jul: "Juli", Aug: "Augustus", Sep: "September", Oct: "Oktober", Nov: "November", Dec: "December"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Zo", Mon: "Ma", Tue: "Di", Wed: "Wo", Thu: "Do", Fri: "Vr", Sat: "Za"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "Z", Mon: "M", Tue: "D", Wed: "W", Thu: "D", Fri: "V", Sat: "Z"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "Zo", Mon: "Ma", Tue: "Di", Wed: "Wo", Thu: "Do", Fri: "Vr", Sat: "Za"},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "Zondag", Mon: "Maandag", Tue: "Dinsdag", Wed: "Woensdag", Thu: "Donderdag", Fri: "Vrijdag", Sat: "Zaterdag"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "a.m.", PM: "p.m."},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "a.m.", PM: "p.m."},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "voormiddag", PM: "p.m."},
		},
	},
}
