package main

import (
	"fmt"

	"github.com/tedcy/fdfs_client"
)

func main() {
	client, err := fdfs_client.NewClientWithConfig("/etc/fdfs/client.conf")
	defer client.Destory()
	if err != nil {
		fmt.Println("初始化客户端错误", err)
		return
	}

	//上传文件  --传入到storage
	resp, err := client.UploadByFilename("头像.jpeg")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp, err)
	// if fileId, err := client.UploadByBuffer([]byte("hello world"), "go"); err != nil {
	// 	fmt.Println(err.Error())
	// }
}
