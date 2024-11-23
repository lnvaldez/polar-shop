package config 

import (
	"os"
	"strconv"
)

type Config struct {
	MongoURI string
	DatabaseName string 
	ServerPort int 
}

func Load() *Config {
	return &Config{
		MongoURI: os.Getenv("MONGO_URI"),
		DatabaseName: os.Getenv("DB_NAME"),
		ServerPort: getEnvAsInt("SERVER_PORT", 5001),
	}
}

func getEnvAsInt(key string, fallback int) int {
	port, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		return fallback
	}
	return port
}