package te

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "d, MMMM y, EEEE", Long: "d MMMM, y", Medium: "d MMM, y", Short: "dd-MM-yy"},
		Time:     cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
		DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "జన", Feb: "ఫిబ్ర", Mar: "మార్చి", Apr: "ఏప్రి", May: "మే", Jun: "జూన్", Jul: "జులై", Aug: "ఆగస్టు", Sep: "సెప్టెం", Oct: "అక్టో", Nov: "నవం", Dec: "డిసెం"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "జ", Feb: "ఫి", Mar: "మా", Apr: "ఏ", May: "మే", Jun: "జూ", Jul: "జు", Aug: "ఆ", Sep: "సె", Oct: "అ", Nov: "న", Dec: "డి"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "జనవరి", Feb: "ఫిబ్రవరి", Mar: "మార్చి", Apr: "ఏప్రిల్", May: "మే", Jun: "జూన్", Jul: "జులై", Aug: "ఆగస్టు", Sep: "సెప్టెంబర్", Oct: "అక్టోబర్", Nov: "నవంబర్", Dec: "డిసెంబర్"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "ఆది", Mon: "సోమ", Tue: "మంగళ", Wed: "బుధ", Thu: "గురు", Fri: "శుక్ర", Sat: "శని"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "ఆ", Mon: "సో", Tue: "మ", Wed: "బు", Thu: "గు", Fri: "శు", Sat: "శ"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "ఆది", Mon: "సోమ", Tue: "మం", Wed: "బుధ", Thu: "గురు", Fri: "శుక్ర", Sat: "శని"},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "ఆదివారం", Mon: "సోమవారం", Tue: "మంగళవారం", Wed: "బుధవారం", Thu: "గురువారం", Fri: "శుక్రవారం", Sat: "శనివారం"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "ఉ", PM: "సా"},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "[AM]", PM: "[PM]"},
		},
	},
}
