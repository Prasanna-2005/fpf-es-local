package config

import (
	"os"
)

type config struct {
	SockPath   string
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
}

var (
	Cfg *config
)

func InitialiseConfig() {
	Cfg = &config{
		SockPath:   os.Getenv("SOCKET_PATH"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBName:     os.Getenv("DB_NAME"),
	}

}
