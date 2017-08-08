package pt_PT

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "", Long: "", Medium: "dd/MM/y", Short: ""},
		Time:     cldr.CalendarDateFormat{},
		DateTime: cldr.CalendarDateFormat{Full: "{1} 'às' {0}", Long: "{1} 'às' {0}", Medium: "{1}, {0}", Short: "{1}, {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{},
		Days:   cldr.CalendarDayFormatNames{},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "a.m.", PM: "p.m."},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "a.m.", PM: "p.m."},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "a.m.", PM: "p.m."},
		},
	},
}
