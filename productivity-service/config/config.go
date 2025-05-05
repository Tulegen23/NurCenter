package config

import "os"

type Config struct {
	ProductivityDBHost     string
	ProductivityDBPort     string
	ProductivityDBUser     string
	ProductivityDBPassword string
	ProductivityDBName     string
	JWTSecret              string
	UserServiceURL         string
}

func Load() *Config {
	return &Config{
		ProductivityDBHost:     getEnv("PRODUCTIVITY_DB_HOST", "localhost"),
		ProductivityDBPort:     getEnv("PRODUCTIVITY_DB_PORT", "5432"),
		ProductivityDBUser:     getEnv("PRODUCTIVITY_DB_USER", "postgres"),
		ProductivityDBPassword: getEnv("PRODUCTIVITY_DB_PASSWORD", "nurbol23"),
		ProductivityDBName:     getEnv("PRODUCTIVITY_DB_NAME", "nurcenter_productivity"),
		JWTSecret:              getEnv("JWT_SECRET", "dXj9kPqL2mWvY8rT3zQhN5bV7cF4gJ0xA1eK9iU6oM="),
		UserServiceURL:         getEnv("USER_SERVICE_URL", "http://localhost:8081"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
