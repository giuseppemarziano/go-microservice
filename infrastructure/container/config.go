package container

import (
	"log"
	"os"
	"strconv"
)

type config struct {
	DatabaseDSN string
	HttpTimeout int
}

type Config interface {
	GetDatabaseDSN() string
	GetHttpTimeout() int
}

func (c *config) GetDatabaseDSN() string {
	return c.DatabaseDSN
}

func (c *config) GetHttpTimeout() int {
	return c.HttpTimeout
}

func NewConfig() Config {
	return &config{
		DatabaseDSN: getEnv("DATABASE_DSN", "default_dsn"),
		HttpTimeout: getEnvAsInt("HTTP_TIMEOUT", 30),
	}
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	valueStr, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		log.Printf("Warning: %s environment variable should be an integer, using default: %d", key, defaultValue)
		return defaultValue
	}
	return value
}
