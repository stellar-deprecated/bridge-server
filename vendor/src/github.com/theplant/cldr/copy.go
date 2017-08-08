package cldr

// Copy only override dst value if src value is not a zero value
//
// TODO: write a auto generator
func Copy(dst, src *Locale) {
	if src.Locale != "" {
		dst.Locale = src.Locale
	}
	if src.PluralRule != "" {
		dst.PluralRule = src.PluralRule
	}
	// Calendar   Calendar

	copyNumber(&dst.Number, &src.Number)
	copyCalendar(&dst.Calendar, &src.Calendar)
}

func copyNumber(dst, src *Number) {
	if src.Symbols.Decimal != "" {
		dst.Symbols.Decimal = src.Symbols.Decimal
	}
	if src.Symbols.Group != "" {
		dst.Symbols.Group = src.Symbols.Group
	}
	if src.Symbols.Negative != "" {
		dst.Symbols.Negative = src.Symbols.Negative
	}
	if src.Symbols.Percent != "" {
		dst.Symbols.Percent = src.Symbols.Percent
	}
	if src.Symbols.PerMille != "" {
		dst.Symbols.PerMille = src.Symbols.PerMille
	}

	if src.Formats.Decimal != "" {
		dst.Formats.Decimal = src.Formats.Decimal
	}
	if src.Formats.Currency != "" {
		dst.Formats.Currency = src.Formats.Currency
	}
	if src.Formats.Percent != "" {
		dst.Formats.Percent = src.Formats.Percent
	}

curLoop:
	for _, scur := range src.Currencies {
		for i, dcur := range dst.Currencies {
			if dcur.Currency == scur.Currency {
				dst.Currencies[i] = scur
				continue curLoop
			}
		}

		dst.Currencies = append(dst.Currencies, scur)
	}
}

func copyCalendar(dst, src *Calendar) {
	copyCalendarDateFormat(&dst.Formats.Date, &src.Formats.Date)
	copyCalendarDateFormat(&dst.Formats.Time, &src.Formats.Time)
	copyCalendarDateFormat(&dst.Formats.DateTime, &src.Formats.DateTime)

	copyCalendarMonthFormatNames(&dst.FormatNames.Months, &src.FormatNames.Months)
	copyCalendarDayFormatNames(&dst.FormatNames.Days, &src.FormatNames.Days)
	copyCalendarPeriodFormatNames(&dst.FormatNames.Periods, &src.FormatNames.Periods)
}

func copyCalendarMonthFormatNames(dst, src *CalendarMonthFormatNames) {
	copyCalendarMonthFormatNameValue(&dst.Abbreviated, &src.Abbreviated)
	copyCalendarMonthFormatNameValue(&dst.Narrow, &src.Narrow)
	copyCalendarMonthFormatNameValue(&dst.Short, &src.Short)
	copyCalendarMonthFormatNameValue(&dst.Wide, &src.Wide)
}

func copyCalendarDayFormatNames(dst, src *CalendarDayFormatNames) {
	copyCalendarDayFormatNameValue(&dst.Abbreviated, &src.Abbreviated)
	copyCalendarDayFormatNameValue(&dst.Narrow, &src.Narrow)
	copyCalendarDayFormatNameValue(&dst.Short, &src.Short)
	copyCalendarDayFormatNameValue(&dst.Wide, &src.Wide)
}

func copyCalendarPeriodFormatNames(dst, src *CalendarPeriodFormatNames) {
	copyCalendarPeriodFormatNameValue(&dst.Abbreviated, &src.Abbreviated)
	copyCalendarPeriodFormatNameValue(&dst.Narrow, &src.Narrow)
	copyCalendarPeriodFormatNameValue(&dst.Short, &src.Short)
	copyCalendarPeriodFormatNameValue(&dst.Wide, &src.Wide)
}

func copyCalendarDateFormat(dst, src *CalendarDateFormat) {
	if src.Full != "" {
		dst.Full = src.Full
	}
	if src.Long != "" {
		dst.Long = src.Long
	}
	if src.Medium != "" {
		dst.Medium = src.Medium
	}
	if src.Short != "" {
		dst.Short = src.Short
	}
}

func copyCalendarDayFormatNameValue(dst, src *CalendarDayFormatNameValue) {
	if src.Sun != "" {
		dst.Sun = src.Sun
	}
	if src.Mon != "" {
		dst.Mon = src.Mon
	}
	if src.Tue != "" {
		dst.Tue = src.Tue
	}
	if src.Wed != "" {
		dst.Wed = src.Wed
	}
	if src.Thu != "" {
		dst.Thu = src.Thu
	}
	if src.Fri != "" {
		dst.Fri = src.Fri
	}
	if src.Sat != "" {
		dst.Sat = src.Sat
	}
}

func copyCalendarMonthFormatNameValue(dst, src *CalendarMonthFormatNameValue) {
	if src.Jan != "" {
		dst.Jan = src.Jan
	}
	if src.Feb != "" {
		dst.Feb = src.Feb
	}
	if src.Mar != "" {
		dst.Mar = src.Mar
	}
	if src.Apr != "" {
		dst.Apr = src.Apr
	}
	if src.May != "" {
		dst.May = src.May
	}
	if src.Jun != "" {
		dst.Jun = src.Jun
	}
	if src.Jul != "" {
		dst.Jul = src.Jul
	}
	if src.Aug != "" {
		dst.Aug = src.Aug
	}
	if src.Sep != "" {
		dst.Sep = src.Sep
	}
	if src.Oct != "" {
		dst.Oct = src.Oct
	}
	if src.Nov != "" {
		dst.Nov = src.Nov
	}
	if src.Dec != "" {
		dst.Dec = src.Dec
	}
}

func copyCalendarPeriodFormatNameValue(dst, src *CalendarPeriodFormatNameValue) {
	if src.AM != "" {
		dst.AM = src.AM
	}
	if src.PM != "" {
		dst.PM = src.PM
	}
}
