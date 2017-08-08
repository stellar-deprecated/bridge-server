package br

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Gen", Feb: "Cʼhwe", Mar: "Meur", Apr: "Ebr", May: "Mae", Jun: "Mezh", Jul: "Goue", Aug: "Eost", Sep: "Gwen", Oct: "Here", Nov: "Du", Dec: "Ker"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "01", Feb: "02", Mar: "03", Apr: "04", May: "05", Jun: "06", Jul: "07", Aug: "08", Sep: "09", Oct: "10", Nov: "11", Dec: "12"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "Genver", Feb: "Cʼhwevrer", Mar: "Meurzh", Apr: "Ebrel", May: "Mae", Jun: "Mezheven", Jul: "Gouere", Aug: "Eost", Sep: "Gwengolo", Oct: "Here", Nov: "Du", Dec: "Kerzu"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Sul", Mon: "Lun", Tue: "Meu.", Wed: "Mer.", Thu: "Yaou", Fri: "Gwe.", Sat: "Sad."},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "Su", Mon: "L", Tue: "Mz", Wed: "Mc", Thu: "Y", Fri: "G", Sat: "Sa"},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "Sul", Mon: "Lun", Tue: "Meurzh", Wed: "Mercʼher", Thu: "Yaou", Fri: "Gwener", Sat: "Sadorn"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "am", PM: "gm"},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "A.M.", PM: "G.M."},
		},
	},
}
