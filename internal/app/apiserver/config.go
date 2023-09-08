package apiserver

import "github.com/andreykvetinsky/http-rest-api/internal/app/store"

// Config ...

type Config struct {
	BindAddress string `toml:"bind_address"`
	LogLevel    string `tom:"log_level"`
	Store       *store.Config
}

// NewConfig returns a new Config
func NewConfig() *Config {
	return &Config{
		BindAddress: ":8080",
		LogLevel:    "debug",
		Store:       store.NewConfig(),
	}
}
