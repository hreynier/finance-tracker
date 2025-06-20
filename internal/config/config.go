package config

type Config struct {
	Currency   string   `yaml:"currency"`
	DateFormat string   `yaml:"date_format"`
	Categories []string `yaml:"categories"`
	UserName   string   `yaml:"user_name"`
	DataFile   string   `yaml:"data_file"`
}
