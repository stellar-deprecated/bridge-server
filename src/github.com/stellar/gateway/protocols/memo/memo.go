package memo

import (
	"encoding/json"
)

type Memo struct {
	Transaction `json:"transaction"`
	Operations  []Operation `json:"operations"`
}

type Transaction struct {
	SenderInfo string `json:"sender_info"`
	Route      string `json:"route"`
	Extra      string `json:"extra"`
	Note       string `json:"note"`
}

type Operation struct {
	// Overriddes Transaction field for this operation
	SenderInfo string `json:"sender_info"`
	// Overriddes Transaction field for this operation
	Route string `json:"route"`
	// Overriddes Transaction field for this operation
	Note string `json:"note"`
}

func (memo *Memo) Marshal() []byte {
	json, _ := json.MarshalIndent(memo, "", "  ")
	return json
}
