package rpc

func NewConfig(address, clientId, clientSecret string) *Config {
	return &Config{
		Address:      address,
		ClientID:     clientId,
		ClientSecret: clientSecret,
	}
}

// Config 客户端配置
type Config struct {
	Address      string `json:"adress" toml:"adress" yaml:"adress" env:"MCENTER_ADDRESS"`
	ClientID     string `json:"client_id" toml:"client_id" yaml:"client_id" env:"MCENTER_CLINET_ID"`
	ClientSecret string `json:"client_secret" toml:"client_secret" yaml:"client_secret" env:"MCENTER_CLIENT_SECRET"`
}

func (c *Config) Credentials() *Authentication {
	return NewAuthentication(c.ClientID, c.ClientSecret)
}
