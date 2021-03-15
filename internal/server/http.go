package server

import (
	demo "halfway_demo/proto"

	"github.com/labstack/echo/v4"
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

func RegisterDemoRouter(srv *http.Server, server demo.DemoServer) {
	srv.GET("/health", health)
	srv.GET("/hello", http.GRPCProxyWrapper(server.SayHello))
}

func health(ctx echo.Context) error {
	return ctx.String(200, "ok")
}
