package or

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "d-M-yy"},
		Time:     cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "ଜାନୁଆରୀ", Feb: "ଫେବୃଆରୀ", Mar: "ମାର୍ଚ୍ଚ", Apr: "ଅପ୍ରେଲ", May: "ମଇ", Jun: "ଜୁନ", Jul: "ଜୁଲାଇ", Aug: "ଅଗଷ୍ଟ", Sep: "ସେପ୍ଟେମ୍ବର", Oct: "ଅକ୍ଟୋବର", Nov: "ନଭେମ୍ବର", Dec: "ଡିସେମ୍ବର"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "ଜା", Feb: "ଫେ", Mar: "ମା", Apr: "ଅ", May: "ମଇ", Jun: "ଜୁ", Jul: "ଜୁ", Aug: "ଅ", Sep: "ସେ", Oct: "ଅ", Nov: "ନ", Dec: "ଡି"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "ଜାନୁଆରୀ", Feb: "ଫେବୃଆରୀ", Mar: "ମାର୍ଚ୍ଚ", Apr: "ଅପ୍ରେଲ", May: "ମଇ", Jun: "ଜୁନ", Jul: "ଜୁଲାଇ", Aug: "ଅଗଷ୍ଟ", Sep: "ସେପ୍ଟେମ୍ବର", Oct: "ଅକ୍ଟୋବର", Nov: "ନଭେମ୍ବର", Dec: "ଡିସେମ୍ବର"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "ରବି", Mon: "ସୋମ", Tue: "ମଙ୍ଗଳ", Wed: "ବୁଧ", Thu: "ଗୁରୁ", Fri: "ଶୁକ୍ର", Sat: "ଶନି"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "ର", Mon: "ସୋ", Tue: "ମ", Wed: "ବୁ", Thu: "ଗୁ", Fri: "ଶୁ", Sat: "ଶ"},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "ରବିବାର", Mon: "ସୋମବାର", Tue: "ମଙ୍ଗଳବାର", Wed: "ବୁଧବାର", Thu: "ଗୁରୁବାର", Fri: "ଶୁକ୍ରବାର", Sat: "ଶନିବାର"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "am", PM: "pm"},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "am", PM: "pm"},
		},
	},
}
