package main

import (
	"fmt"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	store, _ := sessions.NewRedisStore(10, "tcp", "localhost:6379", "", []byte("bj38"))
	router.Use(sessions.Sessions("session", store))
	router.GET("/test", func(context *gin.Context) {
		//设置cookie
		// context.SetCookie("mytest", "itcast", 60*60, "", "", false, true)
		//获取cookie
		// cookieVal, _ := context.Cookie("mytest")
		// fmt.Println("获取的cookie为：", cookieVal)
		// context.Writer.WriteString("测试cookie。。。")

		session := sessions.Default(context)
		// var count int
		// v := session.Get("count")
		// if v == nil {
		//   count = 0
		// } else {
		//   count = v.(int)
		//   count += 1
		// }

		//设置session
		// session.Set("itcast", "itheima")
		// session.Save()
		// // c.JSON(200, gin.H{"count": count})

		v := session.Get("itcast")
		fmt.Println("获取session: ", v.(string))

		context.Writer.WriteString("测试session。。。")
	})
	router.Run(":9999")
}

// import (
// 	"github.com/gin-gonic/contrib/sessions"
// 	"github.com/gin-gonic/gin"
//   )
