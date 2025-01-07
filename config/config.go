package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config -.
	Config struct {
		App   `yaml:"app"`
		HTTP  `yaml:"http"`
		Log   `yaml:"logger"`
		PG    `yaml:"postgres"`
		JWT   `yaml:"jwt"`
		Redis `yaml:"redis"`
		Gmail `yaml:"gmail"`
		MinIO `yaml:"minio"`
	}

	// App -.
	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}

	// HTTP -.
	HTTP struct {
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}

	// Log -.
	Log struct {
		Level string `env-required:"true" yaml:"log_level"   env:"LOG_LEVEL"`
	}

	// PG -.
	PG struct {
		PoolMax int    `env-required:"true" yaml:"pool_max" env:"PG_POOL_MAX"`
		URL     string `env-required:"true"                 env:"PG_URL"`
	}

	// JWT -.
	JWT struct {
		Secret string `env-required:"true" yaml:"secret" env:"JWT_SECRET"`
	}

	// Redis -.
	Redis struct {
		RedisHost string `env-required:"true" yaml:"host" env:"REDIS_HOST"`
		RedisPort int    `env-required:"true" yaml:"port" env:"REDIS_PORT"`
	}

	// Gmail -.
	Gmail struct {
		Email     string `env-required:"true" yaml:"email" env:"EMAIL"`
		EmailPass string `env-required:"true" yaml:"email_pass" env:"EMAIL_PASS"`
		Host      string `env-required:"true" yaml:"host" env:"SMTP_HOST"`
		Port      string    `env-required:"true" yaml:"port" env:"SMTP_PORT"`
	}

	// MinIO -.
	MinIO struct {
		Endpoint   string `env-required:"true" yaml:"endpoint" env:"MINIO_ENDPOINT"`
		AccessKey  string `env-required:"true" yaml:"access_key" env:"MINIO_ACCESS_KEY"`
		SecretKey  string `env-required:"true" yaml:"secret_key" env:"MINIO_SECRET_KEY"`
		BucketName string `env-required:"true" yaml:"bucket_name" env:"MINIO_BUCKET_NAME"`
	}
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
