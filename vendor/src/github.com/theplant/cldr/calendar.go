package cldr

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

// Standard Formats for Dates, Times & DateTimes
// These are the options to pass to the Format method.
const (
	DateFormatFull = iota
	DateFormatLong
	DateFormatMedium
	DateFormatShort
	TimeFormatFull
	TimeFormatLong
	TimeFormatMedium
	TimeFormatShort
	DateTimeFormatFull
	DateTimeFormatLong
	DateTimeFormatMedium
	DateTimeFormatShort
)

// Characters with special meaning in a datetime string:
// Technically, all a-z,A-Z characters should be treated as if they represent a
// datetime unit - but not all actually do. Any a-z,A-Z character that is
// intended to be rendered as a literal a-z,A-Z character should be surrounded
// by single quotes. There is currently no support for rendering a single quote
// literal.
const (
	datetimeFormatUnitEra       = 'G'
	datetimeFormatUnitYear      = 'y'
	datetimeFormatUnitMonth     = 'M'
	datetimeFormatUnitDayOfWeek = 'E'
	datetimeFormatUnitDay       = 'd'
	datetimeFormatUnitHour12    = 'h'
	datetimeFormatUnitHour24    = 'H'
	datetimeFormatUnitMinute    = 'm'
	datetimeFormatUnitSecond    = 's'
	datetimeFormatUnitPeriod    = 'a'
	datetimeForamtUnitQuarter   = 'Q'
	datetimeFormatUnitTimeZone1 = 'z'
	datetimeFormatUnitTimeZone2 = 'v'

	datetimeFormatTimeSeparator = ':'
	datetimeFormatLiteral       = '\''
)

// The sequence length of datetime unit characters indicates how they should be
// rendered.
const (
	datetimeFormatLength1Plus       = 1
	datetimeFormatLength2Plus       = 2
	datetimeFormatLengthAbbreviated = 3
	datetimeFormatLengthWide        = 4
	datetimeFormatLengthNarrow      = 5
)

// datetime formats are a sequences off datetime components and string literals
const (
	datetimePatternComponentUnit = iota
	datetimePatternComponentLiteral
)

// A list of currently unsupported units:
// These still need to be implemented. For now they are ignored.
var (
	datetimeFormatUnitCutset = []rune{
		datetimeFormatUnitEra,
		datetimeForamtUnitQuarter,
		datetimeFormatUnitTimeZone1,
		datetimeFormatUnitTimeZone2,
	}
)

type datetimePatternComponent struct {
	pattern       string
	componentType int
}

func (c Calendar) FmtDateFull(t time.Time) (string, error) {
	return c.Format(t, c.Formats.Date.Full)
}

func (c Calendar) FmtDateLong(t time.Time) (string, error) {
	return c.Format(t, c.Formats.Date.Long)
}

func (c Calendar) FmtDateMedium(t time.Time) (string, error) {
	return c.Format(t, c.Formats.Date.Medium)
}

func (c Calendar) FmtDateShort(t time.Time) (string, error) {
	return c.Format(t, c.Formats.Date.Short)
}

func (c Calendar) FmtTimeFull(t time.Time) (string, error) {
	return c.Format(t, c.Formats.Time.Full)
}

func (c Calendar) FmtTimeLong(t time.Time) (string, error) {
	return c.Format(t, c.Formats.Time.Long)
}

func (c Calendar) FmtTimeMedium(t time.Time) (string, error) {
	return c.Format(t, c.Formats.Time.Medium)
}

func (c Calendar) FmtTimeShort(t time.Time) (string, error) {
	return c.Format(t, c.Formats.Time.Short)
}

func (c Calendar) FmtDateTimeFull(t time.Time) (string, error) {
	pattern := getDateTimePattern(c.Formats.DateTime.Full, c.Formats.Date.Full, c.Formats.Time.Full)
	return c.Format(t, pattern)
}

func (c Calendar) FmtDateTimeLong(t time.Time) (string, error) {
	pattern := getDateTimePattern(c.Formats.DateTime.Long, c.Formats.Date.Long, c.Formats.Time.Long)
	return c.Format(t, pattern)
}

func (c Calendar) FmtDateTimeMedium(t time.Time) (string, error) {
	pattern := getDateTimePattern(c.Formats.DateTime.Medium, c.Formats.Date.Medium, c.Formats.Time.Medium)
	return c.Format(t, pattern)
}

