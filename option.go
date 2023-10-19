package hgo

import "time"

type Option func(config *Config)

func WithConfigPath(cfgPath string) Option {
	return func(cfg *Config) {
		cfg.configPath = cfgPath
	}
}

func WithGracefulExitTimeout(gracefulExitTimeout time.Duration) Option {
	return func(cfg *Config) {
		cfg.gracefulExitTimeout = gracefulExitTimeout
	}
}
