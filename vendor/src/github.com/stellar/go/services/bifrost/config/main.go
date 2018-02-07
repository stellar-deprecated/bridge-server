package config

type Config struct {
	Port                           int             `valid:"required"`
	UsingProxy                     bool            `valid:"optional" toml:"using_proxy"`
	Bitcoin                        *bitcoinConfig  `valid:"optional" toml:"bitcoin"`
	Ethereum                       *ethereumConfig `valid:"optional" toml:"ethereum"`
	AccessControlAllowOriginHeader string          `valid:"optional" toml:"access-control-allow-origin-header"`

	Stellar struct {
		Horizon           string `valid:"required" toml:"horizon"`
		NetworkPassphrase string `valid:"required" toml:"network_passphrase"`
		// TokenAssetCode is asset code of token that will be purchased using BTC or ETH.
		TokenAssetCode string `valid:"required" toml:"token_asset_code"`
		// NeedsAuthorize should be set to true if issuers's authorization required flag is set.
		NeedsAuthorize bool `valid:"optional" toml:"needs_authorize"`
		// IssuerPublicKey is public key of the assets issuer or hot wallet.
		IssuerPublicKey string `valid:"required" toml:"issuer_public_key"`
		// SignerSecretKey is:
		// * Issuer's secret key if only one instance of Bifrost is deployed.
		// * Channel's secret key if more than one instance of Bifrost is deployed.
		// https://www.stellar.org/developers/guides/channels.html
		// Signer's sequence number will be consumed in transaction's sequence number.
		SignerSecretKey string `valid:"required" toml:"signer_secret_key"`
	} `valid:"required" toml:"stellar"`
	Database struct {
		Type string `valid:"matches(^postgres$)"`
		DSN  string `valid:"required"`
	} `valid:"required"`
}

type bitcoinConfig struct {
	MasterPublicKey string `valid:"required" toml:"master_public_key"`
	// Minimum value of transaction accepted by Bifrost in BTC.
	// Everything below will be ignored.
	MinimumValueBtc string `valid:"required" toml:"minimum_value_btc"`
	// Host only
	RpcServer string `valid:"required" toml:"rpc_server"`
	RpcUser   string `valid:"optional" toml:"rpc_user"`
	RpcPass   string `valid:"optional" toml:"rpc_pass"`
	Testnet   bool   `valid:"optional" toml:"testnet"`
}

type ethereumConfig struct {
	NetworkID       string `valid:"required,int" toml:"network_id"`
	MasterPublicKey string `valid:"required" toml:"master_public_key"`
	// Minimum value of transaction accepted by Bifrost in ETH.
	// Everything below will be ignored.
	MinimumValueEth string `valid:"required" toml:"minimum_value_eth"`
	// Host only
	RpcServer string `valid:"required" toml:"rpc_server"`
}