func (c Calendar) FmtDateTimeShort(t time.Time) (string, error) {
	pattern := getDateTimePattern(c.Formats.DateTime.Short, c.Formats.Date.Short, c.Formats.Time.Short)
	return c.Format(t, pattern)
}

// Format takes a time struct and a format and returns a formatted
// string. Callers should use a DateFormat, TimeFormat, or DateTimeFormat
// constant.
func (c Calendar) Format(datetime time.Time, pattern string) (string, error) {
	parsed, err := c.parseDateTimeFormat(pattern)
	if err != nil {
		return "", err
	}

	return c.formatDateTime(datetime, parsed)
}

// formatDateTime takes a time.Time and a sequence of parsed pattern components
// and returns an internationalized string representation.
func (c Calendar) formatDateTime(datetime time.Time, pattern []*datetimePatternComponent) (string, error) {
	formatted := ""
	for _, component := range pattern {
		if component.componentType == datetimePatternComponentLiteral {
			formatted += component.pattern
		} else {
			f, err := c.formatDateTimeComponent(datetime, component.pattern)
			if err != nil {
				return "", err
			}
			formatted += f
		}
	}

	return strings.Trim(formatted, " ,"), nil
}

// formatDateTimeComponent renders a single component of a datetime format
// pattern.
func (c Calendar) formatDateTimeComponent(datetime time.Time, pattern string) (string, error) {
	switch pattern[0:1] {
	case string(datetimeFormatUnitEra):
		return c.formatDateTimeComponentEra(datetime, len(pattern))
	case string(datetimeFormatUnitYear):
		return c.formatDateTimeComponentYear(datetime, len(pattern))
	case string(datetimeFormatUnitMonth):
		return c.formatDateTimeComponentMonth(datetime, len(pattern))
	case string(datetimeFormatUnitDayOfWeek):
		return c.formatDateTimeComponentDayOfWeek(datetime, len(pattern))
	case string(datetimeFormatUnitDay):
		return c.formatDateTimeComponentDay(datetime, len(pattern))
	case string(datetimeFormatUnitHour12):
		return c.formatDateTimeComponentHour12(datetime, len(pattern))
	case string(datetimeFormatUnitHour24):
		return c.formatDateTimeComponentHour24(datetime, len(pattern))
	case string(datetimeFormatUnitMinute):
		return c.formatDateTimeComponentMinute(datetime, len(pattern))
	case string(datetimeFormatUnitSecond):
		return c.formatDateTimeComponentSecond(datetime, len(pattern))
	case string(datetimeFormatUnitPeriod):
		return c.formatDateTimeComponentPeriod(datetime, len(pattern))
	case string(datetimeForamtUnitQuarter):
		return c.formatDateTimeComponentQuarter(datetime, len(pattern))
	case string(datetimeFormatUnitTimeZone1):
		fallthrough
	case string(datetimeFormatUnitTimeZone2):
		return c.formatDateTimeComponentTimeZone(datetime, len(pattern))
	}

	return "", errors.New("unknown datetime format unit: " + pattern[0:1])
}

// formatDateTimeComponentEra renders an era component.
// TODO: not yet implemented
func (c Calendar) formatDateTimeComponentEra(datetime time.Time, length int) (string, error) {
	return "", nil
}

// formatDateTimeComponentYear renders a year component.
func (c Calendar) formatDateTimeComponentYear(datetime time.Time, length int) (string, error) {
	year := datetime.Year()
	switch length {
	case datetimeFormatLength1Plus:
		return c.formatDateTimeComponentYearLengthWide(year), nil
	case datetimeFormatLength2Plus:
		return c.formatDateTimeComponentYearLength2Plus(year), nil
	case datetimeFormatLengthWide:
		return c.formatDateTimeComponentYearLengthWide(year), nil
	}

	return "", fmt.Errorf("unsupported year length: %d", length)
}

// formatDateTimeComponentYearLength2Plus renders a 2-digit year component.
func (c Calendar) formatDateTimeComponentYearLength2Plus(year int) string {
	yearShort := year % 100

	if yearShort < 10 {
		return fmt.Sprintf("0%d", yearShort)
	}

	return fmt.Sprintf("%d", yearShort)
}

// formatDateTimeComponentYearLength2Plus renders a full-year component - for
// all modern dates, that's four digits.
func (c Calendar) formatDateTimeComponentYearLengthWide(year int) string {
	return fmt.Sprintf("%d", year)
}

