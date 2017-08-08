package nnh

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE , 'lyɛ'̌ʼ d 'na' MMMM, y", Long: "'lyɛ'̌ʼ d 'na' MMMM, y", Medium: "d MMM, y", Short: "dd/MM/yy"},
		Time:     cldr.CalendarDateFormat{},
		DateTime: cldr.CalendarDateFormat{Full: "{1},{0}", Long: "{1}, {0}", Medium: "", Short: ""},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "saŋ tsetsɛ̀ɛ lùm", Feb: "saŋ kàg ngwóŋ", Mar: "saŋ lepyè shúm", Apr: "saŋ cÿó", May: "saŋ tsɛ̀ɛ cÿó", Jun: "saŋ njÿoláʼ", Jul: "saŋ tyɛ̀b tyɛ̀b mbʉ̀", Aug: "saŋ mbʉ̀ŋ", Sep: "saŋ ngwɔ̀ʼ mbÿɛ", Oct: "saŋ tàŋa tsetsáʼ", Nov: "saŋ mejwoŋó", Dec: "saŋ lùm"},
			Narrow:      cldr.CalendarMonthFormatNameValue{},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "saŋ tsetsɛ̀ɛ lùm", Feb: "saŋ kàg ngwóŋ", Mar: "saŋ lepyè shúm", Apr: "saŋ cÿó", May: "saŋ tsɛ̀ɛ cÿó", Jun: "saŋ njÿoláʼ", Jul: "saŋ tyɛ̀b tyɛ̀b mbʉ̀", Aug: "saŋ mbʉ̀ŋ", Sep: "saŋ ngwɔ̀ʼ mbÿɛ", Oct: "saŋ tàŋa tsetsáʼ", Nov: "saŋ mejwoŋó", Dec: "saŋ lùm"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "lyɛʼɛ́ sẅíŋtè", Mon: "mvfò lyɛ̌ʼ", Tue: "mbɔ́ɔntè mvfò lyɛ̌ʼ", Wed: "tsètsɛ̀ɛ lyɛ̌ʼ", Thu: "mbɔ́ɔntè tsetsɛ̀ɛ lyɛ̌ʼ", Fri: "mvfò màga lyɛ̌ʼ", Sat: "màga lyɛ̌ʼ"},
			Narrow:      cldr.CalendarDayFormatNameValue{},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "lyɛʼɛ́ sẅíŋtè", Mon: "mvfò lyɛ̌ʼ", Tue: "mbɔ́ɔntè mvfò lyɛ̌ʼ", Wed: "tsètsɛ̀ɛ lyɛ̌ʼ", Thu: "mbɔ́ɔntè tsetsɛ̀ɛ lyɛ̌ʼ", Fri: "mvfò màga lyɛ̌ʼ", Sat: "màga lyɛ̌ʼ"},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "lyɛʼɛ́ sẅíŋtè", Mon: "mvfò lyɛ̌ʼ", Tue: "mbɔ́ɔntè mvfò lyɛ̌ʼ", Wed: "tsètsɛ̀ɛ lyɛ̌ʼ", Thu: "mbɔ́ɔntè tsetsɛ̀ɛ lyɛ̌ʼ", Fri: "mvfò màga lyɛ̌ʼ", Sat: "màga lyɛ̌ʼ"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "mbaʼámbaʼ", PM: "ncwònzém"},
		},
	},
}
