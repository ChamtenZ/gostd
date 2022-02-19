package main

import (
	"context"
	"fmt"
	"pb"

	"google.golang.org/grpc"
)

func main() {
	//1.连接grpc服务
	grpcConn, err := grpc.Dial("127.0.0.1:8800", grpc.WithInsecure())
	if err != nil {
		fmt.Println("grpc Dial err:", err)
		return
	}
	defer grpcConn.Close()
	//2.初始化grpc客户端
	grpcClient := pb.NewSayNameClient(grpcConn)
	//3.调用远程服务
	var teacher pb.Teacher
	teacher.Name = "itcast"
	teacher.Age = 18
	t, err := grpcClient.SayHello(context.TODO(), &teacher)
	fmt.Println(t)
}
