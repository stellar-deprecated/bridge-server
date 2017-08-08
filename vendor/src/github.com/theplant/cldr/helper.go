package cldr

import (
	"fmt"
	"html/template"
	"time"
)

func (l *Locale) FmtDateFull(tim time.Time) (string, error) { return l.Calendar.FmtDateFull(tim) }
func (l *Locale) FmtDateLong(tim time.Time) (string, error) { return l.Calendar.FmtDateLong(tim) }
func (l *Locale) FmtDateMedium(tim time.Time) (string, error) {
	return l.Calendar.FmtDateMedium(tim)
}
func (l *Locale) FmtDateShort(tim time.Time) (string, error) { return l.Calendar.FmtDateShort(tim) }

func (l *Locale) FmtDateTimeFull(tim time.Time) (string, error) {
	return l.Calendar.FmtDateTimeFull(tim)
}
func (l *Locale) FmtDateTimeLong(tim time.Time) (string, error) {
	return l.Calendar.FmtDateTimeLong(tim)
}
func (l *Locale) FmtDateTimeMedium(tim time.Time) (string, error) {
	return l.Calendar.FmtDateTimeMedium(tim)
}
func (l *Locale) FmtDateTimeShort(tim time.Time) (string, error) {
	return l.Calendar.FmtDateTimeShort(tim)
}
func (l *Locale) FmtTimeFull(tim time.Time) (string, error) { return l.Calendar.FmtTimeFull(tim) }
func (l *Locale) FmtTimeLong(tim time.Time) (string, error) { return l.Calendar.FmtTimeLong(tim) }
func (l *Locale) FmtTimeMedium(tim time.Time) (string, error) {
	return l.Calendar.FmtTimeMedium(tim)
}
func (l *Locale) FmtTimeShort(tim time.Time) (string, error) { return l.Calendar.FmtTimeShort(tim) }

func (l *Locale) FmtCurrency(currency string, number interface{}) (formatted string, err error) {
	return l.Number.FmtCurrency(currency, toFloat64(number))
}
func (l *Locale) FmtCurrencyWhole(currency string, number interface{}) (formatted string, err error) {
	return l.Number.FmtCurrencyWhole(currency, toFloat64(number))
}
func (l *Locale) FmtNumber(number interface{}) string {
	return l.Number.FmtNumber(toFloat64(number))
}
func (l *Locale) FmtNumberWhole(number interface{}) string {
	return l.Number.FmtNumberWhole(toFloat64(number))
}
func (l *Locale) FmtPercent(number interface{}) string {
	return l.Number.FmtPercent(toFloat64(number))
}

func (l Locale) FuncMap() template.FuncMap {
	return template.FuncMap{
		"locale":               func() string { return l.Locale },
		"fmt_date_full":        l.FmtDateFull,
		"fmt_date_long":        l.FmtDateLong,
		"fmt_date_medium":      l.FmtDateMedium,
		"fmt_date_short":       l.FmtDateShort,
		"fmt_date_time_full":   l.FmtDateTimeFull,
		"fmt_date_time_long":   l.FmtDateTimeLong,
		"fmt_date_time_medium": l.FmtDateTimeMedium,
		"fmt_date_time_short":  l.FmtDateTimeShort,
		"fmt_time_full":        l.FmtTimeFull,
		"fmt_time_long":        l.FmtTimeLong,
		"fmt_time_medium":      l.FmtTimeMedium,
		"fmt_time_short":       l.FmtTimeShort,
		"fmt_currency":         l.FmtCurrency,
		"fmt_currency_whole":   l.FmtCurrencyWhole,
		"fmt_number":           l.FmtNumber,
		"fmt_number_whole":     l.FmtNumberWhole,
		"fmt_percent":          l.FmtPercent,
	}
}

var FuncMap = template.FuncMap{
	"t":                    Parse,
	"locale":               func() string { return "" },
	"fmt_date_full":        FmtDateFull,
	"fmt_date_long":        FmtDateLong,
	"fmt_date_medium":      FmtDateMedium,
	"fmt_date_short":       FmtDateShort,
	"fmt_date_time_full":   FmtDateTimeFull,
	"fmt_date_time_long":   FmtDateTimeLong,
	"fmt_date_time_medium": FmtDateTimeMedium,
	"fmt_date_time_short":  FmtDateTimeShort,
	"fmt_time_full":        FmtTimeFull,
	"fmt_time_long":        FmtTimeLong,
	"fmt_time_medium":      FmtTimeMedium,
	"fmt_time_short":       FmtTimeShort,
	"fmt_currency":         FmtCurrency,
	"fmt_currency_whole":   FmtCurrencyWhole,
	"fmt_number":           FmtNumber,
	"fmt_number_whole":     FmtNumberWhole,
	"fmt_percent":          FmtPercent,
}

