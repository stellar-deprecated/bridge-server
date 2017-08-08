package as

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, d MMMM, y", Long: "d MMMM, y", Medium: "dd-MM-y", Short: "d-M-y"},
		Time:     cldr.CalendarDateFormat{Full: "h.mm.ss a zzzz", Long: "h.mm.ss a z", Medium: "h.mm.ss a", Short: "h.mm. a"},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "জানু", Feb: "ফেব্ৰু", Mar: "মাৰ্চ", Apr: "এপ্ৰিল", May: "মে", Jun: "জুন", Jul: "জুলাই", Aug: "আগ", Sep: "সেপ্ট", Oct: "অক্টো", Nov: "নভে", Dec: "ডিসে"},
			Narrow:      cldr.CalendarMonthFormatNameValue{},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "জানুৱাৰী", Feb: "ফেব্ৰুৱাৰী", Mar: "মাৰ্চ", Apr: "এপ্ৰিল", May: "মে", Jun: "জুন", Jul: "জুলাই", Aug: "আগষ্ট", Sep: "ছেপ্তেম্বৰ", Oct: "অক্টোবৰ", Nov: "নৱেম্বৰ", Dec: "ডিচেম্বৰ"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "ৰবি", Mon: "সোম", Tue: "মঙ্গল", Wed: "বুধ", Thu: "বৃহষ্পতি", Fri: "শুক্ৰ", Sat: "শনি"},
			Narrow:      cldr.CalendarDayFormatNameValue{},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "দেওবাৰ", Mon: "সোমবাৰ", Tue: "মঙ্গলবাৰ", Wed: "বুধবাৰ", Thu: "বৃহষ্পতিবাৰ", Fri: "শুক্ৰবাৰ", Sat: "শনিবাৰ"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "পূৰ্বাহ্ণ", PM: "অপৰাহ্ণ"},
		},
	},
}
