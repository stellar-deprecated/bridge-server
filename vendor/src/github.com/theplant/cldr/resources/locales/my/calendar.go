package my

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE၊ dd MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd-MM-yy"},
		Time:     cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
		DateTime: cldr.CalendarDateFormat{Full: "{1}မှာ {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "ဇန်", Feb: "ဖေ", Mar: "မတ်", Apr: "ဧပြီ", May: "မေ", Jun: "ဇွန်", Jul: "ဇူ", Aug: "ဩ", Sep: "စက်", Oct: "အောက်", Nov: "နို", Dec: "ဒီ"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "ဇ", Feb: "ဖ", Mar: "မ", Apr: "ဧ", May: "မ", Jun: "ဇ", Jul: "ဇ", Aug: "ဩ", Sep: "စ", Oct: "အ", Nov: "န", Dec: "ဒ"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "ဇန်နဝါရီ", Feb: "ဖေဖော်ဝါရီ", Mar: "မတ်", Apr: "ဧပြီ", May: "မေ", Jun: "ဇွန်", Jul: "ဇူလိုင်", Aug: "ဩဂုတ်", Sep: "စက်တင်ဘာ", Oct: "အောက်တိုဘာ", Nov: "နိုဝင်ဘာ", Dec: "ဒီဇင်ဘာ"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "တနင်္ဂနွေ", Mon: "တနင်္လာ", Tue: "အင်္ဂါ", Wed: "ဗုဒ္ဓဟူး", Thu: "ကြာသပတေး", Fri: "သောကြာ", Sat: "စနေ"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "တ", Mon: "တ", Tue: "အ", Wed: "ဗ", Thu: "က", Fri: "သ", Sat: "စ"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "တနင်္ဂနွေ", Mon: "တနင်္လာ", Tue: "အင်္ဂါ", Wed: "ဗုဒ္ဓဟူး", Thu: "ကြာသပတေး", Fri: "သောကြာ", Sat: "စနေ"},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "တနင်္ဂနွေ", Mon: "တနင်္လာ", Tue: "အင်္ဂါ", Wed: "ဗုဒ္ဓဟူး", Thu: "ကြာသပတေး", Fri: "သောကြာ", Sat: "စနေ"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "နံနက်", PM: "ညနေ"},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "နံနက်", PM: "ညနေ"},
		},
	},
}
