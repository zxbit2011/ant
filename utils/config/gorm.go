package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

// NewGorm 创建DB实例
func NewGorm(conf *Config) (*gorm.DB, error) {
	db, err := gorm.Open(conf.Gorm.DBType, conf.MySQL.DSN())
	if err != nil {
		return nil, err
	}
	if conf.Gorm.Debug {
		db = db.Debug()
	}
	err = db.DB().Ping()
	if err != nil {
		return nil, err
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(conf.Gorm.MaxIdleConns)
	db.DB().SetMaxOpenConns(conf.Gorm.MaxOpenConns)
	db.DB().SetConnMaxLifetime(time.Duration(conf.Gorm.MaxLifetime) * time.Second)
	return db, nil
}
