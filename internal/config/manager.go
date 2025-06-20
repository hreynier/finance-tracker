package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

const (
	AppName    = "finance-tracker"
	ConfigFile = "config.yaml"
)

// Returns the app's config directory path
func GetConfigDir() (string, error) {
	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		return "", fmt.Errorf("Failed to get user config directory: %w", err)
	}

	return filepath.Join(userConfigDir, AppName), nil
}

// Returns the config file path within the config directory
func GetConfigFilePath() (string, error) {
	configDir, err := GetConfigDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(configDir, ConfigFile), nil
}

func CreateConfigDirIfNotExists() error {
	configDir, err := GetConfigDir()
	if err != nil {
		return err
	}

	err = os.MkdirAll(configDir, 0755)
	if err != nil {
		return fmt.Errorf("Failed to create config directory: %w", err)
	}

	return nil
}

func ConfigExists() (bool, error) {
	configPath, err := GetConfigFilePath()
	if err != nil {
		return false, err
	}

	// If there is no file this will return error that we can check
	_, err = os.Stat(configPath)
	if os.IsNotExist(err) {
		return false, nil
	}
	if err != nil {
		return false, fmt.Errorf("Failed to check if config file exists: %w", err)
	}

	return true, nil
}

// Create new config file with default values if it doesn't exist
func CreateDefaultConfig() error {

	if err := CreateConfigDirIfNotExists(); err != nil {
		return err
	}

	configPath, err := GetConfigFilePath()
	if err != nil {
		return err
	}

	defaultCfg := DefaultConfig()

	file, err := os.Create(configPath)
	if err != nil {
		return fmt.Errorf("Failed to create config file: %w", err)
	}
	defer file.Close()

	encoder := yaml.NewEncoder(file)
	defer encoder.Close()

	if err := encoder.Encode(defaultCfg); err != nil {
		return fmt.Errorf("Failed to write default config: %w", err)
	}

	fmt.Printf("DEBUG: Config file created at %s\n", configPath)
	return nil
}

func LoadConfig() (*Config, error) {
	exists, err := ConfigExists()
	if err != nil {
		return nil, err
	}

	if !exists {
		if err := CreateDefaultConfig(); err != nil {
			return nil, err
		}
	}

	configPath, err := GetConfigFilePath()
	if err != nil {
		return nil, err
	}

	file, err := os.Open(configPath)
	if err != nil {
		return nil, fmt.Errorf("Failed to open config file: %w", err)
	}

	defer file.Close()

	decoder := yaml.NewDecoder(file)

	var cfg Config
	if err := decoder.Decode(&cfg); err != nil {
		return nil, fmt.Errorf("Failed to load config: %w", err)
	}

	return &cfg, nil
}
