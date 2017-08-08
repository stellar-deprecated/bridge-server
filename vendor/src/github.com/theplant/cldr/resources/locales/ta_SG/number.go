package ta_SG

import "github.com/theplant/cldr"

var (
	symbols = cldr.Symbols{}
	formats = cldr.NumberFormats{Decimal: "#,##0.###", Currency: "¤\u00a0#,##0.00", Percent: "#,##0%"}
)
