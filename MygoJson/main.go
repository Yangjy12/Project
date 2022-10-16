package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
//控制台输入go mod tidy会自动识别所缺少的包，自动导入
func main() {
	r:=gin.Default()
	// gin.H 是map[string]interface{}的缩写
	r.GET("/someJSON", func(c *gin.Context) {
		// 方式一：自己拼接JSON
		c.JSON(http.StatusOK, gin.H{"message": "Hello world!"})
	})
	r.GET("/moreJSON", func(c *gin.Context) {
		// 方法二：使用结构体
		type msg struct {
			Name    string `json:"user"`//小写默认不可导出
			Message string
			Age     int
		}
		data := msg{
			"xiaowangzi ",
			"hello",
			19,
		}
		c.JSON(http.StatusOK, data)
		/*var msg struct{

		}
		msg.Name = "小王子"
		msg.Message = "Hello world!"
		msg.Age = 18
		c.JSON(http.StatusOK, msg)*/
	})
	/*127.0.0:8080/web?query=杨婧媛*/
	r.GET("/web", func(context *gin.Context) {
		query := context.Query("query")
		context.DefaultQuery("query","yjy")//取不到就用默认值
		context.JSON(http.StatusOK,gin.H{
			"name":query,
		})
	})

	r.Run(":8080")
}
