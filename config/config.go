package config

import (
	"fmt"
)

type Config struct {
	Name     string
	User     string
	Password string
	Host     string
	Port     string
	URI      *string

	Timeout *int64
}

func (cfg *Config) SetURI() error {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)

	cfg.URI = &uri

	return nil
}

func (cfg *Config) DBName() string {
	return cfg.Name
}

func (cfg *Config) DBUser() string {
	return cfg.User
}

func (cfg *Config) DBPassword() string {
	return cfg.Password
}

func (cfg *Config) DBHost() string {
	return cfg.Host
}

func (cfg *Config) DBPort() string {
	return cfg.Port
}

func (cfg *Config) DBURI() *string {
	return cfg.URI
}

func (cfg *Config) DBTimeout() *int64 {
	return cfg.Timeout
}
