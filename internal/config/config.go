package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUser     string
	DBHost     string
	DBName     string
	DBPort     string
	DBPassword string
	DBSecure   string
	JWTSecret  string
	ServerPort string
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func findEnvFile() string {
	dir, _ := os.Getwd()
	for {
		envPath := filepath.Join(dir, ".env")
		if _, err := os.Stat(envPath); err == nil {
			return envPath
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return ""
		}
		dir = parent
	}
}

func Load() (*Config, error) {
	if path := findEnvFile(); path != "" {
		_ = godotenv.Load(path)
	}

	cfg := &Config{
    DBUser:     getEnv("DB_USER", ""),
    DBPassword: getEnv("DB_PASS", ""),
    DBHost:     getEnv("DB_HOST", "localhost"),
    DBPort:     getEnv("DB_PORT", "5433"),
    DBName:     getEnv("DB_NAME", ""),
    DBSecure:   getEnv("DB_SECURE", "disable"),
	ServerPort: getEnv("SERVER_PORT", "8000"),
	JWTSecret:  getEnv("JWT_SECRET_TEST", ""),
	}
	if cfg.DBUser == "" {
		return nil, fmt.Errorf("DB_USER missing")
	}
	if cfg.DBPassword == "" {
		return nil, fmt.Errorf("DB_PASSWORD  Missing")
	}
	if cfg.DBName == "" {
		return nil, fmt.Errorf("DB_NAME  Missing")
	}

	return cfg, nil
}
