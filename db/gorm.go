package db

import (
	"base/config"
	"errors"

	logger "github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	DB_TYPE_SQLITE = "sqlite"
)

func DbInit(typ string, cfg interface{}) (*gorm.DB, error) {
	switch typ {
	case DB_TYPE_SQLITE:
		return sqliteInit(cfg.(config.SqliteConfig))
	}
	return nil, errors.New("error database type")
}

func sqliteInit(config config.SqliteConfig) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(config.DbFile), &gorm.Config{})
	if err != nil {
		logger.Fatalf("gorm init error: %s", err)
		return nil, err
	}
	return db, nil
}
