package config

import "time"

type AppConfig struct {
	Name    string
	Debug   bool
	Profile Profile
}

type HTTPConfig struct {
	Port int
}

type MongoDBConfig struct {
	URI              string
	Name             string
	ConnectTimeout   time.Duration
	OperationTimeout time.Duration
}

type PostgreSQLConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	DBName   string
}

type SMTPConfig struct {
	Port   int
	Domain string
}

type Config struct {
	App        AppConfig
	HTTP       HTTPConfig
	PostgreSQL PostgreSQLConfig
	SMTP       SMTPConfig
}

func NewAdapter() *Config {
	return &Config{
		App: AppConfig{
			Name:    Or(StringFromEnv, "APP_NAME", "bootstrapper"),
			Debug:   Or(BoolFromEnv, "APP_DEBUG", false),
			Profile: Must(ProfileFromEnv, "APP_PROFILE"),
		},
		SMTP: SMTPConfig{
			Port:   Or(IntFromEnv, "SMTP_PORT", 3025),
			Domain: Or(StringFromEnv, "SMTP_DOMAIN", "mail.slake.changhoi.kim"),
		},
		HTTP: HTTPConfig{
			Port: Or(IntFromEnv, "HTTP_PORT", 8080),
		},
		PostgreSQL: PostgreSQLConfig{
			Host:     Must(StringFromEnv, "POSTGRESQL_HOST"),
			Port:     Or(IntFromEnv, "POSTGRESQL_PORT", 5432),
			Username: Must(StringFromEnv, "POSTGRESQL_USERNAME"),
			Password: Must(StringFromEnv, "POSTGRESQL_PASSWORD"),
			DBName:   Must(StringFromEnv, "POSTGRESQL_DBNAME"),
		},
	}
}
