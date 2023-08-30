package apiserver

// Config ...

type Config struct {
	BindAddress string `toml:"bind_address"`
	LogLevel    string `tom:"log_level"`
}

// NewConfig returns a new Config
func NewConfig() *Config {
	return &Config{
		BindAddress: ":8080",
		LogLevel:    "debug",
	}
}
