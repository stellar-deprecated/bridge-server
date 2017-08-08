package ii

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "1", Feb: "2", Mar: "3", Apr: "4", May: "5", Jun: "6", Jul: "7", Aug: "8", Sep: "9", Oct: "10", Nov: "11", Dec: "12"},
			Narrow:      cldr.CalendarMonthFormatNameValue{},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "ꋍꆪ", Feb: "ꑍꆪ", Mar: "ꌕꆪ", Apr: "ꇖꆪ", May: "ꉬꆪ", Jun: "ꃘꆪ", Jul: "ꏃꆪ", Aug: "ꉆꆪ", Sep: "ꈬꆪ", Oct: "ꊰꆪ", Nov: "ꊰꊪꆪ", Dec: "ꊰꑋꆪ"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "ꑭꆏ", Mon: "ꆏꋍ", Tue: "ꆏꑍ", Wed: "ꆏꌕ", Thu: "ꆏꇖ", Fri: "ꆏꉬ", Sat: "ꆏꃘ"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "ꆏ", Mon: "ꋍ", Tue: "ꑍ", Wed: "ꌕ", Thu: "ꇖ", Fri: "ꉬ", Sat: "ꃘ"},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "ꑭꆏꑍ", Mon: "ꆏꊂꋍ", Tue: "ꆏꊂꑍ", Wed: "ꆏꊂꌕ", Thu: "ꆏꊂꇖ", Fri: "ꆏꊂꉬ", Sat: "ꆏꊂꃘ"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "ꎸꄑ", PM: "ꁯꋒ"},
		},
	},
}
