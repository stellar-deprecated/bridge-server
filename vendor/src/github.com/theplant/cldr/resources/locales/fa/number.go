package fa

import "github.com/theplant/cldr"

var (
	symbols = cldr.Symbols{Decimal: "٫", Group: "٬", Negative: "\u200e−", Percent: "٪", PerMille: "؉"}
	formats = cldr.NumberFormats{Decimal: "#,##0.###", Currency: "\u200e¤#,##0.00", Percent: "#,##0%"}
)
