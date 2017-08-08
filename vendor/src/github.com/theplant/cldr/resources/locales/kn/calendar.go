package kn

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, MMMM d, y", Long: "MMMM d, y", Medium: "MMM d, y", Short: "M/d/yy"},
		Time:     cldr.CalendarDateFormat{Full: "hh:mm:ss a zzzz", Long: "hh:mm:ss a z", Medium: "hh:mm:ss a", Short: "hh:mm a"},
		DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "ಜನ", Feb: "ಫೆಬ್ರ", Mar: "ಮಾರ್ಚ್", Apr: "ಏಪ್ರಿ", May: "ಮೇ", Jun: "ಜೂನ್", Jul: "ಜುಲೈ", Aug: "ಆಗ", Sep: "ಸೆಪ್ಟೆಂ", Oct: "ಅಕ್ಟೋ", Nov: "ನವೆಂ", Dec: "ಡಿಸೆಂ"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "ಜ", Feb: "ಫೆ", Mar: "ಮಾ", Apr: "ಏ", May: "ಮೇ", Jun: "ಜೂ", Jul: "ಜು", Aug: "ಆ", Sep: "ಸೆ", Oct: "ಅ", Nov: "ನ", Dec: "ಡಿ"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "ಜನವರಿ", Feb: "ಫೆಬ್ರವರಿ", Mar: "ಮಾರ್ಚ್", Apr: "ಏಪ್ರಿಲ್", May: "ಮೇ", Jun: "ಜೂನ್", Jul: "ಜುಲೈ", Aug: "ಆಗಸ್ಟ್", Sep: "ಸೆಪ್ಟೆಂಬರ್", Oct: "ಅಕ್ಟೋಬರ್", Nov: "ನವೆಂಬರ್", Dec: "ಡಿಸೆಂಬರ್"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "ಭಾನು", Mon: "ಸೋಮ", Tue: "ಮಂಗಳ", Wed: "ಬುಧ", Thu: "ಗುರು", Fri: "ಶುಕ್ರ", Sat: "ಶನಿ"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "ಭಾ", Mon: "ಸೋ", Tue: "ಮಂ", Wed: "ಬು", Thu: "ಗು", Fri: "ಶು", Sat: "ಶ"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "ಭಾನು", Mon: "ಸೋಮ", Tue: "ಮಂಗಳ", Wed: "ಬುಧ", Thu: "ಗುರು", Fri: "ಶುಕ್ರ", Sat: "ಶನಿ"},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "ಭಾನುವಾರ", Mon: "ಸೋಮವಾರ", Tue: "ಮಂಗಳವಾರ", Wed: "ಬುಧವಾರ", Thu: "ಗುರುವಾರ", Fri: "ಶುಕ್ರವಾರ", Sat: "ಶನಿವಾರ"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "ಪೂ", PM: "ಅ"},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "ಪೂರ್ವಾಹ್ನ", PM: "ಅಪರಾಹ್ನ"},
		},
	},
}
