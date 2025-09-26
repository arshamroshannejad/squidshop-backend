package config

import (
	_ "embed"

	"gopkg.in/yaml.v3"

	"time"
)

//go:embed config.yaml
var configurations []byte

type App struct {
	Port    int    `yaml:"port"`
	Debug   bool   `yaml:"debug"`
	BaseAPI string `yaml:"base_api"`
}

type Postgres struct {
	Host            string        `yaml:"host"`
	Port            int           `yaml:"port"`
	Username        string        `yaml:"username"`
	Password        string        `yaml:"password"`
	Database        string        `yaml:"database"`
	MaxOpenConns    int           `yaml:"max_open_conns"`
	MaxIdleConns    int           `yaml:"max_idle_conns"`
	ConnMaxIdleTime time.Duration `yaml:"conn_max_idle_time"`
}

type Redis struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
}

type Jwt struct {
	Secret        string        `yaml:"secret"`
	AccessHourTTL time.Duration `yaml:"access_hour_ttl"`
}

type Sms struct {
	Service string `yaml:"sms_service"`
	Sender  string `yaml:"sms_sender"`
	ApiKey  string `yaml:"sms_api_key"`
}

type S3 struct {
	Bucket    string `yaml:"bucket"`
	Region    string `yaml:"region"`
	AccessKey string `yaml:"access_key"`
	SecretKey string `yaml:"secret_key"`
	Endpoint  string `yaml:"endpoint"`
}

type Config struct {
	App      *App      `yaml:"app"`
	Postgres *Postgres `yaml:"postgres"`
	Redis    *Redis    `yaml:"redis"`
	Jwt      *Jwt      `yaml:"jwt"`
	Sms      *Sms      `yaml:"sms"`
	S3       *S3       `yaml:"s3"`
}

func New() (*Config, error) {
	var cfg Config
	if err := yaml.Unmarshal(configurations, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
