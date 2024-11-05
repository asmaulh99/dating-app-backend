package configs

import "github.com/kelseyhightower/envconfig"

const (
	EnvDevelopment = "development"
	EnvProduction  = "production"
)

type Config struct {
	GeneralConfig
	DatabaseConfig
}

type GeneralConfig struct {
	Env           string `envconfig:"ENV" default:"staging"`
	RestPort      string `envconfig:"REST_PORT" default:"8080"`
	AuthSecretKey []byte `envconfig:"AUTH_SECRET_KEY" default:""`
}

type DatabaseConfig struct {
	DBDebugMode bool   `envconfig:"DB_DEBUG" default:"false"`
	DBHost      string `envconfig:"DB_HOST" default:""`
	DBUsername  string `envconfig:"DB_USERNAME" default:""`
	DBPassword  string `envconfig:"DB_PASSWORD" default:""`
	DBName      string `envconfig:"DB_NAME"`
}

var cfg Config

func GetConfig() *Config {
	envconfig.MustProcess("", &cfg)
	return &cfg
}
