package main

import (
	"context"
	"fmt"

	bj38 "bj38Web/proto"

	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3"
	"github.com/gin-gonic/gin"
)

func CallRemote(c *gin.Context) {
	//初始化服务发现
	reg := consul.NewRegistry()
	server := micro.NewService(
		micro.Registry(reg),
	)
	//修改初始化的客户端
	microClient := bj38.NewBj38Service("bj38", server.Client())
	//1.初始化客户端
	// microClient := bj38.NewBj38Service("bj38", client.DefaultClient)
	resp, err := microClient.Call(context.TODO(), &bj38.Request{
		Name: "xiaowang",
	})
	if err != nil {
		fmt.Println("call err:", err)
		return
	}

	//为了方便查看，在打印之前将结果返回浏览器
	c.Writer.WriteString(resp.Msg)
	fmt.Println(resp, err)
}

func main() {
	//1.初始化路由  官网：初始化web引擎
	router := gin.Default()

	//2.作路由匹配
	// router.GET("/", func(context *gin.Context) {
	// 	context.Writer.WriteString("hello world")
	// })
	router.GET("/", CallRemote)

	//3.启动运行
	router.Run(":8080")
}
