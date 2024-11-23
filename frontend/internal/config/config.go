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
		Port: getPort("SERVER_PORT"),
	}
}

func getPort(key string) int {
	port, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		return 5001
	}
	return port
}