package service

import (
	"context"
	demo "halfway_demo/proto"
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
