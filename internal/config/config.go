package config

import "github.com/caarlos0/env"

// Config ...
type Config struct {
	// Mode: Production or Development
	Mode string `env:"MODE" envDefault:"Development"`
	// Choose Starting Port
	Port string `env:"PORT" envDefault:"8080"`
	// Declare Connection
	DataManagementServiceConnection string `env:"DATA_MANAGEMENT_CONNECTION" envDefault:"127.0.0.1:8082"`
}

// NewConfig ...
func NewConfig() *Config {
	c := &Config{}
	if err := env.Parse(c); err != nil {
		panic(err)
	}
	return c
}
