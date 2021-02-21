package config

import "os"

// Config data type that couses some config properties
type Config struct {
	ServerConfig *ServerConfig `json:"server_config"`
}

// ServerConfig contains the HTTP Server config properties
type ServerConfig struct {
	Address string
	Port    string
}

var (
	// BackendServerAddress config env variable
	BackendServerAddress string = "RANDENG__BACKEND__SERVER__ADDRESS"
	// BackendServerPort config env variable
	BackendServerPort string = "RANDENG__BACKEND__SERVER__PORT"
)

// Load parses environment variables and returns a pointer to a Config object, or an error if exists
func Load() (*Config, error) {
	config := &Config{}
	serverConfig := &ServerConfig{}
	addr := os.Getenv(BackendServerAddress)
	if addr == "" {
		addr = "127.0.0.1"
	}
	serverConfig.Address = addr
	port := os.Getenv(BackendServerPort)
	if port == "" {
		port = "8080"
	}
	serverConfig.Port = port
	config.ServerConfig = serverConfig
	return config, nil
}
