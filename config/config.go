package config

import (
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

var (
	Port    = "3000"
	Release = "development"
	User    = "Anon"
	Time    = time.Now().Format(time.RFC3339)
	Hash    = "N/A"
	Module  = fx.Options(fx.Provide(NewConfig))
)

type (
	HTTP struct {
		Port string
	}

	Config struct {
		HTTP      `yaml:"http"`
		Hash      string
		Time      string
		User      string
		Release   string
		Directory string
	}
)

func NewConfig() *Config {

	cfg := &Config{
		HTTP: HTTP{
			Port: Port,
		},
		Hash:      Hash,
		Time:      Time,
		User:      User,
		Release:   Release,
		Directory: "",
	}

	return cfg
}

func (cfg *Config) InitConfig(dir string) error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	port, is_port_found := os.LookupEnv("PORT")

	if is_port_found {
		cfg.HTTP.Port = port
	}

	cfg.Directory = dir

	return nil
}
