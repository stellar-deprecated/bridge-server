package cldr_test

import (
	"time"

	_ "github.com/theplant/cldr/resources/locales"
	. "gopkg.in/check.v1"

	"github.com/theplant/cldr"
)

var (
	dateTimeString = "Jan 2, 2006 at 3:04:05pm"
	enLocale, _    = cldr.GetLocale("en")
)

type CalendarSuite struct {
	messagesDir       string
	rulesDir          string
	messagesDirStyle1 string
	messagesDirStyle2 string
}

var _ = Suite(&CalendarSuite{})

func (s *CalendarSuite) TestFormatDateTime(c *C) {
	datetime, _ := time.Parse(dateTimeString, dateTimeString)

	// test the public method
	dFull, err := enLocale.Calendar.FmtDateFull(datetime)
	c.Check(err, IsNil)

	dLong, err := enLocale.Calendar.FmtDateLong(datetime)
	c.Check(err, IsNil)

	dMedium, err := enLocale.Calendar.FmtDateMedium(datetime)
	c.Check(err, IsNil)

	dShort, err := enLocale.Calendar.FmtDateShort(datetime)
	c.Check(err, IsNil)

	tFull, err := enLocale.Calendar.FmtTimeFull(datetime)
	c.Check(err, IsNil)

	tLong, err := enLocale.Calendar.FmtTimeLong(datetime)
	c.Check(err, IsNil)

	tMedium, err := enLocale.Calendar.FmtTimeMedium(datetime)
	c.Check(err, IsNil)

	tShort, err := enLocale.Calendar.FmtTimeShort(datetime)
	c.Check(err, IsNil)

	dtFull, err := enLocale.Calendar.FmtDateTimeFull(datetime)
	c.Check(err, IsNil)

	dtLong, err := enLocale.Calendar.FmtDateTimeLong(datetime)
	c.Check(err, IsNil)

	dtMedium, err := enLocale.Calendar.FmtDateTimeMedium(datetime)
	c.Check(err, IsNil)

	dtShort, err := enLocale.Calendar.FmtDateTimeShort(datetime)
	c.Check(err, IsNil)

	c.Check(dFull, Equals, "Monday, January 2, 2006")
	c.Check(dLong, Equals, "January 2, 2006")
	c.Check(dMedium, Equals, "Jan 2, 2006")
	c.Check(dShort, Equals, "1/2/06")
	c.Check(tFull, Equals, "3:04:05 PM")
	c.Check(tLong, Equals, "3:04:05 PM")
	c.Check(tMedium, Equals, "3:04:05 PM")
	c.Check(tShort, Equals, "3:04 PM")
	c.Check(dtFull, Equals, "Monday, January 2, 2006 at 3:04:05 PM")
	c.Check(dtLong, Equals, "January 2, 2006 at 3:04:05 PM")
	c.Check(dtMedium, Equals, "Jan 2, 2006, 3:04:05 PM")
	c.Check(dtShort, Equals, "1/2/06, 3:04 PM")

	r, err := enLocale.Calendar.Format(datetime, "MMMM d yy")
	c.Check(err, Equals, nil)
	c.Check(r, Equals, "January 2 06")

	// // should check other locales too - maybe all of them?
	// locales := map[string]string{
	// 	// "aa":  "02/01/06",
	// 	"af":  "2006-01-02",
	// 	"agq": "2/1/2006",
	// 	"ak":  "06/01/02",
	// 	"am":  "02/01/2006",
	// 	"ar":  "2\u200f/1\u200f/2006",
	// 	"as":  "2-1-2006",
	// 	"asa": "02/01/2006",
	// 	"ast": "2/1/06",
	// 	// "az":      "06/01/02",
	// 	"bas": "2/1/2006",
	// 	"be":  "2.1.06",
	// 	"bem": "02/01/2006",
	// 	"bez": "02/01/2006",
	// 	// "bg":    "02.01.06",
	// 	"bm": "2/1/2006",
	// 	"bn": "2/1/06",
	// 	// "bo":    "2006-01-02",
	// 	// "br":    "2006-01-02",
	// 	"brx": "1/2/06",
	// 	"bs":  "02.01.06.",
	// 	// "byn": "02/01/06",
	// 	// "ca":    "02/01/06",
	// 	"cgg": "02/01/2006",
	// 	"chr": "1/2/06",
	// 	"cs":  "02.01.06",
	// 	// "cy":    "02/01/2006",
	// 	// "da":    "02/01/06",
	// 	"dav":   "02/01/2006",
	// 	"de":    "02.01.06",
	// 	"dje":   "2/1/2006",
	// 	"dua":   "2/1/2006",
	// 	"dyo":   "2/1/2006",
	// 	"dz":    "2006-01-02",
	// 	"ebu":   "02/01/2006",
	// 	"ee":    "1/2/06",
	// 	"el":    "2/1/06",
	// 	"en":    "1/2/06",
	// 	"en_AU": "2/01/06",
	// 	"en_GB": "02/01/2006",
	// 	"eo":    "06-01-02",
	// 	// "es":    "02/01/06",
	// 	"et": "02.01.06",
	// 	// "eu":      "2006-01-02",
	// 	"ewo":   "2/1/2006",
	// 	"fa":    "2006/1/2",
	// 	"ff":    "2/1/2006",
	// 	"fi":    "2.1.2006",
	// 	"fil":   "1/2/06",
	// 	"fo":    "02-01-06",
	// 	"fr":    "02/01/2006",
	// 	"fr_CA": "06-01-02",
	// 	"fur":   "02/01/06",
	// 	"ga":    "02/01/2006",
	// 	"gd":    "02/01/2006",
	// 	"gl":    "02/01/06",
	// 	"gsw":   "02.01.06",
	// 	// "gu":    "2-01-06",
	// 	"guz": "02/01/2006",
	// 	"gv":  "02/01/06",
	// 	"ha":  "2/1/06",
	// 	"haw": "2/1/06",
	// 	// "he":    "02/01/06",
	// 	// "hi":    "2-1-06",
	// 	// "hr":    "2.1.2006.",
	// 	// "ht": "2006-01-02",
	// 	// "hu":    "2006.01.02.",
	// 	"hy": "02.01.06",
	// 	// "ia": "06/01/02",
	// 	"id": "02/01/06",
	// 	"ig": "02/01/2006",
	// 	// "ii":    "2006-01-02",
	// 	"is":  "2.1.2006",
	// 	"it":  "02/01/06",
	// 	"ja":  "2006/01/02",
	// 	"jgo": "2006-01-02",
	// 	"jmc": "02/01/2006",
	// 	"ka":  "02.01.06",
	// 	"kab": "2/1/2006",
	// 	"kam": "02/01/2006",
	// 	"kde": "02/01/2006",
	// 	"kea": "2/1/2006",
	// 	"khq": "2/1/2006",
	// 	"ki":  "02/01/2006",
	// 	// "kk":    "02.01.06",
	// 	"kkj": "02/01 2006",
	// 	"kl":  "2006-01-02",
	// 	"kln": "02/01/2006",
	// 	// "km":    "2/1/2006",
	// 	// "kn":    "2-1-06",
	// 	"ko":  "06. 1. 2.",
	// 	"kok": "2-1-06",
	// 	"ks":  "1/2/06",
	// 	"ksb": "02/01/2006",
	// 	"ksf": "2/1/2006",
	// 	"ksh": "2. 1. 2006",
	// 	"kw":  "02/01/2006",
	// 	"ky":  "02.01.06",
	// 	"lag": "02/01/2006",
	// 	"lg":  "02/01/2006",
	// 	"ln":  "2/1/2006",
	// 	"lo":  "2/1/2006",
	// 	"lt":  "2006-01-02",
	// 	"lu":  "2/1/2006",
	// 	"luo": "02/01/2006",
	// 	"luy": "02/01/2006",
	// 	"lv":  "02.01.06",
	// 	"mas": "02/01/2006",
	// 	"mer": "02/01/2006",
	// 	"mfe": "2/1/2006",
	// 	"mg":  "2/1/2006",
	// 	"mgh": "02/01/2006",
	// 	"mgo": "2006-01-02",
	// 	"mk":  "02.1.06",
	// 	"ml":  "02/01/06",
	// 	"mn":  "2006-01-02",
	// 	// "mr":      "2-1-06",
	// 	"ms":  "2/01/06",
	// 	"mt":  "02/01/2006",
	// 	"mua": "2/1/2006",
	// 	// "my":      "06/01/02",
	// 	"naq": "02/01/2006",
	// 	// "nb":      "02.01.06",
	// 	"nd":    "02/01/2006",
	// 	"ne":    "2006-01-02",
	// 	"nl":    "02-01-06",
	// 	"nl_BE": "2/01/06",
	// 	"nmg":   "2/1/2006",
	// 	// "nn":      "02.01.06",
	// 	"nnh": "02/01/06",
	// 	// "no":  "02.01.06",
	// 	// "nr":  "2006-01-02",
	// 	// "nso": "2006-01-02",
	// 	"nus": "2/01/2006",
	// 	"nyn": "02/01/2006",
	// 	"om":  "02/01/06",
	// 	"or":  "2-1-06",
	// 	"os":  "02.01.06",
	// 	// "pa":      "02/01/2006",
	// 	"pl":    "02.01.2006",
	// 	"ps":    "2006/1/2",
	// 	"pt":    "02/01/06",
	// 	"pt_BR": "02/01/06",
	// 	"rm":    "02-01-06",
	// 	"rn":    "2/1/2006",
	// 	"ro":    "02.01.2006",
	// 	"rof":   "02/01/2006",
	// 	"ru":    "02.01.06",
	// 	"rw":    "06/01/02",
	// 	"rwk":   "02/01/2006",
	// 	"sah":   "06/1/2",
	// 	"saq":   "02/01/2006",
	// 	"sbp":   "02/01/2006",
	// 	// "se":      "2006-01-02",
	// 	"seh": "2/1/2006",
	// 	"ses": "2/1/2006",
	// 	"sg":  "2/1/2006",
	// 	"shi": "2/1/2006",
	// 	// "si":      "2006/01/02",
	// 	// "sk":      "2.1.2006",
	// 	"sl": "2. 01. 06",
	// 	"sn": "02/01/2006",
	// 	"so": "02/01/06",
	// 	// "sq":      "06-01-02",
	// 	"sr": "2.1.06.",
	// 	// "ss":  "2006-01-02",
	// 	// "ssy": "02/01/06",
	// 	// "st":  "2006-01-02",
	// 	"sv": "2006-01-02",
	// 	"sw": "02/01/2006",
	// 	// "swc": "2/1/2006",
	// 	"ta":  "2-1-06",
	// 	"te":  "02-01-06",
	// 	"teo": "02/01/2006",
	// 	// "tg":  "06/01/02",
	// 	"th": "2/1/06",
	// 	"ti": "02/01/06",
	// 	// "tig": "02/01/06",
	// 	// "tn":  "2006-01-02",
	// 	"to": "2/1/06",
	// 	"tr": "2.01.2006",
	// 	// "ts":  "2006-01-02",
	// 	"twq": "2/1/2006",
	// 	"tzm": "02/01/2006",
	// 	"uk":  "02.01.06",
	// 	"ur":  "2/1/06",
	// 	"uz":  "06/01/02",
	// 	"vai": "02/01/2006",
	// 	// "ve":  "2006-01-02",
	// 	"vi": "02/01/2006",
	// 	// "vo":  "2006-01-02",
	// 	"vun": "02/01/2006",
	// 	// "wae":     "2006-01-02",
	// 	// "wal": "02/01/06",
	// 	// "xh":  "2006-01-02",
	// 	"xog": "02/01/2006",
	// 	"yav": "2/1/2006",
	// 	"yo":  "02/01/2006",
	// 	"zh":  "2006/1/2",
	// 	// "zh_HANS": "06-1-2",
	// 	// "zu":      "2006-01-02",
	// }

	// for locale, expected := range locales {
	// 	// t, _ := f.GetTranslator(locale)
	// 	c.Log(locale)
	// 	loc, _ := cldr.GetLocale(locale)
	// 	if !c.Check(loc, NotNil) {
	// 		continue
	// 	}
	// 	d, err := loc.Calendar.FmtDateShort(datetime)

	// 	// if d != expected {
	// 	// 	c.Log(locale)
	// 	// }

	// 	c.Check(err, IsNil)
	// 	c.Check(d, Equals, expected)
	// }

	// // test the private method
	// checkAll := "G y yy yyyy M MM MMM MMMM MMMMM E EE EEE EEEE EEEEE d dd h hh H HH m mm s ss a aaa aaaa aaaaa Q z v 'literal':'literal'   ,   "
	// shouldMatch := "2006 06 2006 1 01 Jan January J Monday Mo Mon Monday M 2 02 3 03 15 15 4 04 5 05 PM PM PM p    literal#literal"
	// separator := tEn.rules.DateTime.TimeSeparator
	// tEn.rules.DateTime.TimeSeparator = "#"
	// patternToCheckEverything, _ := tEn.parseDateTimeFormat(checkAll)
	// tEn.rules.DateTime.TimeSeparator = separator

	// custom, err := tEn.formatDateTime(datetime, patternToCheckEverything)
	// tEn.rules.DateTime.TimeSeparator = separator

	// c.Check(err, IsNil)

	// c.Check(custom, Equals, shouldMatch)
}

