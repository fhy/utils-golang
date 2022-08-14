package db

import (
	"context"
	"fmt"

	"github.com/fhy/utils-golang/config"

	logger "github.com/sirupsen/logrus"

	"github.com/go-redis/redis/v8"
)

func RedisInit(config *config.RedisConfig) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Host, config.Port),
		Password: config.Password,
		DB:       config.DB,
	})
	if _, err := rdb.Ping(context.Background()).Result(); err != nil {
		logger.Errorf("failed to connect redis: %s, error: %s", config.Host, err)
		return nil, err
	}
	return rdb, nil
}
