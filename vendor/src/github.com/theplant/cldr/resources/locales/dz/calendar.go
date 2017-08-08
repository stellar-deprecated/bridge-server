package dz

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, སྤྱི་ལོ་y MMMM ཚེས་dd", Long: "སྤྱི་ལོ་y MMMM ཚེས་ dd", Medium: "སྤྱི་ལོ་y ཟླ་MMM ཚེས་dd", Short: "y-MM-dd"},
		Time:     cldr.CalendarDateFormat{Full: "ཆུ་ཚོད་ h སྐར་མ་ mm:ss a zzzz", Long: "ཆུ་ཚོད་ h སྐར་མ་ mm:ss a z", Medium: "ཆུ་ཚོད་h:mm:ss a", Short: "ཆུ་ཚོད་ h སྐར་མ་ mm a"},
		DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "ཟླ་༡", Feb: "ཟླ་༢", Mar: "ཟླ་༣", Apr: "ཟླ་༤", May: "ཟླ་༥", Jun: "ཟླ་༦", Jul: "ཟླ་༧", Aug: "ཟླ་༨", Sep: "ཟླ་༩", Oct: "ཟླ་༡༠", Nov: "ཟླ་༡༡", Dec: "ཟླ་༡༢"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "༡", Feb: "༢", Mar: "༣", Apr: "༤", May: "༥", Jun: "༦", Jul: "༧", Aug: "༨", Sep: "༩", Oct: "༡༠", Nov: "༡༡", Dec: "༡༢"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "སྤྱི་ཟླ་དངཔ་", Feb: "སྤྱི་ཟླ་གཉིས་པ་", Mar: "སྤྱི་ཟླ་གསུམ་པ་", Apr: "སྤྱི་ཟླ་བཞི་པ", May: "སྤྱི་ཟླ་ལྔ་པ་", Jun: "སྤྱི་ཟླ་དྲུག་པ", Jul: "སྤྱི་ཟླ་བདུན་པ་", Aug: "སྤྱི་ཟླ་བརྒྱད་པ་", Sep: "སྤྱི་ཟླ་དགུ་པ་", Oct: "སྤྱི་ཟླ་བཅུ་པ་", Nov: "སྤྱི་ཟླ་བཅུ་གཅིག་པ་", Dec: "སྤྱི་ཟླ་བཅུ་གཉིས་པ་"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "ཟླ་", Mon: "མིར་", Tue: "ལྷག་", Wed: "ཕུར་", Thu: "སངས་", Fri: "སྤེན་", Sat: "ཉི་"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "ཟླ", Mon: "མིར", Tue: "ལྷག", Wed: "ཕུར", Thu: "སངྶ", Fri: "སྤེན", Sat: "ཉི"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "ཟླ་", Mon: "མིར་", Tue: "ལྷག་", Wed: "ཕུར་", Thu: "སངས་", Fri: "སྤེན་", Sat: "ཉི་"},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "གཟའ་ཟླ་བ་", Mon: "གཟའ་མིག་དམར་", Tue: "གཟའ་ལྷག་པ་", Wed: "གཟའ་ཕུར་བུ་", Thu: "གཟའ་པ་སངས་", Fri: "གཟའ་སྤེན་པ་", Sat: "གཟའ་ཉི་མ་"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "སྔ་ཆ་", PM: "ཕྱི་ཆ་"},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "སྔ་ཆ་", PM: "ཕྱི་ཆ་"},
		},
	},
}
