package ca

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, d MMMM 'de' y", Long: "d MMMM 'de' y", Medium: "d MMM y", Short: "d/M/yy"},
		Time:     cldr.CalendarDateFormat{Full: "H:mm:ss zzzz", Long: "H:mm:ss z", Medium: "H:mm:ss", Short: "H:mm"},
		DateTime: cldr.CalendarDateFormat{Full: "{1} 'a les' {0}", Long: "{1}, {0}", Medium: "{1} , {0}", Short: "{1} {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "gen.", Feb: "febr.", Mar: "març", Apr: "abr.", May: "maig", Jun: "juny", Jul: "jul.", Aug: "ag.", Sep: "set.", Oct: "oct.", Nov: "nov.", Dec: "des."},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "GN", Feb: "FB", Mar: "MÇ", Apr: "AB", May: "MG", Jun: "JN", Jul: "JL", Aug: "AG", Sep: "ST", Oct: "OC", Nov: "NV", Dec: "DS"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "gener", Feb: "febrer", Mar: "març", Apr: "abril", May: "maig", Jun: "juny", Jul: "juliol", Aug: "agost", Sep: "setembre", Oct: "octubre", Nov: "novembre", Dec: "desembre"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "dg.", Mon: "dl.", Tue: "dt.", Wed: "dc.", Thu: "dj.", Fri: "dv.", Sat: "ds."},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "dg", Mon: "dl", Tue: "dt", Wed: "dc", Thu: "dj", Fri: "dv", Sat: "ds"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "dg.", Mon: "dl.", Tue: "dt.", Wed: "dc.", Thu: "dj.", Fri: "dv.", Sat: "ds."},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "diumenge", Mon: "dilluns", Tue: "dimarts", Wed: "dimecres", Thu: "dijous", Fri: "divendres", Sat: "dissabte"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "a.m.", PM: "p.m."},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "a. m.", PM: "p. m."},
		},
	},
}
