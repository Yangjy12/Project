package main

import (
	"github.com/gin-gonic/gin"
)

func sayhello(c *gin.Context)  {
	c.JSON(200,gin.H{
		"message":"Hello golang!",
	})
}
func main() {
	// 创建一个默认的路由引擎
	engine := gin.Default()//返回的是默认引擎
	//get方式请求/sayhello路径时，执行
	engine.GET("/sayhello",sayhello)
	//启动服务
	engine.Run(":9090")//默认是8080
}
