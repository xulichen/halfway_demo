package routers

import (
	"halfway_demo/api/demo"
	"halfway_demo/internal/service"

	"github.com/xulichen/halfway/pkg/net/rpc"
)

func RegisterGrpcRouters(server *rpc.Server, srv *service.Service) {
	// TODO: HelloService http grpc 只实例化一次
	demo.RegisterDemoServer(server.Server, &service.HelloService{Service: srv})
}
