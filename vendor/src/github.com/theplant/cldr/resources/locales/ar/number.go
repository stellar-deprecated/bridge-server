package ar

import "github.com/theplant/cldr"

var (
	symbols = cldr.Symbols{Decimal: "٫", Group: "٬", Negative: "\u200f-", Percent: "٪", PerMille: "؉"}
	formats = cldr.NumberFormats{Decimal: "#,##0.###", Currency: "¤\u00a0#,##0.00", Percent: "#,##0%"}
)
