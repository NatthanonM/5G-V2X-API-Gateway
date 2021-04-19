package config

import "github.com/caarlos0/env"

// Config ...
type Config struct {
	// Mode: Production or Development
	Mode          string `env:"MODE" envDefault:"Production"`
	WebsiteOrigin string `env:"WEBSITE_ORIGIN" envDefault:"http://localhost:3000"`
	WebsiteDomain string `env:"WEBSITE_DOMAIN" envDefault:"localhost"`
	// Choose Starting Port
	Port                string `env:"PORT" envDefault:"8080"`
	AccessTokenLifetime string `env:"ACCESS_TOKEN_LIFETIME" envDefault:"8h"`

	// Declare Connection
	DataManagementServiceConnection string `env:"DATA_MANAGEMENT_CONNECTION" envDefault:"127.0.0.1:8082"`
	UserServiceConnection           string `env:"USER_CONNECTION" envDefault:"127.0.0.1:8083"`
}

// NewConfig ...
func NewConfig() *Config {
	c := &Config{}
	if err := env.Parse(c); err != nil {
		panic(err)
	}
	return c
}
