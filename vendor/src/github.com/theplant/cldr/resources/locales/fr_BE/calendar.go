package fr_BE

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "", Long: "", Medium: "", Short: "d/MM/yy"},
		Time:     cldr.CalendarDateFormat{Full: "H 'h' mm 'min' ss 's' zzzz", Long: "", Medium: "", Short: ""},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{},
}
