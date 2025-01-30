package config

import (
	"log"
	"os"
)

type Config struct {
	DBName   string
	User     string
	Password string
	Host     string
	DBPort   string
	Port     string
}

func LoadConfig() Config {

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		log.Fatal("DB_NAME environment variable is not set")
	}

	user := os.Getenv("DB_USER")
	if user == "" {
		log.Fatal("DB_USER environment variable is not set")
	}

	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		log.Fatal("DB_PASSWORD environment variable is not set")
	}

	host := os.Getenv("DB_HOST")
	if host == "" {
		log.Fatal("DB_HOST environment variable is not set")
	}

	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		log.Fatal("DB_PORT environment variable is not set")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable is not set")
	}

	return Config{
		DBName:   dbName,
		User:     user,
		Password: password,
		Host:     host,
		DBPort:   dbPort,
		Port:     port,
	}

}
