package af

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, dd MMMM y", Long: "dd MMMM y", Medium: "dd MMM y", Short: "y-MM-dd"},
		Time:     cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
		DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Jan.", Feb: "Feb.", Mar: "Mrt.", Apr: "Apr", May: "Mei", Jun: "Jun", Jul: "Jul", Aug: "Aug", Sep: "Sep", Oct: "Okt", Nov: "Nov", Dec: "Des"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Januarie", Feb: "Februarie", Mar: "Maart", Apr: "April", May: "Mei", Jun: "Junie", Jul: "Julie", Aug: "Augustus", Sep: "September", Oct: "Oktober", Nov: "November", Dec: "Desember"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "So", Mon: "Ma", Tue: "Di", Wed: "Wo", Thu: "Do", Fri: "Vr", Sat: "Sa"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "S", Mon: "M", Tue: "D", Wed: "W", Thu: "D", Fri: "V", Sat: "S"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "So.", Mon: "Ma.", Tue: "Di.", Wed: "Wo.", Thu: "Do.", Fri: "Vr.", Sat: "Sa."},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "Sondag", Mon: "Maandag", Tue: "Dinsdag", Wed: "Woensdag", Thu: "Donderdag", Fri: "Vrydag", Sat: "Saterdag"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "v", PM: "n"},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "vm.", PM: "nm."},
		},
	},
}
