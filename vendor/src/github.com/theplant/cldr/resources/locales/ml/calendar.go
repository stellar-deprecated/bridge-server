package ml

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "y, MMMM d, EEEE", Long: "y, MMMM d", Medium: "y, MMM d", Short: "dd/MM/yy"},
		Time:     cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
		DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "ജനു", Feb: "ഫെബ്രു", Mar: "മാർ", Apr: "ഏപ്രി", May: "മേയ്", Jun: "ജൂൺ", Jul: "ജൂലൈ", Aug: "ഓഗ", Sep: "സെപ്റ്റം", Oct: "ഒക്ടോ", Nov: "നവം", Dec: "ഡിസം"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "ജ", Feb: "ഫ", Mar: "മാ", Apr: "ഏ", May: "മേ", Jun: "ജൂ", Jul: "ജൂ", Aug: "ഓ", Sep: "സ", Oct: "ഒ", Nov: "ന", Dec: "ഡി"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "ജനുവരി", Feb: "ഫെബ്രുവരി", Mar: "മാർച്ച്", Apr: "ഏപ്രിൽ", May: "മേയ്", Jun: "ജൂൺ", Jul: "ജൂലൈ", Aug: "ആഗസ്റ്റ്", Sep: "സെപ്റ്റംബർ", Oct: "ഒക്\u200cടോബർ", Nov: "നവംബർ", Dec: "ഡിസംബർ"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "ഞായർ", Mon: "തിങ്കൾ", Tue: "ചൊവ്വ", Wed: "ബുധൻ", Thu: "വ്യാഴം", Fri: "വെള്ളി", Sat: "ശനി"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "ഞ", Mon: "തി", Tue: "ച", Wed: "ബു", Thu: "വ", Fri: "വെ", Sat: "ശ"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "ഞാ", Mon: "തി", Tue: "ചൊ", Wed: "ബു", Thu: "വ്യാ", Fri: "വെ", Sat: "ശ"},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "ഞായറാഴ്\u200cച", Mon: "തിങ്കളാഴ്\u200cച", Tue: "ചൊവ്വാഴ്\u200cച", Wed: "ബുധനാഴ്\u200cച", Thu: "വ്യാഴാഴ്\u200cച", Fri: "വെള്ളിയാഴ്\u200cച", Sat: "ശനിയാഴ്\u200cച"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"},
		},
	},
}
