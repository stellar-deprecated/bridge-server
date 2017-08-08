package en_150

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "", Long: "", Medium: "dd MMM y", Short: "dd/MM/yy"},
		Time:     cldr.CalendarDateFormat{Full: "HH 'h' mm 'min' ss 's' zzzz", Long: "", Medium: "", Short: ""},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{},
}
