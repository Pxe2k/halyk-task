package config

import (
	"os"
)

type Config struct {
	Port string `json:"port"`
	PG   DBConn `json:"pg"`
}

type DBConn struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	DBName   string `json:"db"`
	SSLMode  string `json:"ssl_mode"`
	Timezone string `json:"timezone"`
}

func New(path string) *Config {
	return &Config{
		Port: getEnv("PORT"),
		PG: DBConn{
			User:     getEnv("PG_USER"),
			Password: getEnv("PG_PASSWORD"),
			Host:     getEnv("PG_HOST"),
			Port:     getEnv("PG_PORT"),
			DBName:   getEnv("PG_NAME"),
			SSLMode:  getEnv("PG_SSL_MODE"),
			Timezone: getEnv("PG_TIMEZONE"),
		},
	}
}

func getEnv(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return ""
}
