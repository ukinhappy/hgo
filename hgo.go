package hgo

import (
	"github.com/spf13/viper"
	"github.com/ukinhappy/hgo/core/hgin"
	"github.com/ukinhappy/hgo/core/hmysql"
	"github.com/ukinhappy/hgo/core/hredis"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Hgo struct {
	configPath          string
	Http                *hgin.Http
	Redis               map[string]hredis.Redis
	Mysql               map[string]hmysql.Mysql
	gracefulExitTimeout time.Duration
}

func New() *Hgo {
	return &Hgo{
		configPath:          "config.toml",
		Http:                nil,
		Redis:               make(map[string]hredis.Redis),
		Mysql:               make(map[string]hmysql.Mysql),
		gracefulExitTimeout: time.Second * 10,
	}
}

func (h *Hgo) Init() {
	viper.SetConfigFile(h.configPath)
	viper.SetConfigType("toml")

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	// 初始化mysql
	mysqlMap := viper.GetStringMap("mysql")
	for k, _ := range mysqlMap {
		cfg := hmysql.DefaultConfig()
		viper.UnmarshalKey("mysql."+k, &cfg)

		h.Mysql[k] = hmysql.New(cfg)
	}

	// 初始化redis
	redisMap := viper.GetStringMap("redis")
	for k, _ := range redisMap {
		cfg := hredis.DefaultConfig()
		viper.UnmarshalKey("redis."+k, &cfg)
		h.Redis[k] = hredis.New(cfg)
	}

	// 初始化gin
	ginConfig := hgin.DefaultConfig()
	viper.UnmarshalKey("gin", &ginConfig)
	h.Http = hgin.New(ginConfig)
	h.Http.Init()
}

func (h *Hgo) Run() {
	h.Http.Run()
	h.gracefulExit()
}

func (h *Hgo) gracefulExit() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM, os.Kill)

	sig := <-signalChan
	log.Printf("catch signal, %+v", sig)

	time.Sleep(h.gracefulExitTimeout)
	log.Printf("server exiting")
}
