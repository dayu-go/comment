package config

import (
	"time"

	"github.com/dayu-go/gkit/config"
	"github.com/dayu-go/gkit/config/file"
	"github.com/dayu-go/gkit/log"
)

var Conf Config

type Config struct {
	App struct {
		LogLevel string `json:"log_level"`
		Env      string `json:"env"`
	}
	Server Server
	DB     struct {
		Dayu DBConfig `json:"dayu"`
	}
}

type Server struct {
	Http struct {
		Network string        `json:"network"`
		Addr    string        `json:"addr"`
		Timeout time.Duration `json:"timeout"`
	} `json:"http"`
	Grpc struct {
		Network string `json:"network"`
		Addr    string `json:"addr"`
		Timeout int    `json:"timeout"`
	} `json:"grpc"`
}

type DBConfig struct {
	Driver string `json:"driver"`
	DSN    string `json:"dsn"`
}

func Load() error {
	// load app config
	data, err := loadConfig(file.NewSource("configs/config.yaml"), &Conf)
	if err != nil {
		return err
	}
	log.Infof("source config data: %s", data)
	return nil
}

func loadConfig(source config.Source, v interface{}) (string, error) {
	c := config.New(config.WithSource(
		source,
	))
	if err := c.Load(); err != nil {
		return "", err
	}
	d, err := c.Source()
	if err != nil {
		return "", err
	}
	err = c.Scan(v)
	return string(d), err
}
