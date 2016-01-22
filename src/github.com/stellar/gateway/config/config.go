package config

type Config struct {
	Port     int
	Horizon  string
	ApiKey   string `mapstructure:"api_key"`
	Assets   []string
	Database struct {
		Type string
		Url  string
	}
	Accounts
	Hooks
}

type Accounts struct {
	AuthorizingSeed    string `mapstructure:"authorizing_seed"`
	IssuingSeed        string `mapstructure:"issuing_seed"`
	ReceivingAccountId string `mapstructure:"receiving_account_id"`
}

type Hooks struct {
	Receive string
	Error   string
}
