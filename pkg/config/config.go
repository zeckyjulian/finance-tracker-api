package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Env           string
	Port          string
	Name          string
	AllowedOrigin string
}

type DatabaseConfig struct {
	Host         string
	Port         string
	Name         string
	User         string
	Password     string
	SSLMode      string
	MaxOpenConns int
	MaxIdleConns int
}

type RedisConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
}

type JWTConfig struct {
	AccessSecret       string
	RefreshSecret      string
	AccessExpiryMinute int
	RefreshExpiryDays  int
}

type SuperAdminConfig struct {
	Email    string
	Password string
	Name     string
}

type Config struct {
	App        AppConfig
	Database   DatabaseConfig
	Redis      RedisConfig
	JWT        JWTConfig
	SuperAdmin SuperAdminConfig
}

func (d *DatabaseConfig) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		d.Host, d.Port, d.Name, d.User, d.Password, d.SSLMode,
	)
}

func (r *RedisConfig) Addr() string {
	return fmt.Sprintf("%s:%s", r.Host, r.Port)
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}

func mustGetEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		panic(fmt.Sprintf("the environment variable %q is required", key))
	}
	return val
}

func getEnvInt(key string, fallback int) int {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	n, err := strconv.Atoi(val)
	if err != nil {
		return fallback
	}
	return n
}

func Load() (*Config, error) {
	if os.Getenv("APP_ENV") != "production" {
		if err := godotenv.Load(); err != nil {
			fmt.Println(".env file not found, using existed environment variable")
		}
	}

	cfg := &Config{
		App: AppConfig{
			Env:           getEnv("APP_ENV", "development"),
			Port:          getEnv("APP_PORT", "8080"),
			Name:          getEnv("APP_NAME", "finance-tracker"),
			AllowedOrigin: getEnv("APP_ALLOWED_ORIGIN", "http://localhost:3001"),
		},
		Database: DatabaseConfig{
			Host:         mustGetEnv("DB_HOST"),
			Port:         mustGetEnv("DB_PORT"),
			Name:         mustGetEnv("DB_NAME"),
			User:         mustGetEnv("DB_USER"),
			Password:     mustGetEnv("DB_PASSWORD"),
			SSLMode:      mustGetEnv("DB_SSL_MODE"),
			MaxOpenConns: getEnvInt("DB_MAX_OPEN_CONNS", 25),
			MaxIdleConns: getEnvInt("DB_MAX_IDLE_CONNS", 5),
		},
		Redis: RedisConfig{
			Host:     mustGetEnv("REDIS_HOST"),
			Port:     mustGetEnv("REDIS_PORT"),
			Password: getEnv("REDIS_PASSWORD", ""),
			DB:       getEnvInt("REDIS_DB", 0),
		},
		JWT: JWTConfig{
			AccessSecret:       mustGetEnv("JWT_ACCESS_SECRET"),
			RefreshSecret:      mustGetEnv("JWT_REFRESH_SECRET"),
			AccessExpiryMinute: getEnvInt("JWT_ACCESS_EXPIRY_MINUTES", 15),
			RefreshExpiryDays:  getEnvInt("JWT_REFRESH_EXPIRY_DAYS", 7),
		},
		SuperAdmin: SuperAdminConfig{
			Email:    mustGetEnv("SUPER_ADMIN_EMAIL"),
			Password: mustGetEnv("SUPER_ADMIN_PASSWORD"),
			Name:     getEnv("SUPER_ADMIN_NAME", "Super Admin"),
		},
	}

	return cfg, nil
}
