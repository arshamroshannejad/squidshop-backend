package config

import (
	_ "embed"
	"errors"

	"gopkg.in/yaml.v3"

	"time"
)

//go:embed config.yaml
var configurations []byte

type app struct {
	Port          int           `yaml:"port"`
	Debug         bool          `yaml:"debug"`
	BaseAPI       string        `yaml:"base_api"`
	Secret        string        `yaml:"secret"`
	AccessHourTTL time.Duration `yaml:"access_hour_ttl"`
	SmsService    string        `yaml:"sms_service"`
	SmsSender     string        `yaml:"sms_sender"`
	SmsApiKey     string        `yaml:"sms_api_key"`
}

type postgres struct {
	Host            string        `yaml:"host"`
	Port            int           `yaml:"port"`
	Username        string        `yaml:"username"`
	Password        string        `yaml:"password"`
	Database        string        `yaml:"database"`
	MaxOpenConns    int           `yaml:"max_open_conns"`
	MaxIdleConns    int           `yaml:"max_idle_conns"`
	ConnMaxIdleTime time.Duration `yaml:"conn_max_idle_time"`
}

type redis struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
}

type Config struct {
	App      *app      `yaml:"app"`
	Postgres *postgres `yaml:"postgres"`
	Redis    *redis    `yaml:"redis"`
}

func New() (*Config, error) {
	var cfg Config
	if err := yaml.Unmarshal(configurations, &cfg); err != nil {
		return nil, err
	}
	if cfg.Postgres == nil {
		return nil, errors.New("postgres configuration is required")
	}
	if cfg.Redis == nil {
		return nil, errors.New("redis configuration is required")
	}
	if cfg.App == nil {
		return nil, errors.New("app configuration is required")
	}
	return &cfg, nil
}
