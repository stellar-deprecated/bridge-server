package lt

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "y 'm'. MMMM d 'd'., EEEE", Long: "y 'm'. MMMM d 'd'.", Medium: "y-MM-dd", Short: "y-MM-dd"},
		Time:     cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
		DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "saus.", Feb: "vas.", Mar: "kov.", Apr: "bal.", May: "geg.", Jun: "birž.", Jul: "liep.", Aug: "rugp.", Sep: "rugs.", Oct: "spal.", Nov: "lapkr.", Dec: "gruod."},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "S", Feb: "V", Mar: "K", Apr: "B", May: "G", Jun: "B", Jul: "L", Aug: "R", Sep: "R", Oct: "S", Nov: "L", Dec: "G"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "sausis", Feb: "vasaris", Mar: "kovas", Apr: "balandis", May: "gegužė", Jun: "birželis", Jul: "liepa", Aug: "rugpjūtis", Sep: "rugsėjis", Oct: "spalis", Nov: "lapkritis", Dec: "gruodis"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "sk", Mon: "pr", Tue: "an", Wed: "tr", Thu: "kt", Fri: "pn", Sat: "št"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "S", Mon: "P", Tue: "A", Wed: "T", Thu: "K", Fri: "P", Sat: "Š"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "Sk", Mon: "Pr", Tue: "An", Wed: "Tr", Thu: "Kt", Fri: "Pn", Sat: "Št"},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "sekmadienis", Mon: "pirmadienis", Tue: "antradienis", Wed: "trečiadienis", Thu: "ketvirtadienis", Fri: "penktadienis", Sat: "šeštadienis"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "pr.p.", PM: "pop."},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "pr.p.", PM: "pop."},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "priešpiet", PM: "popiet"},
		},
	},
}
