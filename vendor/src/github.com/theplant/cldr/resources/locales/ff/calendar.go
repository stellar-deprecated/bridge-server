package ff

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM, y", Short: "d/M/y"},
		Time:     cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "sii", Feb: "col", Mar: "mbo", Apr: "see", May: "duu", Jun: "kor", Jul: "mor", Aug: "juk", Sep: "slt", Oct: "yar", Nov: "jol", Dec: "bow"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "s", Feb: "c", Mar: "m", Apr: "s", May: "d", Jun: "k", Jul: "m", Aug: "j", Sep: "s", Oct: "y", Nov: "j", Dec: "b"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "siilo", Feb: "colte", Mar: "mbooy", Apr: "seeɗto", May: "duujal", Jun: "korse", Jul: "morso", Aug: "juko", Sep: "siilto", Oct: "yarkomaa", Nov: "jolal", Dec: "bowte"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "dew", Mon: "aaɓ", Tue: "maw", Wed: "nje", Thu: "naa", Fri: "mwd", Sat: "hbi"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "d", Mon: "a", Tue: "m", Wed: "n", Thu: "n", Fri: "m", Sat: "h"},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "dewo", Mon: "aaɓnde", Tue: "mawbaare", Wed: "njeslaare", Thu: "naasaande", Fri: "mawnde", Sat: "hoore-biir"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "subaka", PM: "kikiiɗe"},
		},
	},
}