// func (s *CalendarSuite) TestFormatDateTimeComponent(c *C) {
// 	datetime, _ := time.Parse(dateTimeString, dateTimeString)

// 	str, err := enLocale.Calendar.formatDateTimeComponent(datetime, "G")
// 	c.Check(err, IsNil)
// 	c.Check(str, Equals, "")

// 	str, err = enLocale.Calendar.formatDateTimeComponent(datetime, "y")
// 	c.Check(err, IsNil)
// 	c.Check(str, Equals, "2006")

// 	str, err = enLocale.Calendar.formatDateTimeComponent(datetime, "yy")
// 	c.Check(err, IsNil)
// 	c.Check(str, Equals, "06")

// 	str, err = enLocale.Calendar.formatDateTimeComponent(datetime, "yyyy")
// 	c.Check(err, IsNil)
// 	c.Check(str, Equals, "2006")

// 	_, err = enLocale.Calendar.formatDateTimeComponent(datetime, "yyyyyy")
// 	c.Check(err, NotNil)

// 	str, err = enLocale.Calendar.formatDateTimeComponent(datetime, "M")
// 	c.Check(err, IsNil)
// 	c.Check(str, Equals, "1")

// 	str, err = enLocale.Calendar.formatDateTimeComponent(datetime, "MM")
// 	c.Check(err, IsNil)
// 	c.Check(str, Equals, "01")

