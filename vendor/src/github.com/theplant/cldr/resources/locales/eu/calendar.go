package eu

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "y('e')'ko' MMMM d, EEEE", Long: "y('e')'ko' MMMM d", Medium: "y MMM d", Short: "y/MM/dd"},
		Time:     cldr.CalendarDateFormat{Full: "HH:mm:ss (zzzz)", Long: "HH:mm:ss (z)", Medium: "HH:mm:ss", Short: "HH:mm"},
		DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Urt.", Feb: "Ots.", Mar: "Mar.", Apr: "Api.", May: "Mai.", Jun: "Eka.", Jul: "Uzt.", Aug: "Abu.", Sep: "Ira.", Oct: "Urr.", Nov: "Aza.", Dec: "Abe."},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "U", Feb: "O", Mar: "M", Apr: "A", May: "M", Jun: "E", Jul: "U", Aug: "A", Sep: "I", Oct: "U", Nov: "A", Dec: "A"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Urtarrila", Feb: "Otsaila", Mar: "Martxoa", Apr: "Apirila", May: "Maiatza", Jun: "Ekaina", Jul: "Uztaila", Aug: "Abuztua", Sep: "Iraila", Oct: "Urria", Nov: "Azaroa", Dec: "Abendua"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Ig.", Mon: "Al.", Tue: "Ar.", Wed: "Az.", Thu: "Og.", Fri: "Or.", Sat: "Lr."},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "I", Mon: "A", Tue: "A", Wed: "A", Thu: "O", Fri: "O", Sat: "L"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "ig.", Mon: "al.", Tue: "ar.", Wed: "az.", Thu: "og.", Fri: "or.", Sat: "lr."},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "Igandea", Mon: "Astelehena", Tue: "Asteartea", Wed: "Asteazkena", Thu: "Osteguna", Fri: "Ostirala", Sat: "Larunbata"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "g", PM: "a"},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"},
		},
	},
}
