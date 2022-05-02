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

func DbInit(dbconfig *config.DbConfig) (*gorm.DB, error) {
	switch dbconfig.Type {
	case DB_TYPE_SQLITE:
		return sqliteInit(dbconfig.Config.(config.SqliteConfig))
	}
	return nil, errors.New("error database type")
}

func sqliteInit(config config.SqliteConfig, dst ...interface{}) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(config.DbFile), &gorm.Config{})
	if err != nil {
		logger.Fatalf("gorm init error: %s", err)
		return nil, err
	}
	
	if err := db.AutoMigrate(dst...); err != nil {
		logger.Fatalf("gorm migrate error: %s", err)
		return nil, err
	}
	return db, nil
}
