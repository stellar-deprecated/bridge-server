package smn

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{},
			Narrow:      cldr.CalendarMonthFormatNameValue{},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "uđđâivemáánu", Feb: "kuovâmáánu", Mar: "njuhčâmáánu", Apr: "cuáŋuimáánu", May: "vyesimáánu", Jun: "kesimáánu", Jul: "syeinimáánu", Aug: "porgemáánu", Sep: "čohčâmáánu", Oct: "roovvâdmáánu", Nov: "skammâmáánu", Dec: "juovlâmáánu"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "pa", Mon: "vu", Tue: "ma", Wed: "ko", Thu: "tu", Fri: "vá", Sat: "lá"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "P", Mon: "V", Tue: "M", Wed: "K", Thu: "T", Fri: "V", Sat: "L"},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "pasepeivi", Mon: "vuossargâ", Tue: "majebargâ", Wed: "koskokko", Thu: "tuorâstâh", Fri: "vástuppeivi", Sat: "lávurdâh"},
		},
		Periods: cldr.CalendarPeriodFormatNames{},
	},
}
