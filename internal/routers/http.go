package routers

import (
	"halfway_demo/api/demo"
	"halfway_demo/internal/service"

	"github.com/labstack/echo/v4"
	"github.com/xulichen/halfway/pkg/net/http"
)

func RegisterHttpRouters(server *http.Server, srv *service.Service) {
	registerDemoRouter(server, &service.HelloService{Service: srv})
}

func registerDemoRouter(srv *http.Server, server demo.DemoServer) {
	srv.GET("/health", health)
	srv.GET("/hello", http.GRPCProxyWrapper(server.SayHello))
}

func health(ctx echo.Context) error {
	return ctx.String(200, "ok")
}
