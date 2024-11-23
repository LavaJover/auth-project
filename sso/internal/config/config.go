package config

import (
	"flag"
	"os"
	"time"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct{
	Env string `yaml:"env" env-default:"env"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	GRPC GRPConfig `yaml:"grpc"`
	TokenTTL time.Duration `yaml:"token_ttl" env-required:"true"`
}

type GRPConfig struct{
	Port int `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

func MustLoad() *Config{
	path := fetchConfigPath()
	if path == ""{
		panic("Config path is empty")
	}

	if _, err := os.Stat(path); os.IsNotExist(err){
		panic("Config file does not exist: " + path)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil{
		panic("Failed to read config: " + err.Error())
	}

	return &cfg

}

func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == ""{
		res = os.Getenv("CONFIG_PATH")
	}
	return res
}