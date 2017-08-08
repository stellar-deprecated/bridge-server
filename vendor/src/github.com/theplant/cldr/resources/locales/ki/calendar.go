package ki

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"},
		Time:     cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "JEN", Feb: "WKR", Mar: "WGT", Apr: "WKN", May: "WTN", Jun: "WTD", Jul: "WMJ", Aug: "WNN", Sep: "WKD", Oct: "WIK", Nov: "WMW", Dec: "DIT"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "K", Mar: "G", Apr: "K", May: "G", Jun: "G", Jul: "M", Aug: "K", Sep: "K", Oct: "I", Nov: "I", Dec: "D"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Njenuarĩ", Feb: "Mwere wa kerĩ", Mar: "Mwere wa gatatũ", Apr: "Mwere wa kana", May: "Mwere wa gatano", Jun: "Mwere wa gatandatũ", Jul: "Mwere wa mũgwanja", Aug: "Mwere wa kanana", Sep: "Mwere wa kenda", Oct: "Mwere wa ikũmi", Nov: "Mwere wa ikũmi na ũmwe", Dec: "Ndithemba"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "KMA", Mon: "NTT", Tue: "NMN", Wed: "NMT", Thu: "ART", Fri: "NMA", Sat: "NMM"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "K", Mon: "N", Tue: "N", Wed: "N", Thu: "A", Fri: "N", Sat: "N"},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "Kiumia", Mon: "Njumatatũ", Tue: "Njumaine", Wed: "Njumatana", Thu: "Aramithi", Fri: "Njumaa", Sat: "Njumamothi"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "Kiroko", PM: "Hwaĩ-inĩ"},
		},
	},
}
