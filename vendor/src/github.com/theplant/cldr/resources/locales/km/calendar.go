package km

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "d/M/yy"},
		Time:     cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
		DateTime: cldr.CalendarDateFormat{Full: "{1} នៅ {0}", Long: "{1} នៅ {0}", Medium: "{1}, {0}", Short: "{1}, {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "មករា", Feb: "កុម្ភៈ", Mar: "មីនា", Apr: "មេសា", May: "ឧសភា", Jun: "មិថុនា", Jul: "កក្កដា", Aug: "សីហា", Sep: "កញ្ញា", Oct: "តុលា", Nov: "វិច្ឆិកា", Dec: "ធ្នូ"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "1", Feb: "2", Mar: "3", Apr: "4", May: "5", Jun: "6", Jul: "7", Aug: "8", Sep: "9", Oct: "10", Nov: "11", Dec: "12"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "មករា", Feb: "កុម្ភៈ", Mar: "មីនា", Apr: "មេសា", May: "ឧសភា", Jun: "មិថុនា", Jul: "កក្កដា", Aug: "សីហា", Sep: "កញ្ញា", Oct: "តុលា", Nov: "វិច្ឆិកា", Dec: "ធ្នូ"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "អាទិត្យ", Mon: "ចន្ទ", Tue: "អង្គារ", Wed: "ពុធ", Thu: "ព្រហស្បតិ៍", Fri: "សុក្រ", Sat: "សៅរ៍"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "1", Mon: "2", Tue: "3", Wed: "4", Thu: "5", Fri: "6", Sat: "7"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "អាទិត្យ", Mon: "ចន្ទ", Tue: "អង្គារ", Wed: "ពុធ", Thu: "ព្រហស្បតិ៍", Fri: "សុក្រ", Sat: "សៅរ៍"},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "អាទិត្យ", Mon: "ចន្ទ", Tue: "អង្គារ", Wed: "ពុធ", Thu: "ព្រហស្បតិ៍", Fri: "សុក្រ", Sat: "សៅរ៍"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "ព្រឹក", PM: "ល្ងាច"},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "ព្រឹក", PM: "ល្ងាច"},
		},
	},
}
