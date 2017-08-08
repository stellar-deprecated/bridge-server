package ln

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "d/M/y"},
		Time:     cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "yan", Feb: "fbl", Mar: "msi", Apr: "apl", May: "mai", Jun: "yun", Jul: "yul", Aug: "agt", Sep: "stb", Oct: "ɔtb", Nov: "nvb", Dec: "dsb"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "y", Feb: "f", Mar: "m", Apr: "a", May: "m", Jun: "y", Jul: "y", Aug: "a", Sep: "s", Oct: "ɔ", Nov: "n", Dec: "d"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "sánzá ya yambo", Feb: "sánzá ya míbalé", Mar: "sánzá ya mísáto", Apr: "sánzá ya mínei", May: "sánzá ya mítáno", Jun: "sánzá ya motóbá", Jul: "sánzá ya nsambo", Aug: "sánzá ya mwambe", Sep: "sánzá ya libwa", Oct: "sánzá ya zómi", Nov: "sánzá ya zómi na mɔ̌kɔ́", Dec: "sánzá ya zómi na míbalé"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "eye", Mon: "ybo", Tue: "mbl", Wed: "mst", Thu: "min", Fri: "mtn", Sat: "mps"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "e", Mon: "y", Tue: "m", Wed: "m", Thu: "m", Fri: "m", Sat: "p"},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "eyenga", Mon: "mokɔlɔ mwa yambo", Tue: "mokɔlɔ mwa míbalé", Wed: "mokɔlɔ mwa mísáto", Thu: "mokɔlɔ ya mínéi", Fri: "mokɔlɔ ya mítáno", Sat: "mpɔ́sɔ"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "ntɔ́ngɔ́", PM: "mpókwa"},
		},
	},
}
