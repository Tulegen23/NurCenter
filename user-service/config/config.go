package config

import (
	"os"
)

type Config struct {
	UserDBHost     string
	UserDBPort     string
	UserDBUser     string
	UserDBPassword string
	UserDBName     string
	JWTSecret      string
}

func Load() *Config {
	return &Config{
		UserDBHost:     getEnv("USER_DB_HOST", "localhost"),
		UserDBPort:     getEnv("USER_DB_PORT", "5432"),
		UserDBUser:     getEnv("USER_DB_USER", "postgres"),
		UserDBPassword: getEnv("USER_DB_PASSWORD", "nurbol23"),
		UserDBName:     getEnv("USER_DB_NAME", "nurcenter_users"),
		JWTSecret:      getEnv("JWT_SECRET", "dXj9kPqL2mWvY8rT3zQhN5bV7cF4gJ0xA1eK9iU6oM="),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
