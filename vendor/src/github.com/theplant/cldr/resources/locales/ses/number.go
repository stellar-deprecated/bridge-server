package ses

import "github.com/theplant/cldr"

var (
	symbols = cldr.Symbols{Decimal: "", Group: "\u00a0", Negative: "", Percent: "", PerMille: ""}
	formats = cldr.NumberFormats{Decimal: "", Currency: "#,##0.00¤", Percent: ""}
)
