package config

import (
	"os"
)

type Config struct {
	DBUsername string
	DBPassword string
	DBHost     string
	DBName     string
	SMTPHost   string
	SMTPUser   string
	SMTPPass   string
	EmailFrom  string
}

func LoadConfig() *Config {
	return &Config{
		DBUsername: os.Getenv("DB_USERNAME"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBHost:     os.Getenv("DB_HOST"),
		DBName:     os.Getenv("DB_NAME"),
		SMTPHost:   os.Getenv("SMTP_HOST"),
		SMTPUser:   os.Getenv("SMTP_USER"),
		SMTPPass:   os.Getenv("SMTP_PASS"),
		EmailFrom:  os.Getenv("EMAIL_FROM"),
	}
}
