package fo

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE dd MMMM y", Long: "d. MMM y", Medium: "dd-MM-y", Short: "dd-MM-yy"},
		Time:     cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
		DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "jan", Feb: "feb", Mar: "mar", Apr: "apr", May: "mai", Jun: "jun", Jul: "jul", Aug: "aug", Sep: "sep", Oct: "okt", Nov: "nov", Dec: "des"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "januar", Feb: "februar", Mar: "mars", Apr: "apríl", May: "mai", Jun: "juni", Jul: "juli", Aug: "august", Sep: "september", Oct: "oktober", Nov: "november", Dec: "desember"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "sun", Mon: "mán", Tue: "týs", Wed: "mik", Thu: "hós", Fri: "frí", Sat: "ley"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "S", Mon: "M", Tue: "T", Wed: "M", Thu: "H", Fri: "F", Sat: "L"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "su", Mon: "má", Tue: "tý", Wed: "mi", Thu: "hó", Fri: "fr", Sat: "le"},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "sunnudagur", Mon: "mánadagur", Tue: "týsdagur", Wed: "mikudagur", Thu: "hósdagur", Fri: "fríggjadagur", Sat: "leygardagur"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "f.p.", PM: "s.p."},
			Narrow:      cldr.CalendarPeriodFormatNameValue{},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "fyrrapartur", PM: "seinnapartur"},
		},
	},
}
