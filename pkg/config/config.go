package config

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
