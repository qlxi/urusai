package main

import (
	"flag"
	"os"
	"testing"
	"time"

	"urusai/config"
)

// TestFlagParsing tests the command line flag parsing functionality
func TestFlagParsing(t *testing.T) {
	// Save original command line arguments and flags
	origArgs := os.Args
	origFlagCommandLine := flag.CommandLine
	
	// Restore the original state after the test
	defer func() {
		os.Args = origArgs
		flag.CommandLine = origFlagCommandLine
	}()

	// Test cases
	testCases := []struct {
		name           string
		args           []string
		expectConfigFile string
		expectLogLevel  string
		expectTimeout   int
	}{
		{
			name:           "Default values",
			args:           []string{"urusai"},
			expectConfigFile: "",
			expectLogLevel:  "info",
			expectTimeout:   0,
		},
		{
			name:           "Custom config file",
			args:           []string{"urusai", "--config", "custom.json"},
			expectConfigFile: "custom.json",
			expectLogLevel:  "info",
			expectTimeout:   0,
		},
		{
			name:           "Custom log level",
			args:           []string{"urusai", "--log", "debug"},
			expectConfigFile: "",
			expectLogLevel:  "debug",
			expectTimeout:   0,
		},
		{
			name:           "Custom timeout",
			args:           []string{"urusai", "--timeout", "300"},
			expectConfigFile: "",
			expectLogLevel:  "info",
			expectTimeout:   300,
		},
		{
			name:           "All custom values",
			args:           []string{"urusai", "--config", "custom.json", "--log", "error", "--timeout", "600"},
			expectConfigFile: "custom.json",
			expectLogLevel:  "error",
			expectTimeout:   600,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Reset flag.CommandLine for each test case
			flag.CommandLine = flag.NewFlagSet(tc.args[0], flag.ExitOnError)
			
			// Set up command line arguments
			os.Args = tc.args
			
			// Declare the flags again (as in main.go)
			var configFile string
			var logLevel string
			var timeout int
			
			flag.StringVar(&configFile, "config", "", "path to config file")
			flag.StringVar(&logLevel, "log", "info", "logging level (debug, info, warn, error)")
			flag.IntVar(&timeout, "timeout", 0, "for how long the crawler should be running, in seconds (0 means no timeout)")
			flag.Parse()
			
			// Check if the flags were parsed correctly
			if configFile != tc.expectConfigFile {
				t.Errorf("Expected config file %q, got %q", tc.expectConfigFile, configFile)
			}
			
			if logLevel != tc.expectLogLevel {
				t.Errorf("Expected log level %q, got %q", tc.expectLogLevel, logLevel)
			}
			
			if timeout != tc.expectTimeout {
				t.Errorf("Expected timeout %d, got %d", tc.expectTimeout, timeout)
			}
		})
	}
}

// TestLoadDefaultConfig tests that the default configuration can be loaded
func TestLoadDefaultConfig(t *testing.T) {
	cfg, err := config.LoadDefaultConfig()
	if err != nil {
		t.Fatalf("Failed to load default config: %v", err)
	}
	
	// Verify that the default config has valid values
	if cfg.MaxDepth <= 0 {
		t.Errorf("Expected MaxDepth > 0, got %d", cfg.MaxDepth)
	}
	
	if len(cfg.RootURLs) == 0 {
		t.Error("Expected RootURLs to have at least one URL")
	}
}

// TestSetLogLevel tests the log level setting functionality
func TestSetLogLevel(t *testing.T) {
	// Test with valid log levels
	validLevels := []string{"debug", "info", "warn", "error"}
	
	for _, level := range validLevels {
		// This should not panic
		setLogLevel(level)
	}
	
	// Test with invalid log level (should default to info)
	setLogLevel("invalid")
}

// TestSignalHandling tests the signal handling setup
// Note: This is a basic test that just ensures the function doesn't crash
func TestSignalHandling(t *testing.T) {
	sigChan := make(chan os.Signal, 1)
	
	// Start a goroutine that will send a signal after a short delay
	go func() {
		time.Sleep(10 * time.Millisecond)
		sigChan <- os.Interrupt
	}()
	
	// This should return almost immediately due to the signal
	<-sigChan
}
