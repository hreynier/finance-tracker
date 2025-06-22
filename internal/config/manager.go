package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

const (
	ConfigDir  = ".config"
	AppName    = "finance-tracker"
	ConfigFile = "config.yaml"
)

func GetHomeDir() (string, error) {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("Failed to get user home directory: %w", err)
	}

	return userHomeDir, nil
}

func GetConfigDir() (string, error) {
	userHomeDir, err := GetHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(userHomeDir, ConfigDir, AppName), nil
}

func GetConfigFilePath() (string, error) {
	configDir, err := GetConfigDir()
	if err != nil {
		return "", err
	}

	_, err = os.Stat(configDir)
	if os.IsNotExist(err) {
		err = os.MkdirAll(configDir, 0755)
		if err != nil {
			return "", fmt.Errorf("Failed to create config directory: %w", err)
		}
		return filepath.Join(configDir, ConfigFile), nil
	}

	if err != nil {
		return "", fmt.Errorf("Failed to check if config directory exists: %w", err)
	}

	return filepath.Join(configDir, ConfigFile), nil
}

func ConfigExists() (bool, error) {
	configPath, err := GetConfigFilePath()

	if err != nil {
		return false, err
	}

	_, err = os.Stat(configPath)
	if os.IsNotExist(err) {
		return false, nil
	}
	if err != nil {
		return false, fmt.Errorf("Failed to check if config file exists: %w", err)
	}

	return true, nil
}

func LoadConfig() (*Config, error) {
	configExists, err := ConfigExists()
	if err != nil {
		return nil, err
	}

	if !configExists {
		return nil, nil
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
		return nil, fmt.Errorf("Failed to read config: %w", err)
	}

	return &cfg, nil
}

func BuildConfig(cfg *Config, loadedConfig *Config) *Config {
	if loadedConfig == nil {
		return cfg
	}

	if loadedConfig.UserName != "" {
		cfg.UserName = loadedConfig.UserName
	}

	if loadedConfig.Currency != "" {
		cfg.Currency = loadedConfig.Currency
	}
	if loadedConfig.DataFile != "" {
		cfg.DataFile = loadedConfig.DataFile
	}
	if loadedConfig.DateFormat != "" {
		cfg.DateFormat = loadedConfig.DateFormat
	}
	if loadedConfig.Categories != nil && len(loadedConfig.Categories) > 0 {
		cfg.Categories = loadedConfig.Categories
	}

	return cfg
}

func GetConfiguration() (*Config, error) {
	fmt.Println("Getting configuration...")
	// Default GetConfiguration
	cfg := DefaultConfig()
	// Try to load config from file
	loadedConfig, err := LoadConfig()
	if err != nil {
		return nil, err
	}
	if loadedConfig != nil {
		cfg = BuildConfig(cfg, loadedConfig)
	}

	return cfg, nil
}
