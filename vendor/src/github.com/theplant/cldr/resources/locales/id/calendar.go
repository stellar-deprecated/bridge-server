package id

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, dd MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/yy"},
		Time:     cldr.CalendarDateFormat{Full: "HH.mm.ss zzzz", Long: "HH.mm.ss z", Medium: "HH.mm.ss", Short: "HH.mm"},
		DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Jan", Feb: "Feb", Mar: "Mar", Apr: "Apr", May: "Mei", Jun: "Jun", Jul: "Jul", Aug: "Agt", Sep: "Sep", Oct: "Okt", Nov: "Nov", Dec: "Des"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Januari", Feb: "Februari", Mar: "Maret", Apr: "April", May: "Mei", Jun: "Juni", Jul: "Juli", Aug: "Agustus", Sep: "September", Oct: "Oktober", Nov: "November", Dec: "Desember"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Min", Mon: "Sen", Tue: "Sel", Wed: "Rab", Thu: "Kam", Fri: "Jum", Sat: "Sab"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "M", Mon: "S", Tue: "S", Wed: "R", Thu: "K", Fri: "J", Sat: "S"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "Min", Mon: "Sen", Tue: "Sel", Wed: "Rab", Thu: "Kam", Fri: "Jum", Sat: "Sab"},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "Minggu", Mon: "Senin", Tue: "Selasa", Wed: "Rabu", Thu: "Kamis", Fri: "Jumat", Sat: "Sabtu"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"},
		},
	},
}
