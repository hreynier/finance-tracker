package config

import (
	"fmt"
)

type Config struct {
	Currency   string   `yaml:"currency"`
	DateFormat string   `yaml:"date_format"`
	Categories []string `yaml:"categories"`
	UserName   string   `yaml:"user_name"`
	DataFile   string   `yaml:"data_file"`
}

func (cfg Config) String() string {
	return fmt.Sprintf("Name: %s, Currency: %s, Date Format: %s, Categories: %v, Data File: %s", cfg.UserName, cfg.Currency, cfg.DateFormat, cfg.Categories, cfg.DataFile)
}
