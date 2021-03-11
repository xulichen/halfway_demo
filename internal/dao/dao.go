package dao

import (
	"github.com/xulichen/halfway/pkg/cache"
	"gorm.io/gorm"
)

// data access object
type Dao struct {
	// TODO: 封装后改小写
	Mysql     *gorm.DB
	DataRedis *cache.Redis
	// ES && Tidb ...
}

func New(db *gorm.DB, client *cache.Redis) *Dao {
	return &Dao{
		Mysql:     db,
		DataRedis: client,
	}
}
