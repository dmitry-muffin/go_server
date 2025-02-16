package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

type Config struct {
	Env         string `yaml:"env" env-default:"local" env-required:"true"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server" env-required:"true"`
}

type HTTPServer struct {
	Addr         string        `yaml:"address" env-required:"true"`
	Timeout      time.Duration `yaml:"timeout" env-default:"5s"`
	IddleTimeout time.Duration `yaml:"iddle_timeout" env-default:"80s"`
}

func MustLoad() *Config {
	//get path
	configPath := "config/local.yml"
	if configPath == "" {
		log.Fatal("CONFIG_PATH variable not set")
	}

	//check existance
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatal("CONFIG_PATH does not exist", configPath)
	}
	//check readable
	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatal("cannot read config", err)
	}

	return &cfg
}
