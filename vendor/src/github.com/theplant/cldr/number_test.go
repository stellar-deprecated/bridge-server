package cldr_test

import (
	"testing"

	"github.com/theplant/cldr/resources/locales/en"
	"github.com/theplant/cldr/resources/locales/hi"
	"github.com/theplant/cldr/resources/locales/uz"
	. "gopkg.in/check.v1"
)

// passes control of tests off to go-check
func Test(t *testing.T) { TestingT(t) }

type NumberSuite struct {
	messagesDir       string
	rulesDir          string
	messagesDirStyle1 string
	messagesDirStyle2 string
}

var _ = Suite(&NumberSuite{})

func (s *NumberSuite) TestFmtCurrency(c *C) {
	cur, err := en.Locale.Number.FmtCurrency("USD", 12345.6789)
	c.Check(err, IsNil)
	c.Check(cur, Equals, "$12,345.68")

	cur, err = en.Locale.Number.FmtCurrency("USD", -12345.6789)
	c.Check(err, IsNil)
	c.Check(cur, Equals, "($12,345.68)")

	cur, err = en.Locale.Number.FmtCurrency("WHAT???", 12345.6789)
	// c.Check(err, NotNil)
	c.Check(cur, Equals, "12,345.68")

	// try some really big numbers to make sure weird floaty stuff doesn't
	// happen
	cur, err = en.Locale.Number.FmtCurrency("USD", 12345000000000.6789)
	c.Check(err, IsNil)
	c.Check(cur, Equals, "$12,345,000,000,000.68")

	cur, err = en.Locale.Number.FmtCurrency("USD", -12345000000000.6789)
	c.Check(err, IsNil)
	c.Check(cur, Equals, "($12,345,000,000,000.68)")

	// ignore for now...
	// // Try something that needs a partial fallback
	// cur, err = saq.Locale.Number.FmtCurrency("USD", 12345.6789)
	// c.Check(err, IsNil)
	// c.Check(cur, Equals, "US$12,345.68")

	// // And one more for with some unusual symbols for good measure
	// cur, err = ar.Locale.Number.FmtCurrency("USD", -12345.6789)
	// c.Check(err, IsNil)
	// c.Check(cur, Equals, "US$ 12,345.68-")

	// And one more for with some unusual symbols for good measure
	cur, err = en.Locale.Number.FmtCurrency("USD", 0.0084)
	c.Check(err, IsNil)
	c.Check(cur, Equals, "$0.01")
}

func (s *NumberSuite) TestFmtCurrencyWhole(c *C) {
	cur, err := en.Locale.Number.FmtCurrencyWhole("USD", 12345.6789)
	c.Check(err, IsNil)
	c.Check(cur, Equals, "$12,346")

	cur, err = en.Locale.Number.FmtCurrencyWhole("USD", -12345.6789)
	c.Check(err, IsNil)
	c.Check(cur, Equals, "($12,346)")

	cur, err = en.Locale.Number.FmtCurrencyWhole("WHAT???", 12345.6789)
	// c.Check(err, NotNil)
	c.Check(cur, Equals, "12,346")

	// try some really big numbers to make sure weird floaty stuff doesn't
	// happen
	cur, err = en.Locale.Number.FmtCurrencyWhole("USD", 12345000000000.6789)
	c.Check(err, IsNil)
	c.Check(cur, Equals, "$12,345,000,000,001")

	cur, err = en.Locale.Number.FmtCurrencyWhole("USD", -12345000000000.6789)
	c.Check(err, IsNil)
	c.Check(cur, Equals, "($12,345,000,000,001)")
}

