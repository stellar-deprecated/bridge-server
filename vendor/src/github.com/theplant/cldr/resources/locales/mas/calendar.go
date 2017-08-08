package mas

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"},
		Time:     cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Dal", Feb: "Ará", Mar: "Ɔɛn", Apr: "Doy", May: "Lép", Jun: "Rok", Jul: "Sás", Aug: "Bɔ́r", Sep: "Kús", Oct: "Gís", Nov: "Shʉ́", Dec: "Ntʉ́"},
			Narrow:      cldr.CalendarMonthFormatNameValue{},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Oladalʉ́", Feb: "Arát", Mar: "Ɔɛnɨ́ɔɨŋɔk", Apr: "Olodoyíóríê inkókúâ", May: "Oloilépūnyīē inkókúâ", Jun: "Kújúɔrɔk", Jul: "Mórusásin", Aug: "Ɔlɔ́ɨ́bɔ́rárɛ", Sep: "Kúshîn", Oct: "Olgísan", Nov: "Pʉshʉ́ka", Dec: "Ntʉ́ŋʉ́s"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Jpi", Mon: "Jtt", Tue: "Jnn", Wed: "Jtn", Thu: "Alh", Fri: "Iju", Sat: "Jmo"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "2", Mon: "3", Tue: "4", Wed: "5", Thu: "6", Fri: "7", Sat: "1"},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "Jumapílí", Mon: "Jumatátu", Tue: "Jumane", Wed: "Jumatánɔ", Thu: "Alaámisi", Fri: "Jumáa", Sat: "Jumamósi"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "Ɛnkakɛnyá", PM: "Ɛndámâ"},
		},
	},
}
