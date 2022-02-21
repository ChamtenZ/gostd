package main

import (
	"context"
	"fmt"
	"gostd/pb"
	"strconv"

	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
)

func main() {
	consulConfig := api.DefaultConfig()
	//创建consul对象（可以重新指定consul属性：IP/port)
	consulClient, err := api.NewClient(consulConfig)
	//服务发现，从consul上获取健康的服务
	services, _, err := consulClient.Health().Service("grpc And Consul", "grpc", true, nil)
	addr := services[0].Service.Address + ":" + strconv.Itoa(services[0].Service.Port)
	grpcConn, _ := grpc.Dial(addr, grpc.WithInsecure())

	// grpcConn, _ := grpc.Dial("127.0.0.1:8800", grpc.WithInsecure())
	defer grpcConn.Close()
	grpcClient := pb.NewSayNameClient(grpcConn)

	var teacher pb.Teacher
	teacher.Name = "Andy"
	teacher.Age = 18
	t, err := grpcClient.SayHello(context.TODO(), &teacher)
	fmt.Println(t, err)
}