func (s *NumberSuite) TestFmtNumber(c *C) {
	// check basic english
	num := en.Locale.Number.FmtNumber(12345.6789)
	c.Check(num, Equals, "12,345.679")

	num = en.Locale.Number.FmtNumber(-12345.6789)
	c.Check(num, Equals, "-12,345.679")

	num = en.Locale.Number.FmtNumber(123456789)
	c.Check(num, Equals, "123,456,789")

	// check Hindi - different group sizes
	num = hi.Locale.Number.FmtNumber(12345.6789)
	c.Check(num, Equals, "12,345.679")

	num = hi.Locale.Number.FmtNumber(-12345.6789)
	c.Check(num, Equals, "-12,345.679")

	num = hi.Locale.Number.FmtNumber(123456789)
	c.Check(num, Equals, "12,34,56,789")

	// check Uzbek - something with a partial fallback
	num = uz.Locale.Number.FmtNumber(12345.6789)
	c.Check(num, Equals, "12٬345٫679")

	num = uz.Locale.Number.FmtNumber(-12345.6789)
	c.Check(num, Equals, "-12٬345٫679")

	num = uz.Locale.Number.FmtNumber(123456789)
	c.Check(num, Equals, "123٬456٬789")

	// format := &(numberFormat{
	// 	positivePrefix:   "p",
	// 	positiveSuffix:   "P%",
	// 	negativePrefix:   "n",
	// 	negativeSuffix:   "N‰",
	// 	multiplier:       2,
	// 	minDecimalDigits: 2,
	// 	maxDecimalDigits: 5,
	// 	minIntegerDigits: 10,
	// 	groupSizeFinal:   6,
	// 	groupSizeMain:    3,
	// })

	// tEn.rules.Numbers.Symbols.Decimal = ".."
	// tEn.rules.Numbers.Symbols.Group = ",,"
	// tEn.rules.Numbers.Symbols.Negative = "--"
	// tEn.rules.Numbers.Symbols.Percent = "%%"
	// tEn.rules.Numbers.Symbols.Permille = "‰‰"

	// // check numbers with too few integer digits & too many decimal digits
	// num = tEn.formatNumber(format, 1.12341234)
	// c.Check(num, Equals, "p0,,000,,000002..24682P%%")
	// num = tEn.formatNumber(format, -1.12341234)
	// c.Check(num, Equals, "n0,,000,,000002..24682N‰‰")

	// // check numbers with more than enough integer digits  & too few decimal digits
	// num = tEn.formatNumber(format, 1234123412341234)
	// c.Check(num, Equals, "p2,,468,,246,,824,,682468..00P%%")
	// num = tEn.formatNumber(format, -1234123412341234)
	// c.Check(num, Equals, "n2,,468,,246,,824,,682468..00N‰‰")
}

func (s *NumberSuite) TestFmtNumberWhole(c *C) {
	num := en.Locale.Number.FmtNumberWhole(12345.6789)
	c.Check(num, Equals, "12,346")

	num = en.Locale.Number.FmtNumberWhole(-12345.6789)
	c.Check(num, Equals, "-12,346")
}

func (s *NumberSuite) TestFmtPercent(c *C) {
	cur := en.Locale.Number.FmtPercent(0.01234)
	c.Check(cur, Equals, "1%")

	cur = en.Locale.Number.FmtPercent(0.1234)
	c.Check(cur, Equals, "12%")

	cur = en.Locale.Number.FmtPercent(1.234)
	c.Check(cur, Equals, "123%")

	cur = en.Locale.Number.FmtPercent(12.34)
	c.Check(cur, Equals, "1,234%")
}

// func (s *NumberSuite) TestParseFormat(c *C) {
// 	enLocale, _ := cldr.GetLocale("en")
// 	format := enLocale.parseFormat("#0", true)
// 	c.Assert(format, NotNil)
// 	c.Check(format.groupSizeFinal, Equals, 0)
// 	c.Check(format.groupSizeMain, Equals, 0)
// 	c.Check(format.maxDecimalDigits, Equals, 0)
// 	c.Check(format.minDecimalDigits, Equals, 0)
// 	c.Check(format.minIntegerDigits, Equals, 1)
// 	c.Check(format.multiplier, Equals, 1)
// 	c.Check(format.negativePrefix, Equals, enLocale.Number.Symbols.Negative)
// 	c.Check(format.negativeSuffix, Equals, "")
// 	c.Check(format.positivePrefix, Equals, "")
// 	c.Check(format.positiveSuffix, Equals, "")

// 	format = enLocale.parseFormat("#0%", true)
// 	c.Assert(format, NotNil)
// 	c.Check(format.groupSizeFinal, Equals, 0)
// 	c.Check(format.groupSizeMain, Equals, 0)
// 	c.Check(format.maxDecimalDigits, Equals, 0)
// 	c.Check(format.minDecimalDigits, Equals, 0)
// 	c.Check(format.minIntegerDigits, Equals, 1)
// 	c.Check(format.multiplier, Equals, 100)
// 	c.Check(format.negativePrefix, Equals, enLocale.Number.Symbols.Negative)
// 	c.Check(format.negativeSuffix, Equals, "%")
// 	c.Check(format.positivePrefix, Equals, "")
// 	c.Check(format.positiveSuffix, Equals, "%")

// 	format = enLocale.parseFormat("#0‰", true)
// 	c.Assert(format, NotNil)
// 	c.Check(format.groupSizeFinal, Equals, 0)
// 	c.Check(format.groupSizeMain, Equals, 0)
// 	c.Check(format.maxDecimalDigits, Equals, 0)
// 	c.Check(format.minDecimalDigits, Equals, 0)
// 	c.Check(format.minIntegerDigits, Equals, 1)
// 	c.Check(format.multiplier, Equals, 1000)
// 	c.Check(format.negativePrefix, Equals, enLocale.Number.Symbols.Negative)
// 	c.Check(format.negativeSuffix, Equals, "‰")
// 	c.Check(format.positivePrefix, Equals, "")
// 	c.Check(format.positiveSuffix, Equals, "‰")