// formatDateTimeComponentMonth renders a month component.
func (c Calendar) formatDateTimeComponentMonth(datetime time.Time, length int) (string, error) {

	month := int(datetime.Month())

	switch length {
	case datetimeFormatLength1Plus:
		return c.formatDateTimeComponentMonth1Plus(month), nil
	case datetimeFormatLength2Plus:
		return c.formatDateTimeComponentMonth2Plus(month), nil
	case datetimeFormatLengthAbbreviated:
		return c.formatDateTimeComponentMonthAbbreviated(month), nil
	case datetimeFormatLengthWide:
		return c.formatDateTimeComponentMonthWide(month), nil
	case datetimeFormatLengthNarrow:
		return c.formatDateTimeComponentMonthNarrow(month), nil
	}

	return "", fmt.Errorf("unsupported month length: %d", length)
}

// formatDateTimeComponentMonth1Plus renders a numeric month component with 1 or
// 2 digits depending on value.
func (c Calendar) formatDateTimeComponentMonth1Plus(month int) string {
	return fmt.Sprintf("%d", month)
}

// formatDateTimeComponentMonth2Plus renders a numeric month component always
// with 2 digits.
func (c Calendar) formatDateTimeComponentMonth2Plus(month int) string {
	if month < 10 {
		return fmt.Sprintf("0%d", month)
	}
	return fmt.Sprintf("%d", month)
}

// formatDateTimeComponentMonthAbbreviated renders an abbreviated text month
// component.
func (c Calendar) formatDateTimeComponentMonthAbbreviated(month int) string {
	switch month {
	case 1:
		return c.FormatNames.Months.Abbreviated.Jan
	case 2:
		return c.FormatNames.Months.Abbreviated.Feb
	case 3:
		return c.FormatNames.Months.Abbreviated.Mar
	case 4:
		return c.FormatNames.Months.Abbreviated.Apr
	case 5:
		return c.FormatNames.Months.Abbreviated.May
	case 6:
		return c.FormatNames.Months.Abbreviated.Jun
	case 7:
		return c.FormatNames.Months.Abbreviated.Jul
	case 8:
		return c.FormatNames.Months.Abbreviated.Aug
	case 9:
		return c.FormatNames.Months.Abbreviated.Sep
	case 10:
		return c.FormatNames.Months.Abbreviated.Oct
	case 11:
		return c.FormatNames.Months.Abbreviated.Nov
	case 12:
		return c.FormatNames.Months.Abbreviated.Dec
	}

	return ""
}

// formatDateTimeComponentMonthWide renders a full text month component.
func (c Calendar) formatDateTimeComponentMonthWide(month int) string {
	switch month {
	case 1:
		return c.FormatNames.Months.Wide.Jan
	case 2:
		return c.FormatNames.Months.Wide.Feb
	case 3:
		return c.FormatNames.Months.Wide.Mar
	case 4:
		return c.FormatNames.Months.Wide.Apr
	case 5:
		return c.FormatNames.Months.Wide.May
	case 6:
		return c.FormatNames.Months.Wide.Jun
	case 7:
		return c.FormatNames.Months.Wide.Jul
	case 8:
		return c.FormatNames.Months.Wide.Aug
	case 9:
		return c.FormatNames.Months.Wide.Sep
	case 10:
		return c.FormatNames.Months.Wide.Oct
	case 11:
		return c.FormatNames.Months.Wide.Nov
	case 12:
		return c.FormatNames.Months.Wide.Dec
	}

	return ""
}

// formatDateTimeComponentMonthNarrow renders a super-short month compontent -
// not guaranteed to be unique for different months.
func (c Calendar) formatDateTimeComponentMonthNarrow(month int) string {
	switch month {
	case 1:
		return c.FormatNames.Months.Narrow.Jan
	case 2:
		return c.FormatNames.Months.Narrow.Feb
	case 3:
		return c.FormatNames.Months.Narrow.Mar
	case 4:
		return c.FormatNames.Months.Narrow.Apr
	case 5:
		return c.FormatNames.Months.Narrow.May
	case 6:
		return c.FormatNames.Months.Narrow.Jun
	case 7:
		return c.FormatNames.Months.Narrow.Jul
	case 8:
		return c.FormatNames.Months.Narrow.Aug
	case 9:
		return c.FormatNames.Months.Narrow.Sep
	case 10:
		return c.FormatNames.Months.Narrow.Oct
	case 11:
		return c.FormatNames.Months.Narrow.Nov
	case 12:
		return c.FormatNames.Months.Narrow.Dec
	}

	return ""
}

