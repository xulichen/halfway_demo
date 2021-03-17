package server

import (
	"fmt"
	"halfway_demo/utils"
	"strconv"

	"github.com/spf13/viper"
	"github.com/xulichen/halfway/pkg/discovery"
	"github.com/xulichen/halfway/pkg/discovery/consul"
	"github.com/xulichen/halfway/pkg/net/rpc"
)

func NewRpcServer() *rpc.Server {
	cfg := &rpc.Config{
		Name:    viper.GetString("name"),
		NetWork: "tcp",
		Host:    viper.GetString("golang-golang-server.host"),
		Port:    viper.GetString("golang-golang-server.rpc-port"),
		Debug:   true,
	}
	srv := rpc.NewServer(*cfg)
	Discovery(srv)
	return srv
}

// consul 服务注册
func Discovery(srv *rpc.Server) {
	c1 := &discovery.ServerConfig{
		Address: viper.GetString("golang-consul.host"),
		Port:    viper.GetInt("golang-consul.port"),
		Token:   viper.GetString("golang-consul.token"),
	}
	res, err := consul.New(c1)
	if err != nil {
		panic("consul init fail")
	}
	port, _ := strconv.Atoi(srv.Port)
	c2 := &discovery.ServiceConfig{
		IP:              utils.LocalIP(),
		Port:            port,
		Name:            srv.Name,
		HealthyCheckURL: fmt.Sprintf("%s:%d/%s", utils.LocalIP(), port, srv.Name),
		IsRPC:           true,
	}
	res.WithServiceConfig(c2)
	_ = srv.Discovery(res)
}
