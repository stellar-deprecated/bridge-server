package hy

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "yթ. MMMM d, EEEE", Long: "dd MMMM, yթ.", Medium: "dd MMM, yթ.", Short: "dd.MM.yy"},
		Time:     cldr.CalendarDateFormat{Full: "H:mm:ss, zzzz", Long: "H:mm:ss, z", Medium: "H:mm:ss", Short: "H:mm"},
		DateTime: cldr.CalendarDateFormat{Full: "{1}, {0}", Long: "{1}, {0}", Medium: "{1}, {0}", Short: "{1}, {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "հնվ", Feb: "փտվ", Mar: "մրտ", Apr: "ապր", May: "մյս", Jun: "հնս", Jul: "հլս", Aug: "օգս", Sep: "սեպ", Oct: "հոկ", Nov: "նոյ", Dec: "դեկ"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "Հ", Feb: "Փ", Mar: "Մ", Apr: "Ա", May: "Մ", Jun: "Հ", Jul: "Հ", Aug: "Օ", Sep: "Ս", Oct: "Հ", Nov: "Ն", Dec: "Դ"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "հունվար", Feb: "փետրվար", Mar: "մարտ", Apr: "ապրիլ", May: "մայիս", Jun: "հունիս", Jul: "հուլիս", Aug: "օգոստոս", Sep: "սեպտեմբեր", Oct: "հոկտեմբեր", Nov: "նոյեմբեր", Dec: "դեկտեմբեր"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "կիր", Mon: "երկ", Tue: "երք", Wed: "չրք", Thu: "հնգ", Fri: "ուր", Sat: "շբթ"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "Կ", Mon: "Ե", Tue: "Ե", Wed: "Չ", Thu: "Հ", Fri: "Ու", Sat: "Շ"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "կիր", Mon: "երկ", Tue: "երք", Wed: "չրք", Thu: "հնգ", Fri: "ուր", Sat: "շբթ"},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "կիրակի", Mon: "երկուշաբթի", Tue: "երեքշաբթի", Wed: "չորեքշաբթի", Thu: "հինգշաբթի", Fri: "ուրբաթ", Sat: "շաբաթ"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "կա", PM: "կհ"},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "կեսօրից առաջ", PM: "կեսօրից հետո"},
		},
	},
}
