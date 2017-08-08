package fi

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "cccc d. MMMM y", Long: "d. MMMM y", Medium: "d.M.y", Short: "d.M.y"},
		Time:     cldr.CalendarDateFormat{Full: "H.mm.ss zzzz", Long: "H.mm.ss z", Medium: "H.mm.ss", Short: "H.mm"},
		DateTime: cldr.CalendarDateFormat{Full: "{1} 'klo' {0}", Long: "{1} 'klo' {0}", Medium: "{1} 'klo' {0}", Short: "{1} {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "tammi", Feb: "helmi", Mar: "maalis", Apr: "huhti", May: "touko", Jun: "kesä", Jul: "heinä", Aug: "elo", Sep: "syys", Oct: "loka", Nov: "marras", Dec: "joulu"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "T", Feb: "H", Mar: "M", Apr: "H", May: "T", Jun: "K", Jul: "H", Aug: "E", Sep: "S", Oct: "L", Nov: "M", Dec: "J"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "tammikuu", Feb: "helmikuu", Mar: "maaliskuu", Apr: "huhtikuu", May: "toukokuu", Jun: "kesäkuu", Jul: "heinäkuu", Aug: "elokuu", Sep: "syyskuu", Oct: "lokakuu", Nov: "marraskuu", Dec: "joulukuu"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "su", Mon: "ma", Tue: "ti", Wed: "ke", Thu: "to", Fri: "pe", Sat: "la"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "S", Mon: "M", Tue: "T", Wed: "K", Thu: "T", Fri: "P", Sat: "L"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "su", Mon: "ma", Tue: "ti", Wed: "ke", Thu: "to", Fri: "pe", Sat: "la"},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "sunnuntai", Mon: "maanantai", Tue: "tiistai", Wed: "keskiviikko", Thu: "torstai", Fri: "perjantai", Sat: "lauantai"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "ap.", PM: "ip."},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "ap.", PM: "ip."},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "ap.", PM: "ip."},
		},
	},
}