// 	format = enLocale.parseFormat("#0P;#0N", true)
// 	c.Assert(format, NotNil)
// 	c.Check(format.groupSizeFinal, Equals, 0)
// 	c.Check(format.groupSizeMain, Equals, 0)
// 	c.Check(format.maxDecimalDigits, Equals, 0)
// 	c.Check(format.minDecimalDigits, Equals, 0)
// 	c.Check(format.minIntegerDigits, Equals, 1)
// 	c.Check(format.multiplier, Equals, 1)
// 	c.Check(format.negativePrefix, Equals, "")
// 	c.Check(format.negativeSuffix, Equals, "N")
// 	c.Check(format.positivePrefix, Equals, "")
// 	c.Check(format.positiveSuffix, Equals, "P")

// 	format = enLocale.parseFormat("P#0;N#0", true)
// 	c.Assert(format, NotNil)
// 	c.Check(format.groupSizeFinal, Equals, 0)
// 	c.Check(format.groupSizeMain, Equals, 0)
// 	c.Check(format.maxDecimalDigits, Equals, 0)
// 	c.Check(format.minDecimalDigits, Equals, 0)
// 	c.Check(format.minIntegerDigits, Equals, 1)
// 	c.Check(format.multiplier, Equals, 1)
// 	c.Check(format.negativePrefix, Equals, "N")
// 	c.Check(format.negativeSuffix, Equals, "")
// 	c.Check(format.positivePrefix, Equals, "P")
// 	c.Check(format.positiveSuffix, Equals, "")

// 	format = enLocale.parseFormat("#00000.00000", true)
// 	c.Assert(format, NotNil)
// 	c.Check(format.groupSizeFinal, Equals, 0)
// 	c.Check(format.groupSizeMain, Equals, 0)
// 	c.Check(format.maxDecimalDigits, Equals, 5)
// 	c.Check(format.minDecimalDigits, Equals, 5)
// 	c.Check(format.minIntegerDigits, Equals, 5)
// 	c.Check(format.multiplier, Equals, 1)
// 	c.Check(format.negativePrefix, Equals, enLocale.Number.Symbols.Negative)
// 	c.Check(format.negativeSuffix, Equals, "")
// 	c.Check(format.positivePrefix, Equals, "")
// 	c.Check(format.positiveSuffix, Equals, "")

// 	format = enLocale.parseFormat("#0.#####", true)
// 	c.Assert(format, NotNil)
// 	c.Check(format.groupSizeFinal, Equals, 0)
// 	c.Check(format.groupSizeMain, Equals, 0)
// 	c.Check(format.maxDecimalDigits, Equals, 5)
// 	c.Check(format.minDecimalDigits, Equals, 0)
// 	c.Check(format.minIntegerDigits, Equals, 1)
// 	c.Check(format.multiplier, Equals, 1)
// 	c.Check(format.negativePrefix, Equals, enLocale.Number.Symbols.Negative)
// 	c.Check(format.negativeSuffix, Equals, "")
// 	c.Check(format.positivePrefix, Equals, "")
// 	c.Check(format.positiveSuffix, Equals, "")

// 	format = enLocale.parseFormat("##,#0", true)
// 	c.Assert(format, NotNil)
// 	c.Check(format.groupSizeFinal, Equals, 2)
// 	c.Check(format.groupSizeMain, Equals, 2)
// 	c.Check(format.maxDecimalDigits, Equals, 0)
// 	c.Check(format.minDecimalDigits, Equals, 0)
// 	c.Check(format.minIntegerDigits, Equals, 1)
// 	c.Check(format.multiplier, Equals, 1)
// 	c.Check(format.negativePrefix, Equals, enLocale.Number.Symbols.Negative)
// 	c.Check(format.negativeSuffix, Equals, "")
// 	c.Check(format.positivePrefix, Equals, "")
// 	c.Check(format.positiveSuffix, Equals, "")

// 	format = enLocale.parseFormat("##,###,#0", true)
// 	c.Assert(format, NotNil)
// 	c.Check(format.groupSizeFinal, Equals, 2)
// 	c.Check(format.groupSizeMain, Equals, 3)
// 	c.Check(format.maxDecimalDigits, Equals, 0)
// 	c.Check(format.minDecimalDigits, Equals, 0)
// 	c.Check(format.minIntegerDigits, Equals, 1)
// 	c.Check(format.multiplier, Equals, 1)
// 	c.Check(format.negativePrefix, Equals, enLocale.Number.Symbols.Negative)
// 	c.Check(format.negativeSuffix, Equals, "")
// 	c.Check(format.positivePrefix, Equals, "")
// 	c.Check(format.positiveSuffix, Equals, "")

