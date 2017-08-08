package he

import "github.com/theplant/cldr"

var (
	symbols = cldr.Symbols{Decimal: ".", Group: ",", Negative: "\u200e-", Percent: "%", PerMille: "‰"}
	formats = cldr.NumberFormats{Decimal: "#,##0.###", Currency: "#,##0.00\u00a0¤", Percent: "#,##0%"}
)
