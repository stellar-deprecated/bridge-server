package cldr

// @xerox
type Locale struct {
	Locale     string
	Number     Number
	Calendar   Calendar
	PluralRule string
}

type Number struct {
	Symbols    Symbols
	Formats    NumberFormats
	Currencies []Currency
}

type Symbols struct {
	Decimal  string
	Group    string
	Negative string
	Percent  string
	PerMille string
}

type NumberFormats struct {
	Decimal  string
	Currency string
	Percent  string
}

type Currency struct {
	Currency    string
	DisplayName string
	Symbol      string
}

type Calendar struct {
	Formats     CalendarFormats
	FormatNames CalendarFormatNames
}

type CalendarFormats struct {
	Date     CalendarDateFormat
	Time     CalendarDateFormat
	DateTime CalendarDateFormat
}

type CalendarDateFormat struct{ Full, Long, Medium, Short string }

type CalendarFormatNames struct {
	Months  CalendarMonthFormatNames
	Days    CalendarDayFormatNames
	Periods CalendarPeriodFormatNames
}

type CalendarMonthFormatNames struct {
	Abbreviated CalendarMonthFormatNameValue
	Narrow      CalendarMonthFormatNameValue
	Short       CalendarMonthFormatNameValue
	Wide        CalendarMonthFormatNameValue
}

type CalendarMonthFormatNameValue struct {
	Jan, Feb, Mar, Apr, May, Jun, Jul, Aug, Sep, Oct, Nov, Dec string
}

type CalendarDayFormatNames struct {
	Abbreviated CalendarDayFormatNameValue
	Narrow      CalendarDayFormatNameValue
	Short       CalendarDayFormatNameValue
	Wide        CalendarDayFormatNameValue
}

type CalendarDayFormatNameValue struct {
	Sun, Mon, Tue, Wed, Thu, Fri, Sat string
}

type CalendarPeriodFormatNames struct {
	Abbreviated CalendarPeriodFormatNameValue
	Narrow      CalendarPeriodFormatNameValue
	Short       CalendarPeriodFormatNameValue
	Wide        CalendarPeriodFormatNameValue
}

type CalendarPeriodFormatNameValue struct {
	AM, PM string
}

var locales = map[string]*Locale{}

// TODO: can override paritally instead of replace it as a whole for existed locales
func RegisterLocale(loc *Locale) {
	old, ok := locales[loc.Locale]
	if ok {
		Copy(loc, old)
	} else {
		locales[loc.Locale] = loc
	}
}

// RegisterLocales registers multiple locales in one go.
func RegisterLocales(locs ...*Locale) {
	for _, loc := range locs {
		RegisterLocale(loc)
	}
}

// GetLocale returns a pointer to an existing locale object and a boolean whether the locale exists.
func GetLocale(locale string) (*Locale, bool) {
	l, ok := locales[locale]
	return l, ok
}

func (n Number) findCurrency(currency string) (c Currency) {
	for _, cur := range n.Currencies {
		if cur.Currency == currency {
			c = cur
			break
		}
	}
	return
}
