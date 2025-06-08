package config

type Config struct {
	Currency   string   `json:"currency"`
	DateFormat string   `json:"date_format"`
	Categories []string `json:"categories"`
	UserName   string   `json:"user_name"`
}
