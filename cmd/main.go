package main

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"github.com/xulichen/halfway/pkg/log"
	"halfway_demo/internal/dao"
	"halfway_demo/internal/server"
	"halfway_demo/internal/service"
	"os"
	"os/signal"
	"syscall"

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
	helloService := &service.HelloService{Service: serv}

	rpcSrv := server.NewRpcServer()
	server.RegisterDemoService(rpcSrv, helloService)
	if err := rpcSrv.Start(); err != nil {
		panic("start rpc server error")
	}
	httpSrv := server.NewHttpServer()
	server.RegisterDemoRouter(httpSrv, helloService)
	httpSrv.Start()

	stopSignal := make(chan os.Signal, 1)
	signal.Notify(stopSignal, os.Interrupt, syscall.SIGTERM)

	select {
		case <- stopSignal:
			log.GetLogger().Info("server GracefulStop start")
			_ = httpSrv.GracefulStop(context.Background())
			rpcSrv.GracefulStop()
			log.GetLogger().Info("server GracefulStop finish")
	}
}
