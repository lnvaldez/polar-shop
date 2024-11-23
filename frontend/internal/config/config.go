package config 

import (
	"os"
	"strconv"
)

type Config struct {
	MongoURI string
	DatabaseName string 
	Port int 
}

func Load() *Config {
	return &Config{
		MongoURI: os.Getenv("MONGO_URI"),
		DatabaseName: os.Getenv("DB_NAME"),
		Port: getEnvAsInt("SERVER_PORT", 5001),
	}
}

func getEnvAsInt(key string, fallback int) int {
	port, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		return fallback
	}
	return port
}