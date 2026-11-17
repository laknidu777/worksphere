package config

import (
	"os"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	Port                   string
	Environment            string
	DatabaseURL            string
	RedisURL               string
	CorsAllowedOrigins     []string
	ShutdownTimeoutSeconds time.Duration
	EnableMockAuth         bool
}

func Load() *Config {
	port := getEnv("PORT", "8080")
	env := getEnv("APP_ENV", "development")
	dbURL := getEnv("DATABASE_URL", "postgres://worksphere:worksphere_local_secret@localhost:5432/worksphere?sslmode=disable")
	redisURL := getEnv("REDIS_URL", "redis://localhost:6379/0")
	originsStr := getEnv("CORS_ALLOWED_ORIGINS", "http://localhost:3000")
	timeoutSecStr := getEnv("SHUTDOWN_TIMEOUT_SECONDS", "15")

	timeoutSec, err := strconv.Atoi(timeoutSecStr)
	if err != nil || timeoutSec <= 0 {
		timeoutSec = 15
	}

	mockAuth := getEnv("ENABLE_MOCK_AUTH", "true") == "true"

	origins := strings.Split(originsStr, ",")
	for i := range origins {
		origins[i] = strings.TrimSpace(origins[i])
	}

	return &Config{
		Port:                   port,
		Environment:            env,
		DatabaseURL:            dbURL,
		RedisURL:               redisURL,
		CorsAllowedOrigins:     origins,
		ShutdownTimeoutSeconds: time.Duration(timeoutSec) * time.Second,
		EnableMockAuth:         mockAuth,
	}
}

func getEnv(key, defaultVal string) string {
	val, ok := os.LookupEnv(key)
	if !ok || strings.TrimSpace(val) == "" {
		return defaultVal
	}
	return strings.TrimSpace(val)
}
