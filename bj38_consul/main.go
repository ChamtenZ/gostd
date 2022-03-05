package main

import (
	"bj38/handler"
	pb "bj38/proto"

	// "github.com/micro/micro/v3/service"
	// "github.com/micro/micro/v3/service/logger"

	"github.com/asim/go-micro/plugins/registry/consul/v3"
	service "github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/logger"
)

func main() {
	//初始化服务发现
	// reg := consul.NewRegistry(func(options *registry.Options) {
	// 	options.Addrs = []string{
	// 		"127.0.0.1:8800",
	// 	}
	// })
	reg := consul.NewRegistry()
	// Create service
	srv := service.NewService(
		service.Name("bj38"),
		service.Registry(reg),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterBj38Handler(srv.Server(), new(handler.Bj38))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
