package models

type TomlConfig struct {
	Owner    OwnerInfo
	DB       Database `toml:"database"`
	SendGrid Sendgrid
}

type OwnerInfo struct {
	Name string
}

type Database struct {
	Server   string
	Port     int
	Name     string
	User     string
	Password string
}

type Sendgrid struct {
	APIKEY string `toml:"api_key"`
}
