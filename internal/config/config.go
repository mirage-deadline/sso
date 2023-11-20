package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"time"
)

type Config struct {
	Env         string        `yaml:"env"`
	StoragePath string        `yaml:"storage_path" env-required:"true"`
	TokenTTL    time.Duration `yaml:"token_ttl" env-required:"true"`
	GRPC        GRPCConfig
}

type GRPCConfig struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

func NewConfig() *Config {
	path := fetchConfigPath()
	if path == "" {
		panic("path not found")
	}
	if _, err := os.Stat(path); os.IsExist(err) {
		panic("config file does not exist" + path)
	}

	var cfg Config
	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("failed to read config" + err.Error())
	}
	return &cfg
}

func fetchConfigPath() string {
	var p string
	flag.StringVar(&p, "config-path", "", "path to config")
	flag.Parse()
	if p == "" {
		return os.Getenv("CONFIG_PATH")
	}
	return p
}
