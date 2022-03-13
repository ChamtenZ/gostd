package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	//链接数据库
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis dial err:", err)
		return
	}
	defer conn.Close()

	//操作数据库
	reply, err := conn.Do("set", "itcast", "itheima")

	//回复助手函数-----确定成具体的数据类型
	r, e := redis.String(reply, err)
	fmt.Println(r, e)
}