// 	// test includeDecimalDigits true vs false
// 	enDecimal := "#,##0.###"
// 	formatWith := enLocale.parseFormat(enDecimal, true)
// 	formatWithout := enLocale.parseFormat(enDecimal, false)
// 	c.Assert(formatWith, NotNil)
// 	c.Check(formatWith.maxDecimalDigits, Equals, 3)
// 	c.Check(formatWith.minDecimalDigits, Equals, 0)
// 	c.Assert(formatWithout, NotNil)
// 	c.Check(formatWithout.maxDecimalDigits, Equals, 0)
// 	c.Check(formatWithout.minDecimalDigits, Equals, 0)

// 	enCurrency := "¤#,##0.00;(¤#,##0.00)"
// 	formatWith = enLocale.parseFormat(enCurrency, true)
// 	formatWithout = enLocale.parseFormat(enCurrency, false)
// 	c.Assert(formatWith, NotNil)
// 	c.Check(formatWith.maxDecimalDigits, Equals, 2)
// 	c.Check(formatWith.minDecimalDigits, Equals, 2)
// 	c.Assert(formatWithout, NotNil)
// 	c.Check(formatWithout.maxDecimalDigits, Equals, 0)
// 	c.Check(formatWithout.minDecimalDigits, Equals, 0)
// }

// func (s *NumberSuite) TestChunkString(c *C) {
// 	str := ""
// 	size := 0
// 	chunks := chunkString(str, size)
// 	c.Check(chunks, HasLen, 0)

// 	str = ""
// 	size = 1
// 	chunks = chunkString(str, size)
// 	c.Check(chunks, HasLen, 0)

// 	str = "What noise annoys a noisy oyster?"
// 	size = 0
// 	chunks = chunkString(str, size)
// 	c.Assert(chunks, HasLen, 1)
// 	c.Check(chunks[0], Equals, str)

// 	str = "What noise annoys a noisy oyster?"
// 	size = 1
// 	chunks = chunkString(str, size)
// 	c.Assert(chunks, HasLen, 33)
// 	c.Check(chunks[0], Equals, "W")
// 	c.Check(chunks[1], Equals, "h")
// 	c.Check(chunks[2], Equals, "a")

// 	str = "What noise annoys a noisy oyster?"
// 	size = 2
// 	chunks = chunkString(str, size)
// 	c.Assert(chunks, HasLen, 17)
// 	c.Check(chunks[0], Equals, "W")
// 	c.Check(chunks[1], Equals, "ha")
// 	c.Check(chunks[2], Equals, "t ")

// 	str = "What noise annoys a noisy oyster?"
// 	size = 3
// 	chunks = chunkString(str, size)
// 	c.Assert(chunks, HasLen, 11)
// 	c.Check(chunks[0], Equals, "Wha")
// 	c.Check(chunks[1], Equals, "t n")
// 	c.Check(chunks[2], Equals, "ois")

// 	str = "What noise annoys a noisy oyster?"
// 	size = 33
// 	chunks = chunkString(str, size)
// 	c.Assert(chunks, HasLen, 1)
// 	c.Check(chunks[0], Equals, str)

// 	str = "What noise annoys a noisy oyster?"
// 	size = 133
// 	chunks = chunkString(str, size)
// 	c.Assert(chunks, HasLen, 1)
// 	c.Check(chunks[0], Equals, str)
// }

// func (s *NumberSuite) TestNumberRound(c *C) {
// 	num := float64(0)
// 	dec := 0
// 	rounded := numberRound(num, dec)
// 	c.Check(rounded, Equals, "0")

// 	num = float64(1)
// 	dec = 1
// 	rounded = numberRound(num, dec)
// 	c.Check(rounded, Equals, "1")

// 	// test round down
// 	num = 1.2
// 	dec = 0
// 	rounded = numberRound(num, dec)
// 	c.Check(rounded, Equals, "1")

// 	// test round up
// 	num = 1.6
// 	dec = 0
// 	rounded = numberRound(num, dec)
// 	c.Check(rounded, Equals, "2")

// 	num = 1.51
// 	dec = 0
// 	rounded = numberRound(num, dec)
// 	c.Check(rounded, Equals, "2")

// 	// test round to even
// 	num = 1.5
// 	dec = 0
// 	rounded = numberRound(num, dec)
// 	c.Check(rounded, Equals, "2")

// 	num = 2.5
// 	dec = 0
// 	rounded = numberRound(num, dec)
// 	c.Check(rounded, Equals, "2")

// 	// a few more
// 	num = 1.99
// 	dec = 1
// 	rounded = numberRound(num, dec)
// 	c.Check(rounded, Equals, "2")

// 	num = 1.23456789
// 	dec = 2
// 	rounded = numberRound(num, dec)
// 	c.Check(rounded, Equals, "1.23")

// 	num = 1.23456789
// 	dec = 3
// 	rounded = numberRound(num, dec)
// 	c.Check(rounded, Equals, "1.235")
// }
