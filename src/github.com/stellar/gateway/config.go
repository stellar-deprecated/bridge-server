package gateway

type Config struct {
	Port     int
	Horizon  string
	Database struct {
		Type string
		Url  string
	}
	Accounts struct {
		AuthorizingSeed string   `mapstructure:"authorizing"`
		IssuingSeed     string   `mapstructure:"issuing"`
		ReceivingSeed   string   `mapstructure:"receiving"`
		ChannelsSeeds   []string `mapstructure:"channels"`
	}
}
