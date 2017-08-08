package bas

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM, y", Short: "d/M/y"},
		Time:     cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "kɔn", Feb: "mac", Mar: "mat", Apr: "mto", May: "mpu", Jun: "hil", Jul: "nje", Aug: "hik", Sep: "dip", Oct: "bio", Nov: "may", Dec: "liɓ"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "k", Feb: "m", Mar: "m", Apr: "m", May: "m", Jun: "h", Jul: "n", Aug: "h", Sep: "d", Oct: "b", Nov: "m", Dec: "l"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Kɔndɔŋ", Feb: "Màcɛ̂l", Mar: "Màtùmb", Apr: "Màtop", May: "M̀puyɛ", Jun: "Hìlòndɛ̀", Jul: "Njèbà", Aug: "Hìkaŋ", Sep: "Dìpɔ̀s", Oct: "Bìòôm", Nov: "Màyɛsèp", Dec: "Lìbuy li ńyèe"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "nɔy", Mon: "nja", Tue: "uum", Wed: "ŋge", Thu: "mbɔ", Fri: "kɔɔ", Sat: "jon"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "n", Mon: "n", Tue: "u", Wed: "ŋ", Thu: "m", Fri: "k", Sat: "j"},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "ŋgwà nɔ̂y", Mon: "ŋgwà njaŋgumba", Tue: "ŋgwà ûm", Wed: "ŋgwà ŋgê", Thu: "ŋgwà mbɔk", Fri: "ŋgwà kɔɔ", Sat: "ŋgwà jôn"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "I bikɛ̂glà", PM: "I ɓugajɔp"},
		},
	},
}
