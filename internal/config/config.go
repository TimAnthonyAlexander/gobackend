package config

// Config holds the application configuration
type Config struct {
	Server ServerConfig
}

// ServerConfig holds the server configuration
type ServerConfig struct {
	Host string
	Port string
}

// NewConfig creates a new configuration with default values
func NewConfig() *Config {
	return &Config{
		Server: ServerConfig{
			Host: "127.0.0.1",
			Port: "3999",
		},
	}
} 