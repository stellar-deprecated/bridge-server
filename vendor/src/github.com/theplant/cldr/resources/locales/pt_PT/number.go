package pt_PT

import "github.com/theplant/cldr"

var (
	symbols = cldr.Symbols{Decimal: "", Group: "\u00a0", Negative: "", Percent: "", PerMille: ""}
	formats = cldr.NumberFormats{Decimal: "0 milhão", Currency: "#,##0.00\u00a0¤", Percent: ""}
)
