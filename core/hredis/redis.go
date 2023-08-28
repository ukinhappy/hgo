package hredis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type Redis struct {
	*redis.ClusterClient
	Config Config
}

func New(cfg Config) Redis {
	config := &redis.ClusterOptions{
		Addrs:        cfg.Addrs,
		Username:     cfg.Username,
		Password:     cfg.Password,
		DialTimeout:  cfg.DialTimeout,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		PoolSize:     cfg.PoolSize,
		MinIdleConns: cfg.MinIdleConns,
		MaxConnAge:   cfg.MaxConnAge,
		PoolTimeout:  cfg.PoolTimeout,
		IdleTimeout:  cfg.IdleTimeout,
	}
	client := redis.NewClusterClient(config)
	err := client.Ping(context.Background()).Err()
	if err != nil {
		panic(err)
	}
	return Redis{ClusterClient: client, Config: cfg}
}
