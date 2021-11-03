package db

import (
	"fmt"
	"strconv"
)

type Config struct {
	Host     string `json:"host" yaml:"host"`
	Port     string `json:"port" yaml:"port"`
	Username string `json:"username" yaml:"username"`
	Password string `json:"password" yaml:"password"`
	Database string `json:"database" yaml:"database"`
}

func (c *Config) Dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.Username, c.Password, c.Host, c.PortToInt(), c.Database)
}

func (c *Config) PortToInt() int {
	intPort, _ := strconv.Atoi(c.Port)

	return intPort
}
