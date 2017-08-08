package si

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "y MMMM d, EEEE", Long: "y MMMM d", Medium: "y MMM d", Short: "y-MM-dd"},
		Time:     cldr.CalendarDateFormat{Full: "a h.mm.ss zzzz", Long: "a h.mm.ss z", Medium: "a h.mm.ss", Short: "a h.mm"},
		DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "ජන", Feb: "පෙබ", Mar: "මාර්", Apr: "අප්\u200dරේල්", May: "මැයි", Jun: "ජූනි", Jul: "ජූලි", Aug: "අගෝ", Sep: "සැප්", Oct: "ඔක්", Nov: "නොවැ", Dec: "දෙසැ"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "ජ", Feb: "පෙ", Mar: "මා", Apr: "අ", May: "මැ", Jun: "ජූ", Jul: "ජූ", Aug: "අ", Sep: "සැ", Oct: "ඔ", Nov: "නෙ", Dec: "දෙ"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "ජනවාරි", Feb: "පෙබරවාරි", Mar: "මාර්තු", Apr: "අප්\u200dරේල්", May: "මැයි", Jun: "ජූනි", Jul: "ජූලි", Aug: "අගෝස්තු", Sep: "සැප්තැම්බර්", Oct: "ඔක්තෝබර්", Nov: "නොවැම්බර්", Dec: "දෙසැම්බර්"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "ඉරිදා", Mon: "සඳුදා", Tue: "අඟහ", Wed: "බදාදා", Thu: "බ්\u200dරහස්", Fri: "සිකු", Sat: "සෙන"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "ඉ", Mon: "ස", Tue: "අ", Wed: "බ", Thu: "බ්\u200dර", Fri: "සි", Sat: "සෙ"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "ඉරි", Mon: "සඳු", Tue: "අඟ", Wed: "බදා", Thu: "බ්\u200dරහ", Fri: "සිකු", Sat: "සෙන"},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "ඉරිදා", Mon: "සඳුදා", Tue: "අඟහරුවාදා", Wed: "බදාදා", Thu: "බ්\u200dරහස්පතින්දා", Fri: "සිකුරාදා", Sat: "සෙනසුරාදා"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "පෙ", PM: "ප"},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "පෙ.ව.", PM: "ප.ව."},
		},
	},
}