// 	str, err = enLocale.Calendar.formatDateTimeComponent(datetime, "MMM")
// 	c.Check(err, IsNil)
// 	c.Check(str, Equals, "Jan")

// 	str, err = enLocale.Calendar.formatDateTimeComponent(datetime, "MMMM")
// 	c.Check(err, IsNil)
// 	c.Check(str, Equals, "January")

// 	str, err = enLocale.Calendar.formatDateTimeComponent(datetime, "MMMMM")
// 	c.Check(err, IsNil)
// 	c.Check(str, Equals, "J")

// 	_, err = enLocale.Calendar.formatDateTimeComponent(datetime, "MMMMMM")
// 	c.Check(err, NotNil)

// 	str, err = enLocale.Calendar.formatDateTimeComponent(datetime, "E")
// 	c.Check(err, IsNil)
// 	c.Check(str, Equals, "Monday")

// 	str, err = enLocale.Calendar.formatDateTimeComponent(datetime, "EE")
// 	c.Check(err, IsNil)
// 	c.Check(str, Equals, "Mo")

// 	str, err = enLocale.Calendar.formatDateTimeComponent(datetime, "EEE")
// 	c.Check(err, IsNil)
// 	c.Check(str, Equals, "Mon")

// 	str, err = enLocale.Calendar.formatDateTimeComponent(datetime, "EEEE")
// 	c.Check(err, IsNil)
// 	c.Check(str, Equals, "Monday")

