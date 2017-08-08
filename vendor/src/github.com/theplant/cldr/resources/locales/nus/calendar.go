package nus

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "d/MM/y"},
		Time:     cldr.CalendarDateFormat{Full: "zzzz h:mm:ss a", Long: "z h:mm:ss a", Medium: "h:mm:ss a", Short: "h:mm a"},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Tiop", Feb: "Pɛt", Mar: "Duɔ̱ɔ̱", Apr: "Guak", May: "Duä", Jun: "Kor", Jul: "Pay", Aug: "Thoo", Sep: "Tɛɛ", Oct: "Laa", Nov: "Kur", Dec: "Tid"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "T", Feb: "P", Mar: "D", Apr: "G", May: "D", Jun: "K", Jul: "P", Aug: "T", Sep: "T", Oct: "L", Nov: "K", Dec: "T"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Tiop thar pɛt", Feb: "Pɛt", Mar: "Duɔ̱ɔ̱ŋ", Apr: "Guak", May: "Duät", Jun: "Kornyoot", Jul: "Pay yie̱tni", Aug: "Tho̱o̱r", Sep: "Tɛɛr", Oct: "Laath", Nov: "Kur", Dec: "Tio̱p in di̱i̱t"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Cäŋ", Mon: "Jiec", Tue: "Rɛw", Wed: "Diɔ̱k", Thu: "Ŋuaan", Fri: "Dhieec", Sat: "Bäkɛl"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "C", Mon: "J", Tue: "R", Wed: "D", Thu: "Ŋ", Fri: "D", Sat: "B"},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "Cäŋ kuɔth", Mon: "Jiec la̱t", Tue: "Rɛw lätni", Wed: "Diɔ̱k lätni", Thu: "Ŋuaan lätni", Fri: "Dhieec lätni", Sat: "Bäkɛl lätni"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "RW", PM: "TŊ"},
		},
	},
}
