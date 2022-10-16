package routers

import (
	"bubble/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter()*gin.Engine  {
	r:= gin.Default()
	//告诉gin框架去哪里找css文件，所以要加载静态文件
	r.Static("./static","static")
	//告诉gin框架去哪里找模板文件
	r.LoadHTMLGlob("./templates/*")
	r.GET("/", controller.IndexHandle)
	v1group := r.Group("v1")
	{
		//待办事项
		//添加
		v1group.POST("/todo",controller.Creat)
		//查看所有的待办事项
		v1group.GET("/todo", controller.Search)
		//查看摸一个
		v1group.GET("/todo/:id", func(context *gin.Context) {

		})
		//修改
		v1group.PUT("/todo/:id", controller.Update)
		//删除
		v1group.DELETE("/todo/:id",controller.Deleted)
	}
	return r
}
