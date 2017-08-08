package az

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "d MMMM y, EEEE", Long: "d MMMM y", Medium: "d MMM y", Short: "dd.MM.yy"},
		Time:     cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
		DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "yan", Feb: "fev", Mar: "mar", Apr: "apr", May: "may", Jun: "iyn", Jul: "iyl", Aug: "avq", Sep: "sen", Oct: "okt", Nov: "noy", Dec: "dek"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "1", Feb: "2", Mar: "3", Apr: "4", May: "5", Jun: "6", Jul: "7", Aug: "8", Sep: "9", Oct: "10", Nov: "11", Dec: "12"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Yanvar", Feb: "Fevral", Mar: "Mart", Apr: "Aprel", May: "May", Jun: "İyun", Jul: "İyul", Aug: "Avqust", Sep: "Sentyabr", Oct: "Oktyabr", Nov: "Noyabr", Dec: "Dekabr"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "B.", Mon: "B.E.", Tue: "Ç.A.", Wed: "Ç.", Thu: "C.A.", Fri: "C.", Sat: "Ş."},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "7", Mon: "1", Tue: "2", Wed: "3", Thu: "4", Fri: "5", Sat: "6"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "B.", Mon: "B.E.", Tue: "Ç.A.", Wed: "Ç.", Thu: "C.A.", Fri: "C.", Sat: "Ş."},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "bazar", Mon: "bazar ertəsi", Tue: "çərşənbə axşamı", Wed: "çərşənbə", Thu: "cümə axşamı", Fri: "cümə", Sat: "şənbə"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "a", PM: "p"},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"},
		},
	},
}
