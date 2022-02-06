package service

import (
	"fmt"
	pb "grpc-demo-simple/api"

	"golang.org/x/net/context"
)

// 定义一个 Hello 服务并实现接口
type HelloService struct{}

func (h HelloService) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	fmt.Println("服务器收到：", req)
	resp := new(pb.HelloReply)
	resp.Message = "Hello，" + req.Name + " ！"
	return resp, nil
}