func T(locale, key string, args ...interface{}) (formatted string, err error) {
	_, ok := GetLocale(locale)
	if !ok {
		err = fmt.Errorf("locale %q not found", locale)
		return
	}
	formatted = "err: not implemented"
	return
}

func FmtDateFull(locale string, tim time.Time) (formatted string, err error) {
	l, ok := GetLocale(locale)
	if !ok {
		err = fmt.Errorf("locale %q not found", locale)
		return
	}
	return l.Calendar.FmtDateFull(tim)
}

func FmtDateLong(locale string, tim time.Time) (formatted string, err error) {
	l, ok := GetLocale(locale)
	if !ok {
		err = fmt.Errorf("locale %q not found", locale)
		return
	}
	return l.Calendar.FmtDateLong(tim)
}

func FmtDateMedium(locale string, tim time.Time) (formatted string, err error) {

	l, ok := GetLocale(locale)
	if !ok {
		err = fmt.Errorf("locale %q not found", locale)
		return
	}
	return l.Calendar.FmtDateMedium(tim)
}

func FmtDateShort(locale string, tim time.Time) (formatted string, err error) {
	l, ok := GetLocale(locale)
	if !ok {
		err = fmt.Errorf("locale %q not found", locale)
		return
	}
	return l.Calendar.FmtDateShort(tim)
}

func FmtDateTimeFull(locale string, tim time.Time) (formatted string, err error) {

	l, ok := GetLocale(locale)
	if !ok {
		err = fmt.Errorf("locale %q not found", locale)
		return
	}
	return l.Calendar.FmtDateTimeFull(tim)
}

func FmtDateTimeLong(locale string, tim time.Time) (formatted string, err error) {

	l, ok := GetLocale(locale)
	if !ok {
		err = fmt.Errorf("locale %q not found", locale)
		return
	}
	return l.Calendar.FmtDateTimeLong(tim)
}

func FmtDateTimeMedium(locale string, tim time.Time) (formatted string, err error) {

	l, ok := GetLocale(locale)
	if !ok {
		err = fmt.Errorf("locale %q not found", locale)
		return
	}
	return l.Calendar.FmtDateTimeMedium(tim)
}

func FmtDateTimeShort(locale string, tim time.Time) (formatted string, err error) {

	l, ok := GetLocale(locale)
	if !ok {
		err = fmt.Errorf("locale %q not found", locale)
		return
	}
	return l.Calendar.FmtDateTimeShort(tim)
}

func FmtTimeFull(locale string, tim time.Time) (formatted string, err error) {
	l, ok := GetLocale(locale)
	if !ok {
		err = fmt.Errorf("locale %q not found", locale)
		return
	}
	return l.Calendar.FmtTimeFull(tim)
}

func FmtTimeLong(locale string, tim time.Time) (formatted string, err error) {
	l, ok := GetLocale(locale)
	if !ok {
		err = fmt.Errorf("locale %q not found", locale)
		return
	}
	return l.Calendar.FmtTimeLong(tim)
}

func FmtTimeMedium(locale string, tim time.Time) (formatted string, err error) {

	l, ok := GetLocale(locale)
	if !ok {
		err = fmt.Errorf("locale %q not found", locale)
		return
	}
	return l.Calendar.FmtTimeMedium(tim)
}

func FmtTimeShort(locale string, tim time.Time) (formatted string, err error) {
	l, ok := GetLocale(locale)
	if !ok {
		err = fmt.Errorf("locale %q not found", locale)
		return
	}
	return l.Calendar.FmtTimeShort(tim)
}

func FmtCurrency(locale string, currency string, number interface{}) (formatted string, err error) {
	l, ok := GetLocale(locale)
	if !ok {
		err = fmt.Errorf("locale %q not found", locale)
		return
	}
	return l.Number.FmtCurrency(currency, toFloat64(number))
}

func FmtCurrencyWhole(locale string, currency string, number interface{}) (formatted string, err error) {
	l, ok := GetLocale(locale)
	if !ok {
		err = fmt.Errorf("locale %q not found", locale)
		return
	}
	return l.Number.FmtCurrencyWhole(currency, toFloat64(number))
}

func FmtNumber(locale string, number interface{}) (formatted string, err error) {
	l, ok := GetLocale(locale)
	if !ok {
		err = fmt.Errorf("locale %q not found", locale)
		return
	}
	formatted = l.Number.FmtNumber(toFloat64(number))
	return
}

func FmtNumberWhole(locale string, number interface{}) (formatted string, err error) {
	l, ok := GetLocale(locale)
	if !ok {
		err = fmt.Errorf("locale %q not found", locale)
		return
	}
	formatted = l.Number.FmtNumberWhole(toFloat64(number))
	return
}

func FmtPercent(locale string, number interface{}) (formatted string, err error) {
	l, ok := GetLocale(locale)
	if !ok {
		err = fmt.Errorf("locale %q not found", locale)
		return
	}
	formatted = l.Number.FmtPercent(toFloat64(number))
	return
}
