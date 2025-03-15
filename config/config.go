package config

import (
	"embed"
	"encoding/json"
	"os"
)

//go:embed default_config.json
var defaultConfigJSON embed.FS

// Config represents the configuration for the crawler
type Config struct {
	MaxDepth        int      `json:"max_depth"`
	MinSleep        int      `json:"min_sleep"`
	MaxSleep        int      `json:"max_sleep"`
	Timeout         int      `json:"timeout"`
	RootURLs        []string `json:"root_urls"`
	BlacklistedURLs []string `json:"blacklisted_urls"`
	UserAgents      []string `json:"user_agents"`
}

// LoadFromFile loads configuration from a JSON file
func LoadFromFile(filePath string) (*Config, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := &Config{}
	err = decoder.Decode(config)
	if err != nil {
		return nil, err
	}

	// Convert timeout=false to 0 (no timeout)
	if config.Timeout < 0 {
		config.Timeout = 0
	}

	return config, nil
}

// LoadDefaultConfig loads the default configuration embedded in the binary
func LoadDefaultConfig() (*Config, error) {
	data, err := defaultConfigJSON.ReadFile("default_config.json")
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = json.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
