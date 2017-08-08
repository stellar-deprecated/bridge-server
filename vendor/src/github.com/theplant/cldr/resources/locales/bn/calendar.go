package bn

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, d MMMM, y", Long: "d MMMM, y", Medium: "d MMM, y", Short: "d/M/yy"},
		Time:     cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
		DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "জানুয়ারী", Feb: "ফেব্রুয়ারী", Mar: "মার্চ", Apr: "এপ্রিল", May: "মে", Jun: "জুন", Jul: "জুলাই", Aug: "আগস্ট", Sep: "সেপ্টেম্বর", Oct: "অক্টোবর", Nov: "নভেম্বর", Dec: "ডিসেম্বর"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "জা", Feb: "ফে", Mar: "মা", Apr: "এ", May: "মে", Jun: "জুন", Jul: "জু", Aug: "আ", Sep: "সে", Oct: "অ", Nov: "ন", Dec: "ডি"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "জানুয়ারী", Feb: "ফেব্রুয়ারী", Mar: "মার্চ", Apr: "এপ্রিল", May: "মে", Jun: "জুন", Jul: "জুলাই", Aug: "আগস্ট", Sep: "সেপ্টেম্বর", Oct: "অক্টোবর", Nov: "নভেম্বর", Dec: "ডিসেম্বর"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "রবি", Mon: "সোম", Tue: "মঙ্গল", Wed: "বুধ", Thu: "বৃহস্পতি", Fri: "শুক্র", Sat: "শনি"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "র", Mon: "সো", Tue: "ম", Wed: "বু", Thu: "বৃ", Fri: "শু", Sat: "শ"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "রঃ", Mon: "সোঃ", Tue: "মঃ", Wed: "বুঃ", Thu: "বৃঃ", Fri: "শুঃ", Sat: "শনি"},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "রবিবার", Mon: "সোমবার", Tue: "মঙ্গলবার", Wed: "বুধবার", Thu: "বৃহষ্পতিবার", Fri: "শুক্রবার", Sat: "শনিবার"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "পূর্বাহ্ণ", PM: "অপরাহ্ণ"},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "পূর্বাহ্ণ", PM: "অপরাহ্ণ"},
		},
	},
}
