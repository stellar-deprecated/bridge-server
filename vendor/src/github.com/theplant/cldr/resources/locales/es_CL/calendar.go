package es_CL

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "", Long: "", Medium: "dd-MM-y", Short: "dd-MM-yy"},
		Time:     cldr.CalendarDateFormat{},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{},
}
