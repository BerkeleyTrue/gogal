package config

import (
	"os"
	"time"

	"github.com/joho/godotenv"
)

var Port string = "3000"
var Release string = "development"
var User string = "Anon"
var Time string = time.Now().Format(time.RFC3339);
var Hash string = "N/A"

type (
	HTTP struct {
		Port    string
	}

	Config struct {
		HTTP    `yaml:"http"`
		Hash    string
		Time    string
		User    string
		Release string
	}
)

func NewConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
	  return nil, err
	}

	cfg := &Config{}

  port, is_found := os.LookupEnv("PORT")

	if is_found {
    cfg.HTTP.Port = port
  } else {
    cfg.HTTP.Port = Port
  }

	cfg.Hash = Hash
	cfg.Time = Time
	cfg.User = User

	cfg.Release = Release

	return cfg, nil
}
