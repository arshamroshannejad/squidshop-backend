package config

import (
	"bytes"
	_ "embed"
	"errors"
	"github.com/spf13/viper"
	"time"
)

//go:embed config.yaml
var configurations []byte

type app struct {
	Port          int           `mapstructure:"port"`
	Debug         bool          `mapstructure:"debug"`
	BaseAPI       string        `mapstructure:"base_api"`
	Secret        string        `mapstructure:"secret"`
	AccessHourTTL time.Duration `mapstructure:"access_hour_ttl"`
}

type postgres struct {
	Host            string        `mapstructure:"host"`
	Port            int           `mapstructure:"port"`
	Username        string        `mapstructure:"username"`
	Password        string        `mapstructure:"password"`
	Database        string        `mapstructure:"database"`
	MaxOpenConns    int           `mapstructure:"max_open_conns"`
	MaxIdleConns    int           `mapstructure:"max_idle_conns"`
	ConnMaxIdleTime time.Duration `mapstructure:"conn_max_idle_time"`
}

type Config struct {
	App      *app      `mapstructure:"app"`
	Postgres *postgres `mapstructure:"postgres"`
}

func New() (*Config, error) {
	viper.SetConfigType("yaml")
	if err := viper.ReadConfig(bytes.NewBuffer(configurations)); err != nil {
		return nil, err
	}
	viper.AutomaticEnv()
	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	if cfg.Postgres == nil {
		return nil, errors.New("postgres configuration is required")
	}
	if cfg.App == nil {
		return nil, errors.New("app configuration is required")
	}
	return &cfg, nil
}
