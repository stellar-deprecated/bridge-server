package uz

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, y MMMM dd", Long: "y MMMM d", Medium: "y MMM d", Short: "yy/MM/dd"},
		Time:     cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
		DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Yanv", Feb: "Fev", Mar: "Mar", Apr: "Apr", May: "May", Jun: "Iyun", Jul: "Iyul", Aug: "Avg", Sep: "Sen", Oct: "Okt", Nov: "Noya", Dec: "Dek"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "Y", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "I", Jul: "I", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Yanvar", Feb: "Fevral", Mar: "Mart", Apr: "Aprel", May: "May", Jun: "Iyun", Jul: "Iyul", Aug: "Avgust", Sep: "Sentabr", Oct: "Oktabr", Nov: "Noyabr", Dec: "Dekabr"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Yaksh", Mon: "Dush", Tue: "Sesh", Wed: "Chor", Thu: "Pay", Fri: "Jum", Sat: "Shan"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "Y", Mon: "D", Tue: "S", Wed: "C", Thu: "P", Fri: "J", Sat: "S"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "Ya", Mon: "Du", Tue: "Se", Wed: "Cho", Thu: "Pa", Fri: "Ju", Sat: "Sha"},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "yakshanba", Mon: "dushanba", Tue: "seshanba", Wed: "chorshanba", Thu: "payshanba", Fri: "juma", Sat: "shanba"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "o", PM: "k"},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "TO", PM: "TK"},
		},
	},
}
