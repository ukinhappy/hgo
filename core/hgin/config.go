package hgin

import "time"

type Config struct {
	Host         string        `toml:"host"`
	Port         int           `toml:"port"`
	ReadTimeout  time.Duration `toml:"read_timeout"`
	WriteTimeout time.Duration `toml:"write_timeout"`
}

func DefaultConfig() Config {
	return Config{
		Host:         "0.0.0.0",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}
