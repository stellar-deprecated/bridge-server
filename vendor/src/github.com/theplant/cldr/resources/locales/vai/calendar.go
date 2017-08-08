package vai

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"},
		Time:     cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{},
			Narrow:      cldr.CalendarMonthFormatNameValue{},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "ꖨꕪꖃ ꔞꕮ", Feb: "ꕒꕡꖝꖕ", Mar: "ꕾꖺ", Apr: "ꖢꖕ", May: "ꖑꕱ", Jun: "6", Jul: "7", Aug: "ꗛꔕ", Sep: "ꕢꕌ", Oct: "ꕭꖃ", Nov: "ꔞꘋꕔꕿ ꕸꖃꗏ", Dec: "ꖨꕪꕱ ꗏꕮ"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{},
			Narrow:      cldr.CalendarDayFormatNameValue{},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "ꕞꕌꔵ", Mon: "ꗳꗡꘉ", Tue: "ꕚꕞꕚ", Wed: "ꕉꕞꕒ", Thu: "ꕉꔤꕆꕢ", Fri: "ꕉꔤꕀꕮ", Sat: "ꔻꔬꔳ"},
		},
		Periods: cldr.CalendarPeriodFormatNames{},
	},
}