// formatDateTimeComponentDayOfWeek renders a day-of-week component.
func (c Calendar) formatDateTimeComponentDayOfWeek(datetime time.Time, length int) (string, error) {
	switch length {
	case datetimeFormatLength1Plus:
		return c.formatDateTimeComponentDayOfWeekWide(datetime.Weekday()), nil
	case datetimeFormatLength2Plus:
		return c.formatDateTimeComponentDayOfWeekShort(datetime.Weekday()), nil
	case datetimeFormatLengthAbbreviated:
		return c.formatDateTimeComponentDayOfWeekAbbreviated(datetime.Weekday()), nil
	case datetimeFormatLengthWide:
		return c.formatDateTimeComponentDayOfWeekWide(datetime.Weekday()), nil
	case datetimeFormatLengthNarrow:
		return c.formatDateTimeComponentDayOfWeekNarrow(datetime.Weekday()), nil
	}

	return "", fmt.Errorf("unsupported year day-of-week: %d", length)
}

// formatDateTimeComponentDayOfWeekAbbreviated renders an abbreviated text
// day-of-week component.
func (c Calendar) formatDateTimeComponentDayOfWeekAbbreviated(dayOfWeek time.Weekday) string {
	switch dayOfWeek {
	case time.Sunday:
		return c.FormatNames.Days.Abbreviated.Sun
	case time.Monday:
		return c.FormatNames.Days.Abbreviated.Mon
	case time.Tuesday:
		return c.FormatNames.Days.Abbreviated.Tue
	case time.Wednesday:
		return c.FormatNames.Days.Abbreviated.Wed
	case time.Thursday:
		return c.FormatNames.Days.Abbreviated.Thu
	case time.Friday:
		return c.FormatNames.Days.Abbreviated.Fri
	case time.Saturday:
		return c.FormatNames.Days.Abbreviated.Sat
	}

	return ""
}

// formatDateTimeComponentDayOfWeekAbbreviated renders a
// shorter-then-abbreviated but still unique text day-of-week component.
func (c Calendar) formatDateTimeComponentDayOfWeekShort(dayOfWeek time.Weekday) string {
	switch dayOfWeek {
	case time.Sunday:
		return c.FormatNames.Days.Short.Sun
	case time.Monday:
		return c.FormatNames.Days.Short.Mon
	case time.Tuesday:
		return c.FormatNames.Days.Short.Tue
	case time.Wednesday:
		return c.FormatNames.Days.Short.Wed
	case time.Thursday:
		return c.FormatNames.Days.Short.Thu
	case time.Friday:
		return c.FormatNames.Days.Short.Fri
	case time.Saturday:
		return c.FormatNames.Days.Short.Sat
	}

	return ""
}

// formatDateTimeComponentDayOfWeekWide renders a full text day-of-week
// component.
func (c Calendar) formatDateTimeComponentDayOfWeekWide(dayOfWeek time.Weekday) string {
	switch dayOfWeek {
	case time.Sunday:
		return c.FormatNames.Days.Wide.Sun
	case time.Monday:
		return c.FormatNames.Days.Wide.Mon
	case time.Tuesday:
		return c.FormatNames.Days.Wide.Tue
	case time.Wednesday:
		return c.FormatNames.Days.Wide.Wed
	case time.Thursday:
		return c.FormatNames.Days.Wide.Thu
	case time.Friday:
		return c.FormatNames.Days.Wide.Fri
	case time.Saturday:
		return c.FormatNames.Days.Wide.Sat
	}

	return ""
}

// formatDateTimeComponentDayOfWeekNarrow renders a super-short day-of-week
// compontent - not guaranteed to be unique for different days.
func (c Calendar) formatDateTimeComponentDayOfWeekNarrow(dayOfWeek time.Weekday) string {
	switch dayOfWeek {
	case time.Sunday:
		return c.FormatNames.Days.Narrow.Sun
	case time.Monday:
		return c.FormatNames.Days.Narrow.Mon
	case time.Tuesday:
		return c.FormatNames.Days.Narrow.Tue
	case time.Wednesday:
		return c.FormatNames.Days.Narrow.Wed
	case time.Thursday:
		return c.FormatNames.Days.Narrow.Thu
	case time.Friday:
		return c.FormatNames.Days.Narrow.Fri
	case time.Saturday:
		return c.FormatNames.Days.Narrow.Sat
	}

	return ""
}

