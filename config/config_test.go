package config

import (
	"testing"
)

func TestLoadDefaultConfig(t *testing.T) {
	cfg, err := LoadDefaultConfig()
	if err != nil {
		t.Fatalf("Failed to load default config: %v", err)
	}

	// Verify that the default config has valid values
	if cfg.MaxDepth <= 0 {
		t.Errorf("Expected MaxDepth > 0, got %d", cfg.MaxDepth)
	}

	if cfg.MinSleep <= 0 {
		t.Errorf("Expected MinSleep > 0, got %d", cfg.MinSleep)
	}

	if cfg.MaxSleep <= 0 {
		t.Errorf("Expected MaxSleep > 0, got %d", cfg.MaxSleep)
	}

	if len(cfg.RootURLs) == 0 {
		t.Error("Expected RootURLs to have at least one URL")
	}

	if len(cfg.UserAgents) == 0 {
		t.Error("Expected UserAgents to have at least one user agent")
	}
}
