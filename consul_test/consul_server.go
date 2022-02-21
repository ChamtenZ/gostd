package main

import (
	"context"
	"fmt"
	"gostd/pb"
	"net"

	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
)

type Children struct {
	*pb.UnimplementedSayNameServer
}

func (this *Children) SayHello(ctx context.Context, p *pb.Teacher) (*pb.Teacher, error) {
	p.Name = "hello" + p.Name
	return p, nil
}

func main() {
	///////////////把grpc服务注册到consul////////////
	//1.初始化consul配置
	consulConfig := api.DefaultConfig()
	//2.创建consul对象
	consulClient, err := api.NewClient(consulConfig)
	if err != nil {
		fmt.Println("api.NewClient err:", err)
		return
	}
	//3.告诉consul，即将注册的服务的配置信息
	reg := api.AgentServiceRegistration{
		ID:      "bj38",
		Tags:    []string{"grpc", "consul"},
		Name:    "grpc And Consul",
		Address: "127.0.0.1",
		Port:    8800,
		Check: &api.AgentServiceCheck{
			CheckID:  "consul grpc test",
			TCP:      "127.0.0.1:8800",
			Timeout:  "1s",
			Interval: "5s",
		},
	}
	//4.注册grpc服务到consul上
	consulClient.Agent().ServiceRegister(&reg)

	////////////////////grpc服务 远程调用/////////////
	grpcServer := grpc.NewServer()
	pb.RegisterSayNameServer(grpcServer, new(Children))
	listener, err := net.Listen("tcp", "127.0.0.1:8800")
	if err != nil {
		fmt.Println("Listen err:", err)
		return
	}
	defer listener.Close()
	grpcServer.Serve(listener)
}
