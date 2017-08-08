package agq

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM, y", Short: "d/M/y"},
		Time:     cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "nùm", Feb: "kɨz", Mar: "tɨd", Apr: "taa", May: "see", Jun: "nzu", Jul: "dum", Aug: "fɔe", Sep: "dzu", Oct: "lɔm", Nov: "kaa", Dec: "fwo"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "n", Feb: "k", Mar: "t", Apr: "t", May: "s", Jun: "z", Jul: "k", Aug: "f", Sep: "d", Oct: "l", Nov: "c", Dec: "f"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "ndzɔ̀ŋɔ̀nùm", Feb: "ndzɔ̀ŋɔ̀kƗ̀zùʔ", Mar: "ndzɔ̀ŋɔ̀tƗ̀dʉ̀ghà", Apr: "ndzɔ̀ŋɔ̀tǎafʉ̄ghā", May: "ndzɔ̀ŋèsèe", Jun: "ndzɔ̀ŋɔ̀nzùghò", Jul: "ndzɔ̀ŋɔ̀dùmlo", Aug: "ndzɔ̀ŋɔ̀kwîfɔ̀e", Sep: "ndzɔ̀ŋɔ̀tƗ̀fʉ̀ghàdzughù", Oct: "ndzɔ̀ŋɔ̀ghǔuwelɔ̀m", Nov: "ndzɔ̀ŋɔ̀chwaʔàkaa wo", Dec: "ndzɔ̀ŋèfwòo"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "nts", Mon: "kpa", Tue: "ghɔ", Wed: "tɔm", Thu: "ume", Fri: "ghɨ", Sat: "dzk"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "n", Mon: "k", Tue: "g", Wed: "t", Thu: "u", Fri: "g", Sat: "d"},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "tsuʔntsɨ", Mon: "tsuʔukpà", Tue: "tsuʔughɔe", Wed: "tsuʔutɔ̀mlò", Thu: "tsuʔumè", Fri: "tsuʔughɨ̂m", Sat: "tsuʔndzɨkɔʔɔ"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "a.g", PM: "a.k"},
		},
	},
}
