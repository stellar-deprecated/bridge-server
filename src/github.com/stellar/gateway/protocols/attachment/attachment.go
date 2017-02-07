package attachment

import (
	"encoding/json"
)

// Attachment represents core object in Stellar attachment convention
type Attachment struct {
	Transaction `json:"transaction"`
	Operations  []Operation `json:"operations"`
}

// Transaction represents transaction field in Stellar attachment
type Transaction struct {
	Nonce      string     `json:"nonce"`
	SenderInfo SenderInfo `json:"sender_info"`
	Route      string     `json:"route"`
	Note       string     `json:"note"`
	Extra      string     `json:"extra"`
}

// Operation represents a single operation object in Stellar attachment
type Operation struct {
	// Overriddes Transaction field for this operation
	SenderInfo SenderInfo `json:"sender_info"`
	// Overriddes Transaction field for this operation
	Route string `json:"route"`
	// Overriddes Transaction field for this operation
	Note string `json:"note"`
}

// SenderInfo contain information about sender
type SenderInfo struct {
	FirstName   string `json:"first_name,omitempty"`
	MiddleName  string `json:"middle_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	Address     string `json:"address,omitempty"`
	City        string `json:"city,omitempty"`
	Province    string `json:"province,omitempty"`
	Country     string `json:"country,omitempty"`
	DateOfBirth string `json:"date_of_birth,omitempty"`
	CompanyName string `json:"company_name,omitempty"`
}

// Marshal marshals Attachment
func (attachment *Attachment) Marshal() []byte {
	json, _ := json.MarshalIndent(attachment, "", "  ")
	return json
}

// Marshal marshals SenderInfo
func (senderInfo SenderInfo) Marshal() []byte {
	json, _ := json.MarshalIndent(senderInfo, "", "  ")
	return json
}
