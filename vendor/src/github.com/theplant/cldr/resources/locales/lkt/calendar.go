package lkt

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, MMMM d, y", Long: "MMMM d, y", Medium: "MMM d, y", Short: "M/d/yy"},
		Time:     cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{},
			Narrow:      cldr.CalendarMonthFormatNameValue{},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Wiótheȟika Wí", Feb: "Thiyóȟeyuŋka Wí", Mar: "Ištáwičhayazaŋ Wí", Apr: "Pȟežítȟo Wí", May: "Čhaŋwápetȟo Wí", Jun: "Wípazukȟa-wašté Wí", Jul: "Čhaŋpȟásapa Wí", Aug: "Wasútȟuŋ Wí", Sep: "Čhaŋwápeǧi Wí", Oct: "Čhaŋwápe-kasná Wí", Nov: "Waníyetu Wí", Dec: "Tȟahékapšuŋ Wí"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "A", Mon: "W", Tue: "N", Wed: "Y", Thu: "T", Fri: "Z", Sat: "O"},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "Aŋpétuwakȟaŋ", Mon: "Aŋpétuwaŋži", Tue: "Aŋpétunuŋpa", Wed: "Aŋpétuyamni", Thu: "Aŋpétutopa", Fri: "Aŋpétuzaptaŋ", Sat: "Owáŋgyužažapi"},
		},
		Periods: cldr.CalendarPeriodFormatNames{},
	},
}
