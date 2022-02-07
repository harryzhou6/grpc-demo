package main

import (
	"fmt"
	pb "grpc-demo-simple/api"
	"grpc-demo-simple/server/service"
	"net"

	"google.golang.org/grpc/reflection"

	"google.golang.org/grpc"
	_ "google.golang.org/grpc/encoding/gzip"
)

const (
	// gRPC 服务地址
	Address = "127.0.0.1:54321"
)

func main() {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
	}
	// 实现 grpc Server
	s := grpc.NewServer()
	// reflection包提供的反射服务查询gRPC服务或调用gRPC方法
	reflection.Register(s)
	// 注册 service.HelloService 为客户端提供服务
	pb.RegisterHelloServer(s, service.HelloService{})
	fmt.Println("Listen on" + Address)

	_ = s.Serve(listen)
}
