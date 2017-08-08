package lo

import "github.com/theplant/cldr"

var calendar = cldr.Calendar{
	Formats: cldr.CalendarFormats{
		Date:     cldr.CalendarDateFormat{Full: "EEEE ທີ d MMMM G y", Long: "d MMMM y", Medium: "d MMM y", Short: "d/M/y"},
		Time:     cldr.CalendarDateFormat{Full: "H ໂມງ m ນາທີ ss ວິນາທີ zzzz", Long: "H ໂມງ m ນາທີ ss ວິນາທີ z", Medium: "H:mm:ss", Short: "H:mm"},
		DateTime: cldr.CalendarDateFormat{Full: "{1}, {0}", Long: "{1}, {0}", Medium: "{1}, {0}", Short: "{1}, {0}"},
	},
	FormatNames: cldr.CalendarFormatNames{
		Months: cldr.CalendarMonthFormatNames{
			Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "ມ.ກ.", Feb: "ກ.ພ.", Mar: "ມ.ນ.", Apr: "ມ.ສ.", May: "ພ.ພ.", Jun: "ມິ.ຖ.", Jul: "ກ.ລ.", Aug: "ສ.ຫ.", Sep: "ກ.ຍ.", Oct: "ຕ.ລ.", Nov: "ພ.ຈ.", Dec: "ທ.ວ."},
			Narrow:      cldr.CalendarMonthFormatNameValue{Jan: "1", Feb: "2", Mar: "3", Apr: "4", May: "5", Jun: "6", Jul: "7", Aug: "8", Sep: "9", Oct: "10", Nov: "11", Dec: "12"},
			Short:       cldr.CalendarMonthFormatNameValue{},
			Wide:        cldr.CalendarMonthFormatNameValue{Jan: "ມັງກອນ", Feb: "ກຸມພາ", Mar: "ມີນາ", Apr: "ເມສາ", May: "ພຶດສະພາ", Jun: "ມິຖຸນາ", Jul: "ກໍລະກົດ", Aug: "ສິງຫາ", Sep: "ກັນຍາ", Oct: "ຕຸລາ", Nov: "ພະຈິກ", Dec: "ທັນວາ"},
		},
		Days: cldr.CalendarDayFormatNames{
			Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "ອາທິດ", Mon: "ຈັນ", Tue: "ອັງຄານ", Wed: "ພຸດ", Thu: "ພະຫັດ", Fri: "ສຸກ", Sat: "ເສົາ"},
			Narrow:      cldr.CalendarDayFormatNameValue{Sun: "ອ", Mon: "ຈ", Tue: "ອ", Wed: "ພ", Thu: "ພຫ", Fri: "ສຸ", Sat: "ສ"},
			Short:       cldr.CalendarDayFormatNameValue{Sun: "ອາ.", Mon: "ຈ.", Tue: "ອ.", Wed: "ພ.", Thu: "ພຫ.", Fri: "ສຸ.", Sat: "ສ."},
			Wide:        cldr.CalendarDayFormatNameValue{Sun: "ວັນອາທິດ", Mon: "ວັນຈັນ", Tue: "ວັນອັງຄານ", Wed: "ວັນພຸດ", Thu: "ວັນພະຫັດ", Fri: "ວັນສຸກ", Sat: "ວັນເສົາ"},
		},
		Periods: cldr.CalendarPeriodFormatNames{
			Abbreviated: cldr.CalendarPeriodFormatNameValue{},
			Narrow:      cldr.CalendarPeriodFormatNameValue{AM: "ກທ", PM: "ຫຼທ"},
			Short:       cldr.CalendarPeriodFormatNameValue{},
			Wide:        cldr.CalendarPeriodFormatNameValue{AM: "ກ່ອນທ່ຽງ", PM: "ຫຼັງທ່ຽງ"},
		},
	},
}
