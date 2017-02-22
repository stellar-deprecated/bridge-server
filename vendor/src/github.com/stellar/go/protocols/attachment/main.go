package attachment

// Attachment represents core object in Stellar attachment convention
type Attachment struct {
	Nonce       string `json:"nonce"`
	Transaction `json:"transaction"`
	Operations  []Operation `json:"operations"`
}

// Transaction represents transaction field in Stellar attachment
type Transaction struct {
	SenderInfo map[string]string `json:"sender_info"`
	Route      string            `json:"route"`
	Note       string            `json:"note"`
	Extra      string            `json:"extra"`
}

// Operation represents a single operation object in Stellar attachment
type Operation struct {
	// Overriddes Transaction field for this operation
	SenderInfo map[string]string `json:"sender_info"`
	// Overriddes Transaction field for this operation
	Route string `json:"route"`
	// Overriddes Transaction field for this operation
	Note string `json:"note"`
}

// SenderInfo is a helper structure with standardized fields that contains
// information about the sender. Use Map() method to transform it to
// map[string]string used in Transaction and Operation structs.
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
