package server

import (
	"data-seeder/internal/db"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	Port int        `json:"port" yaml:"port"`
	Db   *db.Config `json:"db" yaml:"db"`
}

func NewConfigFromFile(name string) (*Config, error) {
	var cfg *Config

	config, err := ioutil.ReadFile(name)

	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(config, &cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
