package bo

import "github.com/theplant/cldr"

var (
	symbols = cldr.Symbols{Decimal: ".", Group: ",", Negative: "", Percent: "", PerMille: ""}
	formats = cldr.NumberFormats{Decimal: "#,##0.###", Currency: "¤\u00a0#,##0.00", Percent: "#,##0%"}
)
