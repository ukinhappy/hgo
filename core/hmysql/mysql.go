package hmysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Mysql struct {
	*gorm.DB
	Config Config
}

func New(cfg Config) Mysql {
	dsns := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", cfg.UserName, cfg.Password, cfg.Host, cfg.Port, cfg.Database) + "parseTime=true&charset=utf8&loc=Asia%2fShanghai"
	db, err := gorm.Open(mysql.Open(dsns))
	if err != nil {
		panic(err)
	}

	DB, err := db.DB()
	if err != nil {
		panic(err)
	}
	DB.SetMaxIdleConns(cfg.MaxIdleConns)
	DB.SetMaxOpenConns(cfg.MaxOpenConns)
	DB.SetConnMaxIdleTime(cfg.ConnMaxIdleTime)
	DB.SetConnMaxLifetime(cfg.ConnMaxLifetime)
	return Mysql{
		DB:     db,
		Config: cfg,
	}
}
