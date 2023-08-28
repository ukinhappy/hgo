package hredis

import (
	"time"
)

type Config struct {
	Addrs        []string      `toml:"addrs"`
	Username     string        `toml:"username"`
	Password     string        `toml:"password"`
	DialTimeout  time.Duration `toml:"dial_timeout"`
	ReadTimeout  time.Duration `toml:"read_timeout"`
	WriteTimeout time.Duration `toml:"write_timeout"`
	PoolSize     int           `toml:"pool_size"`
	MinIdleConns int           `toml:"min_idle_conns"`
	MaxConnAge   time.Duration `toml:"max_conn_age"`
	PoolTimeout  time.Duration `toml:"pool_timeout"`
	IdleTimeout  time.Duration `toml:"idle_timeout"`
}

func DefaultConfig() Config {
	return Config{
		DialTimeout:  time.Second * 3,
		ReadTimeout:  time.Second * 3,
		WriteTimeout: time.Second * 3,
		PoolSize:     100,
		MinIdleConns: 100,
		MaxConnAge:   time.Second * 3,
		PoolTimeout:  time.Second * 3,
		IdleTimeout:  time.Second * 3,
	}
}
