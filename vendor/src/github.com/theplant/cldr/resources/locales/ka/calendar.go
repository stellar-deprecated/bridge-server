package ka

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE, dd MMMM, y", Long: "d MMMM, y", Medium: "d MMM, y", Short: "dd.MM.yy"},
		Time:     cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
		DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1}, {0}", Medium: "{1} {0}", Short: "{1}, {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "იან", Feb: "თებ", Mar: "მარ", Apr: "აპრ", May: "მაი", Jun: "ივნ", Jul: "ივლ", Aug: "აგვ", Sep: "სექ", Oct: "ოქტ", Nov: "ნოე", Dec: "დეკ"},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "ი", Feb: "თ", Mar: "მ", Apr: "ა", May: "მ", Jun: "ი", Jul: "ი", Aug: "ა", Sep: "ს", Oct: "ო", Nov: "ნ", Dec: "დ"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "იანვარი", Feb: "თებერვალი", Mar: "მარტი", Apr: "აპრილი", May: "მაისი", Jun: "ივნისი", Jul: "ივლისი", Aug: "აგვისტო", Sep: "სექტემბერი", Oct: "ოქტომბერი", Nov: "ნოემბერი", Dec: "დეკემბერი"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "კვი", Mon: "ორშ", Tue: "სამ", Wed: "ოთხ", Thu: "ხუთ", Fri: "პარ", Sat: "შაბ"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "კ", Mon: "ო", Tue: "ს", Wed: "ო", Thu: "ხ", Fri: "პ", Sat: "შ"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "კვ", Mon: "ორ", Tue: "სმ", Wed: "ოთ", Thu: "ხთ", Fri: "პრ", Sat: "შბ"},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "კვირა", Mon: "ორშაბათი", Tue: "სამშაბათი", Wed: "ოთხშაბათი", Thu: "ხუთშაბათი", Fri: "პარასკევი", Sat: "შაბათი"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "a", PM: "p"},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"},
		},
	},
}
