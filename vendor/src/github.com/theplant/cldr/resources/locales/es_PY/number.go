package es_PY

import "github.com/theplant/cldr"

var (
	symbols = cldr.Symbols{Decimal: ",", Group: ".", Negative: "", Percent: "", PerMille: ""}
	formats = cldr.NumberFormats{Decimal: "", Currency: "¤\u00a0#,##0.00;¤\u00a0-#,##0.00", Percent: ""}
)
