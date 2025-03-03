package config

import (
	"log"
	"os"
	"time"
	"flag"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env          string        `yaml:"env" env-default:"local"`
	Storage_path string        `yaml:"storage_path" env-required:"true"`
	TokenTTL     time.Duration `yaml:"token_ttl" env-required:"true"`
	GRPC         GRPCConfig    `yaml:"grpc"`
}

type GRPCConfig struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

func MustLoad() *Config {
	path := fetchConfigPath()
	if path == "" {
		panic("config path is empty")
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config path is not exist: " + path)
	}
	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err.Error())
	}

	return &cfg

	// configPath := os.Getenv("CONFIG_PATH")
	// if configPath == "" {
	// 	log.Fatal("CONFIG_PATH is not set")
	// }

	// // check if file exists
	// if _, err := os.Stat(configPath); os.IsNotExist(err) {
	// 	log.Fatalf("config file does not exist: %s", configPath)
	// }

	// var cfg Config

	// if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
	// 	log.Fatalf("cannot read config: %s", err)
	// }

	// return &cfg
}

func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")	
	}
	return res
}
