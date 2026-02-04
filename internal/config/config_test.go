package config

import (
	"testing"
)

func TestLoad(t *testing.T) {
	cfg, err := Load()
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	if cfg == nil {
		t.Fatal("Expected non-nil config")
	}

	if cfg.Server.Port == "" {
		t.Error("Expected server port to be set")
	}

	if cfg.Server.Host == "" {
		t.Error("Expected server host to be set")
	}

	if cfg.Database.Host == "" {
		t.Error("Expected database host to be set")
	}

	if cfg.Logging.Level == "" {
		t.Error("Expected logging level to be set")
	}
}

func TestDefaultValues(t *testing.T) {
	cfg, err := Load()
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	// Test default values
	if cfg.Server.Port != "8080" {
		t.Errorf("Expected default port '8080', got '%s'", cfg.Server.Port)
	}

	if cfg.Database.Port != "5432" {
		t.Errorf("Expected default database port '5432', got '%s'", cfg.Database.Port)
	}

	if cfg.Logging.Level != "info" {
		t.Errorf("Expected default log level 'info', got '%s'", cfg.Logging.Level)
	}
}
