package gu

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, d MMMM, y", Long: "d MMMM, y", Medium: "d MMM, y", Short: "d/M/yy"},
		Time:     cldr.CalendarDateFormat{Full: "hh:mm:ss a zzzz", Long: "hh:mm:ss a z", Medium: "hh:mm:ss a", Short: "hh:mm a"},
		DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "જાન્યુ", Feb: "ફેબ્રુ", Mar: "માર્ચ", Apr: "એપ્રિલ", May: "મે", Jun: "જૂન", Jul: "જુલાઈ", Aug: "ઑગસ્ટ", Sep: "સપ્ટે", Oct: "ઑક્ટો", Nov: "નવે", Dec: "ડિસે"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "જા", Feb: "ફે", Mar: "મા", Apr: "એ", May: "મે", Jun: "જૂ", Jul: "જુ", Aug: "ઑ", Sep: "સ", Oct: "ઑ", Nov: "ન", Dec: "ડિ"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "જાન્યુઆરી", Feb: "ફેબ્રુઆરી", Mar: "માર્ચ", Apr: "એપ્રિલ", May: "મે", Jun: "જૂન", Jul: "જુલાઈ", Aug: "ઑગસ્ટ", Sep: "સપ્ટેમ્બર", Oct: "ઑક્ટોબર", Nov: "નવેમ્બર", Dec: "ડિસેમ્બર"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "રવિ", Mon: "સોમ", Tue: "મંગળ", Wed: "બુધ", Thu: "ગુરુ", Fri: "શુક્ર", Sat: "શનિ"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "ર", Mon: "સો", Tue: "મં", Wed: "બુ", Thu: "ગુ", Fri: "શુ", Sat: "શ"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "ર", Mon: "સો", Tue: "મં", Wed: "બુ", Thu: "ગુ", Fri: "શુ", Sat: "શ"},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "રવિવાર", Mon: "સોમવાર", Tue: "મંગળવાર", Wed: "બુધવાર", Thu: "ગુરુવાર", Fri: "શુક્રવાર", Sat: "શનિવાર"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"},
		},
	},
}
