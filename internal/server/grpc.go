package server

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/xulichen/halfway/pkg/discovery"
	"github.com/xulichen/halfway/pkg/discovery/consul"
	"github.com/xulichen/halfway/pkg/net/rpc"
	demo "halfway_demo/proto"
	"halfway_demo/utils"
	"strconv"
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

func Discovery(srv *rpc.Server) {
	c1 := &discovery.ServerConfig{
		Address: viper.GetString("golang-consul.host"),
		Port:    viper.GetInt("golang-consul.port"),
		Token:   viper.GetString("golang-consul.token"),
	}
	port, _ := strconv.Atoi(srv.Port)
	c2 := &discovery.ServiceConfig{
		IP:              utils.LocalIP(),
		Port:            port,
		Name:            srv.Name,
		HealthyCheckURL: fmt.Sprintf("%s:%d/%s", utils.LocalIP(), port, srv.Name),
		IsRPC:           false,
	}
	res, err := consul.New(c1, c2)
	if err != nil {
		panic("consul init fail")
	}
	_ = srv.Discovery(res)
}

func RegisterDemoService(srv *rpc.Server, service demo.DemoServer) {
	demo.RegisterDemoServer(srv.Server, service)
}
