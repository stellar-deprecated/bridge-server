package memo

import (
	"encoding/json"
)

// Memo represents memo in Stellar memo convention
type Memo struct {
	Transaction `json:"transaction"`
	Operations  []Operation `json:"operations"`
}

// Transaction represents transaction field in Stellar memo
type Transaction struct {
	SenderInfo string `json:"sender_info"`
	Route      string `json:"route"`
	Extra      string `json:"extra"`
	Note       string `json:"note"`
}

// Operation represents a single operation object in Stellar memo
type Operation struct {
	// Overriddes Transaction field for this operation
	SenderInfo string `json:"sender_info"`
	// Overriddes Transaction field for this operation
	Route string `json:"route"`
	// Overriddes Transaction field for this operation
	Note string `json:"note"`
}

// Marshal marshals Memo
func (memo *Memo) Marshal() []byte {
	json, _ := json.MarshalIndent(memo, "", "  ")
	return json
}
