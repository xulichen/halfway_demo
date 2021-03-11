package dao

import (
	"github.com/spf13/viper"
	"github.com/xulichen/halfway/pkg/db"
	"gorm.io/gorm"
)

func NewMysql() *gorm.DB {
	cfg := &db.MySqlConfig{
		Host:     viper.GetString("golang-mysql.host"),
		Port:     viper.GetString("golang-mysql.port"),
		User:     viper.GetString("golang-mysql.user"),
		Password: viper.GetString("golang-mysql.password"),
		DB:       viper.GetString("golang-mysql.db"),
	}
	return db.NewDBWithAPM(cfg)
}
