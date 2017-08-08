package nyn

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"},
		Time:     cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "KBZ", Feb: "KBR", Mar: "KST", Apr: "KKN", May: "KTN", Jun: "KMK", Jul: "KMS", Aug: "KMN", Sep: "KMW", Oct: "KKM", Nov: "KNK", Dec: "KNB"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Okwokubanza", Feb: "Okwakabiri", Mar: "Okwakashatu", Apr: "Okwakana", May: "Okwakataana", Jun: "Okwamukaaga", Jul: "Okwamushanju", Aug: "Okwamunaana", Sep: "Okwamwenda", Oct: "Okwaikumi", Nov: "Okwaikumi na kumwe", Dec: "Okwaikumi na ibiri"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "SAN", Mon: "ORK", Tue: "OKB", Wed: "OKS", Thu: "OKN", Fri: "OKT", Sat: "OMK"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "S", Mon: "K", Tue: "R", Wed: "S", Thu: "N", Fri: "T", Sat: "M"},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "Sande", Mon: "Orwokubanza", Tue: "Orwakabiri", Wed: "Orwakashatu", Thu: "Orwakana", Fri: "Orwakataano", Sat: "Orwamukaaga"},
		},
		Periods: cldr.CalendarPeriodFormatNames{},
	},
}
