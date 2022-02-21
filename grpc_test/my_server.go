package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"pb"
)

type Children struct {
	*pb.UnimplementedSayNameServer
}

func (this *Children) SayHello(ctx context.Context, t *pb.Teacher) (*pb.Teacher, error) {
	t.Name += "is sleeping"
	return t, nil
}

func main() {
	grpcServer := grpc.NewServer()

	pb.RegisterSayNameServer(grpcServer, new(Children))
	listener, err := net.Listen("tcp", "127.0.0.1:8800")
	if err != nil {
		fmt.Println("Listen err: ", err)
		return
	}
	defer listener.Close()

	fmt.Println("启动 server!")
	grpcServer.Serve(listener)
}