// formatDateTimeComponentDay renders a day-of-year component.
func (c Calendar) formatDateTimeComponentDay(datetime time.Time, length int) (string, error) {
	day := datetime.Day()

	switch length {
	case datetimeFormatLength1Plus:
		return fmt.Sprintf("%d", day), nil
	case datetimeFormatLength2Plus:
		if day < 10 {
			return fmt.Sprintf("0%d", day), nil
		}
		return fmt.Sprintf("%d", day), nil
	}

	return "", fmt.Errorf("unsupported day-of-year: %d", length)
}

// formatDateTimeComponentHour12 renders an hour-component using a 12-hour
// clock.
func (c Calendar) formatDateTimeComponentHour12(datetime time.Time, length int) (string, error) {
	hour := datetime.Hour()
	if hour > 12 {
		hour = hour - 12
	}

	switch length {
	case datetimeFormatLength1Plus:
		return fmt.Sprintf("%d", hour), nil
	case datetimeFormatLength2Plus:
		if hour < 10 {
			return fmt.Sprintf("0%d", hour), nil
		}
		return fmt.Sprintf("%d", hour), nil
	}

	return "", fmt.Errorf("unsupported hour-12: %d", length)
}

// formatDateTimeComponentHour24 renders an hour-component using a 24-hour
// clock.
func (c Calendar) formatDateTimeComponentHour24(datetime time.Time, length int) (string, error) {
	hour := datetime.Hour()

	switch length {
	case datetimeFormatLength1Plus:
		return fmt.Sprintf("%d", hour), nil
	case datetimeFormatLength2Plus:
		if hour < 10 {
			return fmt.Sprintf("0%d", hour), nil
		}
		return fmt.Sprintf("%d", hour), nil
	}

	return "", fmt.Errorf("unsupported hour-24: %d", length)
}

// formatDateTimeComponentMinute renders a minute component.
func (c Calendar) formatDateTimeComponentMinute(datetime time.Time, length int) (string, error) {
	minute := datetime.Minute()

	switch length {
	case datetimeFormatLength1Plus:
		return fmt.Sprintf("%d", minute), nil
	case datetimeFormatLength2Plus:
		if minute < 10 {
			return fmt.Sprintf("0%d", minute), nil
		}
		return fmt.Sprintf("%d", minute), nil
	}

	return "", fmt.Errorf("unsupported minute: %d", length)
}

// formatDateTimeComponentSecond renders a second component
func (c Calendar) formatDateTimeComponentSecond(datetime time.Time, length int) (string, error) {
	second := datetime.Second()

	switch length {
	case datetimeFormatLength1Plus:
		return fmt.Sprintf("%d", second), nil
	case datetimeFormatLength2Plus:
		if second < 10 {
			return fmt.Sprintf("0%d", second), nil
		}
		return fmt.Sprintf("%d", second), nil
	}

	return "", fmt.Errorf("unsupported second: %d", length)
}

// formatDateTimeComponentPeriod renders a period component (AM/PM).
func (c Calendar) formatDateTimeComponentPeriod(datetime time.Time, length int) (string, error) {
	hour := datetime.Hour()

	switch length {
	case datetimeFormatLength1Plus:
		return c.formatDateTimeComponentPeriodWide(hour), nil
	case datetimeFormatLengthAbbreviated:
		return c.formatDateTimeComponentPeriodAbbreviated(hour), nil
	case datetimeFormatLengthWide:
		return c.formatDateTimeComponentPeriodWide(hour), nil
	case datetimeFormatLengthNarrow:
		return c.formatDateTimeComponentPeriodNarrow(hour), nil
	}

	return "", fmt.Errorf("unsupported day-period: %d", length)
}

// formatDateTimeComponentPeriodAbbreviated renders an abbreviated period
// component.
func (c Calendar) formatDateTimeComponentPeriodAbbreviated(hour int) string {
	if hour < 12 {
		return c.FormatNames.Periods.Abbreviated.AM
	}

	return c.FormatNames.Periods.Abbreviated.PM
}