// 	str, err = enLocale.Calendar.formatDateTimeComponent(datetime, "EEEEE")
// 	c.Check(err, IsNil)
// 	c.Check(str, Equals, "M")

// 	_, err = enLocale.Calendar.formatDateTimeComponent(datetime, "EEEEEE")
// 	c.Check(err, NotNil)

// 	str, err = enLocale.Calendar.formatDateTimeComponent(datetime, "d")
// 	c.Check(err, IsNil)
// 	c.Check(str, Equals, "2")

// 	str, err = enLocale.Calendar.formatDateTimeComponent(datetime, "dd")
// 	c.Check(err, IsNil)
// 	c.Check(str, Equals, "02")

// 	_, err = enLocale.Calendar.formatDateTimeComponent(datetime, "dddddd")
// 	c.Check(err, NotNil)

// 	str, err = enLocale.Calendar.formatDateTimeComponent(datetime, "h")
// 	c.Check(err, IsNil)
// 	c.Check(str, Equals, "3")

// 	str, err = enLocale.Calendar.formatDateTimeComponent(datetime, "hh")
// 	c.Check(err, IsNil)
// 	c.Check(str, Equals, "03")

// 	_, err = enLocale.Calendar.formatDateTimeComponent(datetime, "hhhhhh")
// 	c.Check(err, NotNil)

