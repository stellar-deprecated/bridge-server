package ta

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, d MMMM, y", Long: "d MMMM, y", Medium: "d MMM, y", Short: "d-M-yy"},
		Time:     cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
		DateTime: cldr.CalendarDateFormat{Full: "{1} ’அன்று’ {0}", Long: "{1} ’அன்று’ {0}", Medium: "{1}, {0}", Short: "{1}, {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "ஜன.", Feb: "பிப்.", Mar: "மார்.", Apr: "ஏப்.", May: "மே", Jun: "ஜூன்", Jul: "ஜூலை", Aug: "ஆக.", Sep: "செப்.", Oct: "அக்.", Nov: "நவ.", Dec: "டிச."},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "ஜ", Feb: "பி", Mar: "மா", Apr: "ஏ", May: "மே", Jun: "ஜூ", Jul: "ஜூ", Aug: "ஆ", Sep: "செ", Oct: "அ", Nov: "ந", Dec: "டி"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "ஜனவரி", Feb: "பிப்ரவரி", Mar: "மார்ச்", Apr: "ஏப்ரல்", May: "மே", Jun: "ஜூன்", Jul: "ஜூலை", Aug: "ஆகஸ்டு", Sep: "செப்டம்பர்", Oct: "அக்டோபர்", Nov: "நவம்பர்", Dec: "டிசம்பர்"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "ஞா", Mon: "தி", Tue: "செ", Wed: "பு", Thu: "வி", Fri: "வெ", Sat: "ச"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "ஞா", Mon: "தி", Tue: "செ", Wed: "பு", Thu: "வி", Fri: "வெ", Sat: "ச"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "ஞா", Mon: "தி", Tue: "செ", Wed: "பு", Thu: "வி", Fri: "வெ", Sat: "ச"},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "ஞாயிறு", Mon: "திங்கள்", Tue: "செவ்வாய்", Wed: "புதன்", Thu: "வியாழன்", Fri: "வெள்ளி", Sat: "சனி"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "முற்பகல்", PM: "பிற்பகல்"},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "மு.ப", PM: "பி.ப"},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "முற்பகல்", PM: "பிற்பகல்"},
		},
	},
}
