package chr

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, MMMM d, y", Long: "MMMM d, y", Medium: "MMM d, y", Short: "M/d/yy"},
		Time:     cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "ᎤᏃ", Feb: "ᎧᎦ", Mar: "ᎠᏅ", Apr: "ᎧᏬ", May: "ᎠᏂ", Jun: "ᏕᎭ", Jul: "ᎫᏰ", Aug: "ᎦᎶ", Sep: "ᏚᎵ", Oct: "ᏚᏂ", Nov: "ᏅᏓ", Dec: "ᎥᏍ"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "Ꭴ", Feb: "Ꭷ", Mar: "Ꭰ", Apr: "Ꭷ", May: "Ꭰ", Jun: "Ꮥ", Jul: "Ꭻ", Aug: "Ꭶ", Sep: "Ꮪ", Oct: "Ꮪ", Nov: "Ꮕ", Dec: "Ꭵ"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "ᎤᏃᎸᏔᏅ", Feb: "ᎧᎦᎵ", Mar: "ᎠᏅᏱ", Apr: "ᎧᏬᏂ", May: "ᎠᏂᏍᎬᏘ", Jun: "ᏕᎭᎷᏱ", Jul: "ᎫᏰᏉᏂ", Aug: "ᎦᎶᏂ", Sep: "ᏚᎵᏍᏗ", Oct: "ᏚᏂᏅᏗ", Nov: "ᏅᏓᏕᏆ", Dec: "ᎥᏍᎩᏱ"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "ᏆᏍᎬ", Mon: "ᏉᏅᎯ", Tue: "ᏔᎵᏁ", Wed: "ᏦᎢᏁ", Thu: "ᏅᎩᏁ", Fri: "ᏧᎾᎩ", Sat: "ᏈᏕᎾ"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "Ꮖ", Mon: "Ꮙ", Tue: "Ꮤ", Wed: "Ꮶ", Thu: "Ꮕ", Fri: "Ꮷ", Sat: "Ꭴ"},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "ᎤᎾᏙᏓᏆᏍᎬ", Mon: "ᎤᎾᏙᏓᏉᏅᎯ", Tue: "ᏔᎵᏁᎢᎦ", Wed: "ᏦᎢᏁᎢᎦ", Thu: "ᏅᎩᏁᎢᎦ", Fri: "ᏧᎾᎩᎶᏍᏗ", Sat: "ᎤᎾᏙᏓᏈᏕᎾ"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "ᏌᎾᎴ", PM: "ᏒᎯᏱᎢᏗᏢ"},
		},
	},
}
