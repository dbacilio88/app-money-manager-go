package config

import (
	"os"
)

type Config struct {
	Env               string
	Port              string
	MongoURI          string
	MongoDB           string
	JWTSecret         string
	JWTExpiration     string
	RefreshExpiration string
	GoogleClientID    string
	GoogleSecret      string
	GoogleRedirectURL string
	AppURL            string
}

func Load() *Config {
	return &Config{
		Env:               getEnv("ENV", "development"),
		Port:              getEnv("PORT", "8080"),
		MongoURI:          getEnv("MONGO_URI", "mongodb://localhost:27017"),
		MongoDB:           getEnv("MONGO_DB", "control_financiero"),
		JWTSecret:         getEnv("JWT_SECRET", "your-secret-key-change-in-production"),
		JWTExpiration:     getEnv("JWT_EXPIRATION", "15m"),
		RefreshExpiration: getEnv("REFRESH_EXPIRATION", "168h"),
		GoogleClientID:    getEnv("GOOGLE_CLIENT_ID", ""),
		GoogleSecret:      getEnv("GOOGLE_CLIENT_SECRET", ""),
		GoogleRedirectURL: getEnv("GOOGLE_REDIRECT_URL", ""),
		AppURL:            getEnv("APP_URL", "http://localhost:8080"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
