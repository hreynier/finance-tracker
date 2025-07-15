package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/charmbracelet/log"
	"gopkg.in/yaml.v3"
)

const (
	ConfigDir  = ".config"
	AppName    = "finance-tracker"
	ConfigFile = "config.yaml"
)

type Config struct {
	Currency   string   `yaml:"currency"`
	DateFormat string   `yaml:"date_format"`
	Categories []string `yaml:"categories"`
	UserName   string   `yaml:"user_name"`
	DataFile   string   `yaml:"data_file"`
}

func (cfg Config) String() string {
	configStr := fmt.Sprintf("Name: %s, Currency: %s, Date Format: %s, Categories: %v, Data File: %s",
		cfg.UserName, cfg.Currency, cfg.DateFormat, cfg.Categories, cfg.DataFile)
	log.Debug("Config string generated", "config", configStr)
	return configStr
}

func DefaultConfig() *Config {
	log.Debug("Creating default configuration")
	cfg := &Config{
		Currency:   "GBP",
		DateFormat: "02/01/2006", // DD/MM/YYYY
		UserName:   "User",
		Categories: []string{
			"Rent",
			"Utilities",
			"Insurance",
			"Subscriptions",
			"Groceries",
			"Eating Out",
			"Coffee",
			"Public Transport",
			"Car Expenses",
			"Cycling",
			"Other Transport",
			"Work & Hobbies",
			"Activites & Socialising",
			"Drinking & Indulgences",
			"Shopping",
			"Gifts",
			"Self-Care & Health",
			"Holiday Saving",
			"Holiday Spending",
			"Other",
		},
		DataFile: "transactions.json",
	}
	log.Debug("Default config created", "categories_count", len(cfg.Categories), "user", cfg.UserName)
	return cfg
}

func ConfigExists() (bool, error) {
	log.Debug("Checking if config file exists")
	configPath, err := getConfigFilePath()
	if err != nil {
		log.Debug("Failed to get config file path", "err", err)
		return false, err
	}

	log.Debug("Checking config file", "path", configPath)
	_, err = os.Stat(configPath)
	if os.IsNotExist(err) {
		log.Debug("Config file does not exist", "path", configPath)
		return false, nil
	}
	if err != nil {
		log.Error("Failed to check if config file exists", "err", err, "path", configPath)
		return false, err
	}

	log.Debug("Config file exists", "path", configPath)
	return true, nil
}

func LoadConfig() (*Config, error) {
	log.Debug("Loading configuration from file")

	configExists, err := ConfigExists()
	if err != nil {
		return nil, err
	}

	if !configExists {
		log.Debug("Config file doesn't exist, returning nil")
		return nil, nil
	}

	configPath, err := getConfigFilePath() // Fixed: was GetConfigFilePath()
	if err != nil {
		return nil, err
	}

	log.Debug("Opening config file", "path", configPath)
	file, err := os.Open(configPath)
	if err != nil {
		log.Error("Failed to open config file", "err", err, "path", configPath)
		return nil, fmt.Errorf("Failed to open config file: %w", err)
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)

	var cfg Config
	log.Debug("Decoding YAML config")
	if err := decoder.Decode(&cfg); err != nil {
		log.Error("Failed to decode config YAML", "err", err, "path", configPath)
		return nil, fmt.Errorf("Failed to read config: %w", err)
	}

	log.Debug("Config loaded successfully", "user", cfg.UserName, "currency", cfg.Currency, "categories_count", len(cfg.Categories))
	return &cfg, nil
}

func GetConfiguration() (*Config, error) {
	log.Debug("Getting configuration...")

	// Default configuration
	cfg := DefaultConfig()

	// Try to load config from file
	log.Debug("Attempting to load config from file")
	loadedConfig, err := LoadConfig()
	if err != nil {
		log.Error("Failed to load config from file", "err", err)
		return nil, err
	}

	if loadedConfig != nil {
		log.Debug("Merging loaded config with defaults")
		cfg = buildConfig(cfg, loadedConfig)
		log.Debug("Config merge completed", "final_user", cfg.UserName)
	} else {
		log.Debug("No config file found, using defaults")
	}

	return cfg, nil
}

func getHomeDir() (string, error) {
	log.Debug("Getting user home directory")
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		log.Error("Failed to get user home directory", "err", err)
		return "", err
	}

	log.Debug("User home directory found", "path", userHomeDir)
	return userHomeDir, nil
}

func getConfigDir() (string, error) {
	log.Debug("Getting config directory path")
	userHomeDir, err := getHomeDir()
	if err != nil {
		return "", err
	}

	configDir := filepath.Join(userHomeDir, ConfigDir, AppName)
	log.Debug("Config directory path determined", "path", configDir)
	return configDir, nil
}

func getConfigFilePath() (string, error) {
	log.Debug("Getting config file path")
	configDir, err := getConfigDir()
	if err != nil {
		return "", err
	}

	log.Debug("Checking if config directory exists", "dir", configDir)
	_, err = os.Stat(configDir)
	if os.IsNotExist(err) {
		log.Debug("Config directory doesn't exist, creating it", "dir", configDir)
		err = os.MkdirAll(configDir, 0755)
		if err != nil {
			log.Error("Failed to create config directory", "err", err, "dir", configDir)
			return "", err
		}
		log.Debug("Config directory created successfully", "dir", configDir)
	} else if err != nil {
		log.Error("Failed to check if config directory exists", "err", err, "dir", configDir)
		return "", err
	} else {
		log.Debug("Config directory already exists", "dir", configDir)
	}

	configPath := filepath.Join(configDir, ConfigFile)
	log.Debug("Config file path determined", "path", configPath)
	return configPath, nil
}

func buildConfig(cfg *Config, loadedConfig *Config) *Config {
	log.Debug("Building merged configuration")

	if loadedConfig == nil {
		log.Debug("No loaded config to merge, returning default")
		return cfg
	}

	originalUser := cfg.UserName
	if loadedConfig.UserName != "" {
		log.Debug("Overriding username", "from", cfg.UserName, "to", loadedConfig.UserName)
		cfg.UserName = loadedConfig.UserName
	}

	if loadedConfig.Currency != "" {
		log.Debug("Overriding currency", "from", cfg.Currency, "to", loadedConfig.Currency)
		cfg.Currency = loadedConfig.Currency
	}

	if loadedConfig.DataFile != "" {
		log.Debug("Overriding data file", "from", cfg.DataFile, "to", loadedConfig.DataFile)
		cfg.DataFile = loadedConfig.DataFile
	}

	if loadedConfig.DateFormat != "" {
		log.Debug("Overriding date format", "from", cfg.DateFormat, "to", loadedConfig.DateFormat)
		cfg.DateFormat = loadedConfig.DateFormat
	}

	if loadedConfig.Categories != nil && len(loadedConfig.Categories) > 0 {
		log.Debug("Overriding categories", "from_count", len(cfg.Categories), "to_count", len(loadedConfig.Categories))
		cfg.Categories = loadedConfig.Categories
	}

	log.Debug("Config merge completed", "original_user", originalUser, "final_user", cfg.UserName)
	return cfg
}
