package hgo

import "time"

type Config struct {
	gracefulExitTimeout time.Duration
	configPath          string
}

var DefaultCfg = Config{
	gracefulExitTimeout: time.Second * 5,
	configPath:          "config.toml",
}
