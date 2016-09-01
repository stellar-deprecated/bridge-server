package horizon

type Problem struct {
	Type     string                 `json:"type"`
	Title    string                 `json:"title"`
	Status   int                    `json:"status"`
	Detail   string                 `json:"detail,omitempty"`
	Instance string                 `json:"instance,omitempty"`
	Extras   map[string]interface{} `json:"extras,omitempty"`
}

type Account struct {
	Links struct {
		Self         Link `json:"self"`
		Transactions Link `json:"transactions"`
		Operations   Link `json:"operations"`
		Payments     Link `json:"payments"`
		Effects      Link `json:"effects"`
		Offers       Link `json:"offers"`
	} `json:"_links"`

	HistoryAccount
	Sequence             string            `json:"sequence"`
	SubentryCount        int32             `json:"subentry_count"`
	InflationDestination string            `json:"inflation_destination,omitempty"`
	HomeDomain           string            `json:"home_domain,omitempty"`
	Thresholds           AccountThresholds `json:"thresholds"`
	Flags                AccountFlags      `json:"flags"`
	Balances             []Balance         `json:"balances"`
	Signers              []Signer          `json:"signers"`
}

func (a Account) GetNativeBalance() string {
	for _, balance := range a.Balances {
		if balance.Asset.Type == "native" {
			return balance.Balance
		}
	}

	return "0"
}

type AccountFlags struct {
	AuthRequired  bool `json:"auth_required"`
	AuthRevocable bool `json:"auth_revocable"`
}

type AccountThresholds struct {
	LowThreshold  byte `json:"low_threshold"`
	MedThreshold  byte `json:"med_threshold"`
	HighThreshold byte `json:"high_threshold"`
}

type Asset struct {
	Type   string `json:"asset_type"`
	Code   string `json:"asset_code,omitempty"`
	Issuer string `json:"asset_issuer,omitempty"`
}

type Balance struct {
	Balance string `json:"balance"`
	Limit   string `json:"limit,omitempty"`
	Asset
}

type HistoryAccount struct {
	ID        string `json:"id"`
	PT        string `json:"paging_token"`
	AccountID string `json:"account_id"`
}

type Link struct {
	Href      string `json:"href"`
	Templated bool   `json:"templated,omitempty"`
}


type TransactionSuccess struct {
	Links struct {
		Transaction Link `json:"transaction"`
	} `json:"_links"`
	Hash   string `json:"hash"`
	Ledger int32  `json:"ledger"`
	Env    string `json:"envelope_xdr"`
	Result string `json:"result_xdr"`
	Meta   string `json:"result_meta_xdr"`
}

type Signer struct {
	PublicKey string `json:"public_key"`
	Weight    int32  `json:"weight"`
}

type PaymentResponse struct {
	ID          string `json:"id"`
	Type        string `json:"type"`
	PagingToken string `json:"paging_token"`

	Links struct {
		Transaction struct {
			Href string `json:"href"`
		} `json:"transaction"`
	} `json:"_links"`

	// payment/path_payment fields
	From        string `json:"from"`
	To          string `json:"to"`
	AssetType   string `json:"asset_type"`
	AssetCode   string `json:"asset_code"`
	AssetIssuer string `json:"asset_issuer"`
	Amount      string `json:"amount"`

	// transaction fields
	Memo struct {
		Type  string `json:"memo_type"`
		Value string `json:"memo"`
	}
}