// 	str, err = enLocale.Calendar.formatDateTimeComponent(datetime, "H")
// 	c.Check(err, IsNil)
// 	c.Check(str, Equals, "15")

// 	str, err = enLocale.Calendar.formatDateTimeComponent(datetime, "HH")
// 	c.Check(err, IsNil)
// 	c.Check(str, Equals, "15")

// 	_, err = enLocale.Calendar.formatDateTimeComponent(datetime, "HHHHHH")
// 	c.Check(err, NotNil)

// 	str, err = enLocale.Calendar.formatDateTimeComponent(datetime, "m")
// 	c.Check(err, IsNil)
// 	c.Check(str, Equals, "4")

// 	str, err = enLocale.Calendar.formatDateTimeComponent(datetime, "mm")
// 	c.Check(err, IsNil)
// 	c.Check(str, Equals, "04")

// 	_, err = enLocale.Calendar.formatDateTimeComponent(datetime, "mmmmmm")
// 	c.Check(err, NotNil)

// 	str, err = enLocale.Calendar.formatDateTimeComponent(datetime, "s")
// 	c.Check(err, IsNil)
// 	c.Check(str, Equals, "5")

// 	str, err = enLocale.Calendar.formatDateTimeComponent(datetime, "ss")
// 	c.Check(err, IsNil)
// 	c.Check(str, Equals, "05")

// 	_, err = enLocale.Calendar.formatDateTimeComponent(datetime, "ssssss")
// 	c.Check(err, NotNil)

// 	abbr := tEn.rules.DateTime.FormatNames.Periods.Abbreviated.PM
// 	narrow := tEn.rules.DateTime.FormatNames.Periods.Narrow.PM
// 	wide := tEn.rules.DateTime.FormatNames.Periods.Wide.PM
// 	tEn.rules.DateTime.FormatNames.Periods.Abbreviated.PM = "abbr"
// 	tEn.rules.DateTime.FormatNames.Periods.Narrow.PM = "narrow"
// 	tEn.rules.DateTime.FormatNames.Periods.Wide.PM = "wide"
// 	str, err = enLocale.Calendar.formatDateTimeComponent(datetime, "a")
// 	c.Check(err, IsNil)
// 	c.Check(str, Equals, "wide")
// 	str, err = enLocale.Calendar.formatDateTimeComponent(datetime, "aaa")
// 	c.Check(err, IsNil)
// 	c.Check(str, Equals, "abbr")
// 	str, err = enLocale.Calendar.formatDateTimeComponent(datetime, "aaaa")
// 	c.Check(err, IsNil)
// 	c.Check(str, Equals, "wide")
// 	str, err = enLocale.Calendar.formatDateTimeComponent(datetime, "aaaaa")
// 	c.Check(err, IsNil)
// 	c.Check(str, Equals, "narrow")
// 	tEn.rules.DateTime.FormatNames.Periods.Abbreviated.PM = abbr
// 	tEn.rules.DateTime.FormatNames.Periods.Narrow.PM = narrow
// 	tEn.rules.DateTime.FormatNames.Periods.Wide.PM = wide

// 	_, err = enLocale.Calendar.formatDateTimeComponent(datetime, "aaaaaa")
// 	c.Check(err, NotNil)

// 	str, err = enLocale.Calendar.formatDateTimeComponent(datetime, "Q")
// 	c.Check(err, IsNil)
// 	c.Check(str, Equals, "")

// 	str, err = enLocale.Calendar.formatDateTimeComponent(datetime, "z")
// 	c.Check(err, IsNil)
// 	c.Check(str, Equals, "")

// 	str, err = enLocale.Calendar.formatDateTimeComponent(datetime, "v")
// 	c.Check(err, IsNil)
// 	c.Check(str, Equals, "")

// 	_, err = enLocale.Calendar.formatDateTimeComponent(datetime, "%")
// 	c.Check(err, NotNil)
// }

// func (s *CalendarSuite) TestParseDateTimeFormat(c *C) {
// 	f, _ := NewTranslatorFactory(
// 		[]string{"data/rules"},
// 		[]string{"data/messages"},
// 		"en",
// 	)

// 	c.Assert(f, NotNil)

// 	tEn, _ := f.GetTranslator("en")

