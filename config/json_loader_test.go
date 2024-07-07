package config

import (
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	loader := JSONConfigLoader{}

	type TestConfig struct {
		TestField string `json:"test_field"`
	}

	configData := `{"test_field": "test_value"}`
	configFile, err := os.CreateTemp("", "config_test.json")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(configFile.Name())

	configFile.WriteString(configData)
	configFile.Close()

	var config TestConfig
	err = loader.Load(configFile.Name(), &config)
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	if config.TestField != "test_value" {
		t.Errorf("Expected 'test_value', got '%s'", config.TestField)
	}
}
