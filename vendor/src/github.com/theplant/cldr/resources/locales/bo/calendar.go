package bo

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "", Long: "སྤྱི་ལོ་y MMMMའི་ཙེས་dད", Medium: "y ལོ་འི་MMMཙེས་d", Short: ""},
		Time:     cldr.CalendarDateFormat{},
		DateTime: cldr.CalendarDateFormat{},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "ཟླ་༡", Feb: "ཟླ་༢", Mar: "ཟླ་༣", Apr: "ཟླ་༤", May: "ཟླ་༥", Jun: "ཟླ་༦", Jul: "ཟླ་༧", Aug: "ཟླ་༨", Sep: "ཟླ་༩", Oct: "ཟླ་༡༠", Nov: "ཟླ་༡༡", Dec: "ཟླ་༡༢"},
			Narrow:      cldr.CalendarMonthFormatNameValue{},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "ཟླ་བ་དང་པོ་", Feb: "ཟླ་བ་གཉིས་པ་", Mar: "ཟླ་བ་སུམ་པ་", Apr: "ཟླ་བ་བཞི་པ་", May: "ཟླ་བ་ལྔ་པ་", Jun: "ཟླ་བ་དྲུག་པ་", Jul: "ཟླ་བ་བདུན་པ་", Aug: "ཟླ་བ་བརྒྱད་པ་", Sep: "ཟླ་བ་དགུ་པ་", Oct: "ཟླ་བ་བཅུ་པ་", Nov: "ཟླ་བ་བཅུ་གཅིག་པ་", Dec: "ཟླ་བ་བཅུ་གཉིས་པ་"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "ཉི་མ་", Mon: "ཟླ་བ་", Tue: "མིག་དམར་", Wed: "ལྷག་པ་", Thu: "ཕུར་བུ་", Fri: "པ་སངས་", Sat: "སྤེན་པ་"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "ཉི", Mon: "ཟླ", Tue: "མི", Wed: "ལྷ", Thu: "ཕུ", Fri: "པ", Sat: "སྤེ"},
			Short:       cldr.CalendarDayFormatNameValue{},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "གཟའ་ཉི་མ་", Mon: "གཟའ་ཟླ་བ་", Tue: "གཟའ་མིག་དམར་", Wed: "གཟའ་ལྷག་པ་", Thu: "གཟའ་ཕུར་བུ་", Fri: "གཟའ་པ་སངས་", Sat: "གཟའ་སྤེན་པ་"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "སྔ་དྲོ་", PM: "ཕྱི་དྲོ་"},
		},
	},
}
