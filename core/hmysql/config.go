package hmysql

import "time"

type Config struct {
	UserName        string        `toml:"user_name"`
	Password        string        `toml:"password"`
	Host            string        `toml:"host"`
	Port            int           `toml:"port"`
	Database        string        `toml:"database"`
	MaxIdleConns    int           `toml:"max_idle_conns"`
	MaxOpenConns    int           `toml:"max_open_conns"`
	ConnMaxIdleTime time.Duration `toml:"conn_max_idle_time"`
	ConnMaxLifetime time.Duration `toml:"conn_max_lifetime"`
}

func DefaultConfig() Config {
	return Config{
		MaxOpenConns:    500,
		MaxIdleConns:    100,
		ConnMaxLifetime: time.Second * 3,
		ConnMaxIdleTime: time.Second * 3,
	}
}
