package tr

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "d MMMM y EEEE", Long: "d MMMM y", Medium: "d MMM y", Short: "d.MM.y"},
		Time:     cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
		DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Oca", Feb: "Şub", Mar: "Mar", Apr: "Nis", May: "May", Jun: "Haz", Jul: "Tem", Aug: "Ağu", Sep: "Eyl", Oct: "Eki", Nov: "Kas", Dec: "Ara"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "O", Feb: "Ş", Mar: "M", Apr: "N", May: "M", Jun: "H", Jul: "T", Aug: "A", Sep: "E", Oct: "E", Nov: "K", Dec: "A"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Ocak", Feb: "Şubat", Mar: "Mart", Apr: "Nisan", May: "Mayıs", Jun: "Haziran", Jul: "Temmuz", Aug: "Ağustos", Sep: "Eylül", Oct: "Ekim", Nov: "Kasım", Dec: "Aralık"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Paz", Mon: "Pzt", Tue: "Sal", Wed: "Çar", Thu: "Per", Fri: "Cum", Sat: "Cmt"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "P", Mon: "P", Tue: "S", Wed: "Ç", Thu: "P", Fri: "C", Sat: "C"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "Pa", Mon: "Pt", Tue: "Sa", Wed: "Ça", Thu: "Pe", Fri: "Cu", Sat: "Ct"},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "Pazar", Mon: "Pazartesi", Tue: "Salı", Wed: "Çarşamba", Thu: "Perşembe", Fri: "Cuma", Sat: "Cumartesi"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "öö", PM: "ös"},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "ÖÖ", PM: "ÖS"},
		},
	},
}
