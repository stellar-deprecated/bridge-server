package nmg

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "d/M/y"},
		Time:     cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "ng1", Feb: "ng2", Mar: "ng3", Apr: "ng4", May: "ng5", Jun: "ng6", Jul: "ng7", Aug: "ng8", Sep: "ng9", Oct: "ng10", Nov: "ng11", Dec: "kris"},
			Narrow:      cldr.CalendarMonthFormatNameValue{},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "ngwɛn matáhra", Feb: "ngwɛn ńmba", Mar: "ngwɛn ńlal", Apr: "ngwɛn ńna", May: "ngwɛn ńtan", Jun: "ngwɛn ńtuó", Jul: "ngwɛn hɛmbuɛrí", Aug: "ngwɛn lɔmbi", Sep: "ngwɛn rɛbvuâ", Oct: "ngwɛn wum", Nov: "ngwɛn wum navǔr", Dec: "krísimin"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "sɔ́n", Mon: "mɔ́n", Tue: "smb", Wed: "sml", Thu: "smn", Fri: "mbs", Sat: "sas"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "s", Mon: "m", Tue: "s", Wed: "s", Thu: "s", Fri: "m", Sat: "s"},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "sɔ́ndɔ", Mon: "mɔ́ndɔ", Tue: "sɔ́ndɔ mafú mába", Wed: "sɔ́ndɔ mafú málal", Thu: "sɔ́ndɔ mafú mána", Fri: "mabágá má sukul", Sat: "sásadi"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "maná", PM: "kugú"},
		},
	},
}
