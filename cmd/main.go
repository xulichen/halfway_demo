package main

import (
	"fmt"
	"halfway_demo/internal/dao"
	"halfway_demo/internal/routers"
	"halfway_demo/internal/server"
	"halfway_demo/internal/service"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/viper"
	"github.com/xulichen/halfway/pkg/log"

	"github.com/xulichen/halfway/pkg/config"
	"go.elastic.co/apm"
)

func main() {
	config.InitViper("./config/config.yaml")
	viper.SetDefault("name", "halfway_demo")
	apm.DefaultTracer.Service.Name = "halfway_demo"
	apm.DefaultTracer.Service.Environment = viper.GetString("golang-base.env")
	database := dao.NewMysql()
	client, err := dao.NewRedis()
	if err != nil {
		panic(fmt.Sprintf("New redis client error, err: %s", err.Error()))
	}
	d := dao.New(database, client)
	serv := service.New(d)

	rpcSrv := server.NewRpcServer()
	routers.RegisterGrpcRouters(rpcSrv, serv)
	if err := rpcSrv.Start(); err != nil {
		panic("start rpc server error")
	}
	httpSrv := server.NewHttpServer()
	routers.RegisterHttpRouters(httpSrv, serv)
	httpSrv.Start()

	stopSignal := make(chan os.Signal, 1)
	signal.Notify(stopSignal, os.Interrupt, syscall.SIGTERM)

	select {
	case <-stopSignal:
		log.GetLogger().Info("server GracefulStop start")
		_ = httpSrv.GracefulStop()
		rpcSrv.GracefulStop()
		log.GetLogger().Info("server GracefulStop finish")
	}
}
