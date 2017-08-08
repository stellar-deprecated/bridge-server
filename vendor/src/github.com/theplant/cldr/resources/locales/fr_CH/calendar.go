package fr_CH

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "", Medium: "", Short: "dd.MM.yy"},
		Time:     cldr.CalendarDateFormat{Full: "HH.mm:ss 'h' zzzz", Long: "", Medium: "", Short: ""},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{},
}
