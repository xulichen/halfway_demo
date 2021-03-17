package server

import (
	"github.com/spf13/viper"
	"github.com/xulichen/halfway/pkg/net/http"
)

func NewHttpServer() *http.Server {
	cfg := &http.Config{
		Name: viper.GetString("name"),
		Host: viper.GetString("golang-golang-server.host"),
		Port: viper.GetString("golang-golang-server.http-port"),
	}
	return http.NewServer(cfg)
}