// formatDateTimeComponentPeriodWide renders a full period component.
func (c Calendar) formatDateTimeComponentPeriodWide(hour int) string {
	if hour < 12 {
		return c.FormatNames.Periods.Wide.AM
	}

	return c.FormatNames.Periods.Wide.PM
}

// formatDateTimeComponentPeriodNarrow renders a super-short period component.
func (c Calendar) formatDateTimeComponentPeriodNarrow(hour int) string {
	if hour < 12 {
		return c.FormatNames.Periods.Narrow.AM
	}

	return c.FormatNames.Periods.Narrow.PM
}

// formatDateTimeComponentQuarter renders a calendar quarter component - this
// is calendar quarters and not fiscal quarters.
//  - Q1: Jan-Mar
//  - Q2: Apr-Jun
//  - Q3: Jul-Sep
//  - Q4: Oct-Dec
// TODO: not yet implemented
func (c Calendar) formatDateTimeComponentQuarter(datetime time.Time, length int) (string, error) {
	return "", nil
}

// formatDateTimeComponentTimeZone renders a time zone component.
// TODO: this has not yet been implemented
func (c Calendar) formatDateTimeComponentTimeZone(datetime time.Time, length int) (string, error) {
	return "", nil
}

// parseDateTimeFormat takes a format pattern string and returns a sequence of
// components.
func (c Calendar) parseDateTimeFormat(pattern string) ([]*datetimePatternComponent, error) {
	// every thing between single quotes should become a literal
	// all non a-z, A-Z characters become a literal
	// everything else, repeat character sequences become a component
	format := []*datetimePatternComponent{}
	for i := 0; i < len(pattern); {
		char := pattern[i : i+1]

		skip := false
		// for units we don't support yet, just skip over them
		for _, r := range datetimeFormatUnitCutset {
			if char == string(r) {
				skip = true
				break
			}
		}

		if skip {
			i++
			continue
		}

		if char == string(datetimeFormatLiteral) {
			// find the next single quote
			// create a literal out of everything between the quotes
			// and set i to the position after the second quote

			if i == len(pattern)-1 {
				return []*datetimePatternComponent{}, errors.New("malformed datetime format")
			}

			nextQuote := strings.Index(pattern[i+1:], string(datetimeFormatLiteral))
			if nextQuote == -1 {
				return []*datetimePatternComponent{}, errors.New("malformed datetime format")
			}

			component := &datetimePatternComponent{
				pattern:       pattern[i+1 : nextQuote+i+1],
				componentType: datetimePatternComponentLiteral,
			}

			format = append(format, component)
			i = nextQuote + i + 2
			continue

		}
		if (char >= "a" && char <= "z") || (char >= "A" && char <= "Z") {
			// this represents a format unit
			// find the entire sequence of the same character
			endChar := lastSequenceIndex(pattern[i:]) + i

			component := &datetimePatternComponent{
				pattern:       pattern[i : endChar+1],
				componentType: datetimePatternComponentUnit,
			}

			format = append(format, component)
			i = endChar + 1
			continue

		}
		if char == string(datetimeFormatTimeSeparator) {
			component := &datetimePatternComponent{
				// pattern:       c.TimeSeparator,
				pattern:       string(datetimeFormatTimeSeparator),
				componentType: datetimePatternComponentLiteral,
			}
			format = append(format, component)
			i++
			continue

		}

		component := &datetimePatternComponent{
			pattern:       char,
			componentType: datetimePatternComponentLiteral,
		}

		format = append(format, component)
		i++
		continue

	}

	return format, nil
}

// getDateTimePattern combines a date pattern and a time pattern into a datetime
// pattern. The datetimePattern argument includes a {0} placeholder for the time
// pattern, and a {1} placeholder for the date component.
func getDateTimePattern(datetimePattern, datePattern, timePattern string) string {
	return strings.Replace(strings.Replace(datetimePattern, "{1}", datePattern, 1), "{0}", timePattern, 1)
}

// lastSequenceIndex looks at the first character in a string and returns the
// last digits of the first sequence of that character. For example:
//  - ABC: 0
//  - AAB: 1
//  - ABA: 0
//  - AAA: 2
func lastSequenceIndex(str string) int {
	if len(str) == 0 {
		return -1
	}

	if len(str) == 1 {
		return 0
	}

	sequenceChar := str[0:1]
	lastPos := 0
	for i := 1; i < len(str); i++ {
		if str[i:i+1] != sequenceChar {
			break
		}

		lastPos = i
	}

	return lastPos
}