// 	c.Assert(tEn, NotNil)

// 	checkAll := "G y yy yyyy M MM MMM MMMM MMMMM E EE EEE EEEE EEEEE d dd h hh H HH m mm s ss a aaa aaaa aaaaa Q z v 'literal':'literal'"
// 	patternToCheckEverything, err := tEn.parseDateTimeFormat(checkAll)

// 	c.Check(err, IsNil)
// 	c.Check(patternToCheckEverything, HasLen, 61)
// 	c.Check(patternToCheckEverything[0].pattern, Equals, " ")
// 	c.Check(patternToCheckEverything[0].componentType, Equals, datetimePatternComponentLiteral)
// 	c.Check(patternToCheckEverything[1].pattern, Equals, "y")
// 	c.Check(patternToCheckEverything[1].componentType, Equals, datetimePatternComponentUnit)
// 	c.Check(patternToCheckEverything[2].pattern, Equals, " ")
// 	c.Check(patternToCheckEverything[2].componentType, Equals, datetimePatternComponentLiteral)
// 	c.Check(patternToCheckEverything[3].pattern, Equals, "yy")
// 	c.Check(patternToCheckEverything[3].componentType, Equals, datetimePatternComponentUnit)
// 	c.Check(patternToCheckEverything[4].pattern, Equals, " ")
// 	c.Check(patternToCheckEverything[4].componentType, Equals, datetimePatternComponentLiteral)
// 	c.Check(patternToCheckEverything[5].pattern, Equals, "yyyy")
// 	c.Check(patternToCheckEverything[5].componentType, Equals, datetimePatternComponentUnit)
// 	c.Check(patternToCheckEverything[6].pattern, Equals, " ")
// 	c.Check(patternToCheckEverything[6].componentType, Equals, datetimePatternComponentLiteral)
// 	c.Check(patternToCheckEverything[7].pattern, Equals, "M")
// 	c.Check(patternToCheckEverything[7].componentType, Equals, datetimePatternComponentUnit)
// 	c.Check(patternToCheckEverything[8].pattern, Equals, " ")
// 	c.Check(patternToCheckEverything[8].componentType, Equals, datetimePatternComponentLiteral)
// 	c.Check(patternToCheckEverything[9].pattern, Equals, "MM")
// 	c.Check(patternToCheckEverything[9].componentType, Equals, datetimePatternComponentUnit)
// 	c.Check(patternToCheckEverything[10].pattern, Equals, " ")
// 	c.Check(patternToCheckEverything[10].componentType, Equals, datetimePatternComponentLiteral)
// 	c.Check(patternToCheckEverything[11].pattern, Equals, "MMM")
// 	c.Check(patternToCheckEverything[11].componentType, Equals, datetimePatternComponentUnit)
// 	c.Check(patternToCheckEverything[12].pattern, Equals, " ")
// 	c.Check(patternToCheckEverything[12].componentType, Equals, datetimePatternComponentLiteral)
// 	c.Check(patternToCheckEverything[13].pattern, Equals, "MMMM")
// 	c.Check(patternToCheckEverything[13].componentType, Equals, datetimePatternComponentUnit)
// 	c.Check(patternToCheckEverything[14].pattern, Equals, " ")
// 	c.Check(patternToCheckEverything[14].componentType, Equals, datetimePatternComponentLiteral)
// 	c.Check(patternToCheckEverything[15].pattern, Equals, "MMMMM")
// 	c.Check(patternToCheckEverything[15].componentType, Equals, datetimePatternComponentUnit)
// 	c.Check(patternToCheckEverything[16].pattern, Equals, " ")
// 	c.Check(patternToCheckEverything[16].componentType, Equals, datetimePatternComponentLiteral)
// 	c.Check(patternToCheckEverything[17].pattern, Equals, "E")
// 	c.Check(patternToCheckEverything[17].componentType, Equals, datetimePatternComponentUnit)
// 	c.Check(patternToCheckEverything[18].pattern, Equals, " ")
// 	c.Check(patternToCheckEverything[18].componentType, Equals, datetimePatternComponentLiteral)
// 	c.Check(patternToCheckEverything[19].pattern, Equals, "EE")
// 	c.Check(patternToCheckEverything[19].componentType, Equals, datetimePatternComponentUnit)
// 	c.Check(patternToCheckEverything[20].pattern, Equals, " ")
// 	c.Check(patternToCheckEverything[20].componentType, Equals, datetimePatternComponentLiteral)
// 	c.Check(patternToCheckEverything[21].pattern, Equals, "EEE")
// 	c.Check(patternToCheckEverything[21].componentType, Equals, datetimePatternComponentUnit)
// 	c.Check(patternToCheckEverything[22].pattern, Equals, " ")
// 	c.Check(patternToCheckEverything[22].componentType, Equals, datetimePatternComponentLiteral)
// 	c.Check(patternToCheckEverything[23].pattern, Equals, "EEEE")
// 	c.Check(patternToCheckEverything[23].componentType, Equals, datetimePatternComponentUnit)
// 	c.Check(patternToCheckEverything[24].pattern, Equals, " ")
// 	c.Check(patternToCheckEverything[24].componentType, Equals, datetimePatternComponentLiteral)
// 	c.Check(patternToCheckEverything[25].pattern, Equals, "EEEEE")
// 	c.Check(patternToCheckEverything[25].componentType, Equals, datetimePatternComponentUnit)
// 	c.Check(patternToCheckEverything[26].pattern, Equals, " ")
// 	c.Check(patternToCheckEverything[26].componentType, Equals, datetimePatternComponentLiteral)
// 	c.Check(patternToCheckEverything[27].pattern, Equals, "d")
// 	c.Check(patternToCheckEverything[27].componentType, Equals, datetimePatternComponentUnit)
// 	c.Check(patternToCheckEverything[28].pattern, Equals, " ")
// 	c.Check(patternToCheckEverything[28].componentType, Equals, datetimePatternComponentLiteral)
// 	c.Check(patternToCheckEverything[29].pattern, Equals, "dd")
// 	c.Check(patternToCheckEverything[29].componentType, Equals, datetimePatternComponentUnit)
// 	c.Check(patternToCheckEverything[30].pattern, Equals, " ")
// 	c.Check(patternToCheckEverything[30].componentType, Equals, datetimePatternComponentLiteral)
// 	c.Check(patternToCheckEverything[31].pattern, Equals, "h")
// 	c.Check(patternToCheckEverything[31].componentType, Equals, datetimePatternComponentUnit)
// 	c.Check(patternToCheckEverything[32].pattern, Equals, " ")
// 	c.Check(patternToCheckEverything[32].componentType, Equals, datetimePatternComponentLiteral)
// 	c.Check(patternToCheckEverything[33].pattern, Equals, "hh")
// 	c.Check(patternToCheckEverything[33].componentType, Equals, datetimePatternComponentUnit)
// 	c.Check(patternToCheckEverything[34].pattern, Equals, " ")
// 	c.Check(patternToCheckEverything[34].componentType, Equals, datetimePatternComponentLiteral)
// 	c.Check(patternToCheckEverything[35].pattern, Equals, "H")
// 	c.Check(patternToCheckEverything[35].componentType, Equals, datetimePatternComponentUnit)
// 	c.Check(patternToCheckEverything[36].pattern, Equals, " ")
// 	c.Check(patternToCheckEverything[36].componentType, Equals, datetimePatternComponentLiteral)
// 	c.Check(patternToCheckEverything[37].pattern, Equals, "HH")
// 	c.Check(patternToCheckEverything[37].componentType, Equals, datetimePatternComponentUnit)
// 	c.Check(patternToCheckEverything[38].pattern, Equals, " ")
// 	c.Check(patternToCheckEverything[38].componentType, Equals, datetimePatternComponentLiteral)
// 	c.Check(patternToCheckEverything[39].pattern, Equals, "m")
// 	c.Check(patternToCheckEverything[39].componentType, Equals, datetimePatternComponentUnit)
// 	c.Check(patternToCheckEverything[40].pattern, Equals, " ")
// 	c.Check(patternToCheckEverything[40].componentType, Equals, datetimePatternComponentLiteral)
// 	c.Check(patternToCheckEverything[41].pattern, Equals, "mm")
// 	c.Check(patternToCheckEverything[41].componentType, Equals, datetimePatternComponentUnit)
// 	c.Check(patternToCheckEverything[42].pattern, Equals, " ")
// 	c.Check(patternToCheckEverything[42].componentType, Equals, datetimePatternComponentLiteral)
// 	c.Check(patternToCheckEverything[43].pattern, Equals, "s")
// 	c.Check(patternToCheckEverything[43].componentType, Equals, datetimePatternComponentUnit)
// 	c.Check(patternToCheckEverything[44].pattern, Equals, " ")
// 	c.Check(patternToCheckEverything[44].componentType, Equals, datetimePatternComponentLiteral)
// 	c.Check(patternToCheckEverything[45].pattern, Equals, "ss")
// 	c.Check(patternToCheckEverything[45].componentType, Equals, datetimePatternComponentUnit)
// 	c.Check(patternToCheckEverything[46].pattern, Equals, " ")
// 	c.Check(patternToCheckEverything[46].componentType, Equals, datetimePatternComponentLiteral)
// 	c.Check(patternToCheckEverything[47].pattern, Equals, "a")
// 	c.Check(patternToCheckEverything[47].componentType, Equals, datetimePatternComponentUnit)
// 	c.Check(patternToCheckEverything[48].pattern, Equals, " ")
// 	c.Check(patternToCheckEverything[48].componentType, Equals, datetimePatternComponentLiteral)
// 	c.Check(patternToCheckEverything[49].pattern, Equals, "aaa")
// 	c.Check(patternToCheckEverything[49].componentType, Equals, datetimePatternComponentUnit)
// 	c.Check(patternToCheckEverything[50].pattern, Equals, " ")
// 	c.Check(patternToCheckEverything[50].componentType, Equals, datetimePatternComponentLiteral)
// 	c.Check(patternToCheckEverything[51].pattern, Equals, "aaaa")
// 	c.Check(patternToCheckEverything[51].componentType, Equals, datetimePatternComponentUnit)
// 	c.Check(patternToCheckEverything[52].pattern, Equals, " ")
// 	c.Check(patternToCheckEverything[52].componentType, Equals, datetimePatternComponentLiteral)
// 	c.Check(patternToCheckEverything[53].pattern, Equals, "aaaaa")
// 	c.Check(patternToCheckEverything[53].componentType, Equals, datetimePatternComponentUnit)
// 	c.Check(patternToCheckEverything[54].pattern, Equals, " ")
// 	c.Check(patternToCheckEverything[54].componentType, Equals, datetimePatternComponentLiteral)
// 	c.Check(patternToCheckEverything[55].pattern, Equals, " ")
// 	c.Check(patternToCheckEverything[55].componentType, Equals, datetimePatternComponentLiteral)
// 	c.Check(patternToCheckEverything[56].pattern, Equals, " ")
// 	c.Check(patternToCheckEverything[56].componentType, Equals, datetimePatternComponentLiteral)
// 	c.Check(patternToCheckEverything[57].pattern, Equals, " ")
// 	c.Check(patternToCheckEverything[57].componentType, Equals, datetimePatternComponentLiteral)
// 	c.Check(patternToCheckEverything[58].pattern, Equals, "literal")
// 	c.Check(patternToCheckEverything[58].componentType, Equals, datetimePatternComponentLiteral)
// 	c.Check(patternToCheckEverything[59].pattern, Equals, ":")
// 	c.Check(patternToCheckEverything[59].componentType, Equals, datetimePatternComponentLiteral)
// 	c.Check(patternToCheckEverything[60].pattern, Equals, "literal")
// 	c.Check(patternToCheckEverything[60].componentType, Equals, datetimePatternComponentLiteral)

// 	// check bad-quote errors
// 	_, err = tEn.parseDateTimeFormat("'a")
// 	c.Check(err, NotNil)

// 	_, err = tEn.parseDateTimeFormat("a'")
// 	c.Check(err, NotNil)

// 	_, err = tEn.parseDateTimeFormat("a'a'a'a")
// 	c.Check(err, NotNil)
// }

// func (s *CalendarSuite) TestGetDateTimePattern(c *C) {
// 	datetimePattern := "___-{0}-{1}-___"
// 	datePattern := "date"
// 	timePattern := "time"

// 	p := getDateTimePattern(datetimePattern, datePattern, timePattern)
// 	c.Check(p, Equals, "___-time-date-___")
// }

// func (s *CalendarSuite) TestLastSequenceIndex(c *C) {
// 	str := "a"
// 	c.Check(lastSequenceIndex(str), Equals, 0)

// 	str = "aa"
// 	c.Check(lastSequenceIndex(str), Equals, 1)

// 	str = "aba"
// 	c.Check(lastSequenceIndex(str), Equals, 0)

// 	str = "aab"
// 	c.Check(lastSequenceIndex(str), Equals, 1)
// }
