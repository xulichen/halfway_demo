package dao

import (
	"github.com/spf13/viper"
	"github.com/xulichen/halfway/pkg/cache"
)

func NewRedis() (*cache.Redis, error) {
	cfg := &cache.Config{
		Host:         viper.GetString("golang-redis-data.host"),
		Port:         viper.GetString("golang-redis-data.port"),
		Password:     viper.GetString("golang-redis-data.password"),
	}
	return cache.NewRedis(cfg)
}