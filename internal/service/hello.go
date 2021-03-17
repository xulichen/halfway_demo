package service

import (
	"context"
	"halfway_demo/api/demo"
)

type HelloService struct {
	//demo.UnimplementedDemoServer
	*Service
}

// SayHello demo.DemoServer func
func (hs *HelloService) SayHello(ctx context.Context, req *demo.HelloReq) (reply *demo.HelloResp, err error) {
	reply = &demo.HelloResp{
		Content: "hello " + req.Content,
	}
	return
}
