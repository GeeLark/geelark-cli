package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// Config holds the GeeLark API credentials.
type Config struct {
	Token string `json:"token"`
	// BaseURL defaults to https://openapi.geelark.com (Cloud Phone API)
	BaseURL string `json:"base_url,omitempty"`
	// BrowserBaseURL defaults to http://localhost:40185 (Browser API, local)
	BrowserBaseURL string `json:"browser_base_url,omitempty"`
}

// DefaultBaseURL is the default GeeLark Cloud Phone API base URL.
const DefaultBaseURL = "https://openapi.geelark.com"

// DefaultBrowserBaseURL is the default GeeLark Browser API base URL (local).
const DefaultBrowserBaseURL = "http://localhost:40185"

// ConfigDir returns the path to the config directory.
func ConfigDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ".geelark"
	}
	return filepath.Join(home, ".geelark")
}

// ConfigFile returns the path to the config file.
func ConfigFile() string {
	return filepath.Join(ConfigDir(), "config.json")
}

// Load reads the config from disk.
func Load() (*Config, error) {
	data, err := os.ReadFile(ConfigFile())
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("config not found, run `geelark-cli config init` to set up credentials")
		}
		return nil, fmt.Errorf("failed to read config: %w", err)
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	if cfg.BaseURL == "" {
		cfg.BaseURL = DefaultBaseURL
	}
	if cfg.BrowserBaseURL == "" {
		cfg.BrowserBaseURL = DefaultBrowserBaseURL
	}

	return &cfg, nil
}

// Save writes the config to disk.
func Save(cfg *Config) error {
	dir := ConfigDir()
	if err := os.MkdirAll(dir, 0700); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	if err := os.WriteFile(ConfigFile(), data, 0600); err != nil {
		return fmt.Errorf("failed to write config: %w", err)
	}

	return nil
}
