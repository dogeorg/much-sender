package config

import (
	"os"

	"github.com/BurntSushi/toml"
)

// Such Config struct for TOML configuration
type Config struct {
	SMTP struct {
		Server   string `toml:"server"`
		Port     int    `toml:"port"`
		Username string `toml:"username"`
		Password string `toml:"password"`
	} `toml:"smtp"`
	Security struct {
		BearerToken string `toml:"bearer_token"`
	} `toml:"security"`
	Server struct {
		Port int `toml:"port"`
	} `toml:"server"`
}

// Much LoadConfig loads the TOML configuration file
func LoadConfig(filename string) (Config, error) {
	var config Config
	data, err := os.ReadFile(filename)
	if err != nil {
		return config, err
	}
	err = toml.Unmarshal(data, &config)
	return config, err
}
