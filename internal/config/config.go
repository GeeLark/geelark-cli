package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
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

// Environment variables that override the stored configuration.
const (
	EnvToken          = "GEELARK_TOKEN"
	EnvBaseURL        = "GEELARK_BASE_URL"
	EnvBrowserBaseURL = "GEELARK_BROWSER_BASE_URL"
)

// envOrder lists the overrides for display, most important first.
var envOrder = []string{EnvToken, EnvBaseURL, EnvBrowserBaseURL}

// envValue reads an override and trims it. Secrets injected through .env
// files, Docker and Kubernetes routinely carry a trailing newline, which
// would otherwise be signed into the request and rejected as a bad token.
func envValue(name string) string {
	return strings.TrimSpace(os.Getenv(name))
}

// EnvOverrides lists the overrides that are currently set, so callers can
// point out that the config file is not what actually takes effect.
func EnvOverrides() []string {
	var names []string
	for _, name := range envOrder {
		if envValue(name) != "" {
			names = append(names, name)
		}
	}
	return names
}

// Load returns the effective settings for talking to the API, failing when no
// token is available from either source.
func Load() (*Config, error) {
	cfg, err := LoadForDisplay()
	if err != nil {
		return nil, err
	}

	// A missing file and a file carrying no token are the same problem to the
	// caller. Without this the second case reaches the API with an empty
	// token and comes back as a signature failure, which reads like the token
	// is wrong rather than absent.
	if cfg.Token == "" {
		return nil, fmt.Errorf("no API token configured, run `geelark-cli config init` to set up credentials")
	}

	return cfg, nil
}

// LoadForDisplay reads the config from disk, then applies environment variable
// overrides: GEELARK_TOKEN, GEELARK_BASE_URL, GEELARK_BROWSER_BASE_URL.
// Environment variables take precedence over the config file, so MCP
// servers and containers can inject credentials without touching the
// shared file. A config file is not required when GEELARK_TOKEN is set.
//
// Unlike Load it tolerates a missing token, because `config show` has to stay
// usable before setup — that is precisely when the resolved endpoints and the
// active overrides need inspecting.
func LoadForDisplay() (*Config, error) {
	var cfg Config

	data, err := os.ReadFile(ConfigFile())
	switch {
	case err == nil:
		if err := json.Unmarshal(data, &cfg); err != nil {
			return nil, fmt.Errorf("failed to parse config %s: %w", ConfigFile(), err)
		}
	case !os.IsNotExist(err):
		return nil, fmt.Errorf("failed to read config: %w", err)
	}

	// Hand-edited and copy-pasted files pick up stray whitespace just like
	// injected secrets do, so both sources are trimmed the same way.
	cfg.Token = strings.TrimSpace(cfg.Token)
	cfg.BaseURL = strings.TrimSpace(cfg.BaseURL)
	cfg.BrowserBaseURL = strings.TrimSpace(cfg.BrowserBaseURL)

	// Environment overrides (12-factor style: env beats file).
	if v := envValue(EnvToken); v != "" {
		cfg.Token = v
	}
	if v := envValue(EnvBaseURL); v != "" {
		cfg.BaseURL = v
	}
	if v := envValue(EnvBrowserBaseURL); v != "" {
		cfg.BrowserBaseURL = v
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
