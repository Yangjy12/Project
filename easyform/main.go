package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./Login.html")
	r.GET("/login", func(context *gin.Context) {
		context.HTML(http.StatusOK,"Login.html",nil)
	})
	r.POST("/userlogin", func(context *gin.Context) {
		//name := context.PostForm("name")//取到就返回值，取不到就返回空字符串
		//pass := context.PostForm("password")
		//name:= context.DefaultPostForm("name", "用户名未输入")
		pass := context.DefaultPostForm("password", "*******")
		name, ok := context.GetPostForm("name")
		if !ok {
			name="请输入正确的应户名！"
		}
		context.JSON(http.StatusOK,gin.H{
			"Name":name,
			"Password":pass,
		})
	})
	//获得path路径,所有返回值都是string类型。注意url的匹配不要冲突
	r.GET("/username/:name/:age", func(context *gin.Context) {
		name := context.Param("name")
		age := context.Param("age")
		context.JSON(http.StatusOK,gin.H{
			"Name":name,
			"Age":age,
		})
	})
	type userinfo struct {
		Name string `form:"username" json:"user"`
		Password string `form:"password" json:"pass"`
	}
	//使用c.ShouldBind()完成数据绑定
	r.GET("/form", func(context *gin.Context) {
		var u userinfo
		err := context.ShouldBind(&u)//注意这里要取地址
		if err!=nil {
			context.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
		}else {
			fmt.Printf("%v\n",u)
			context.JSON(http.StatusOK,gin.H{
				"message":"ok",
			})
		}
	})
	/*文件上传*/
	r.LoadHTMLFiles("./index.html")
	r.GET("/index", func(context *gin.Context) {
		context.HTML(200,"index.html",nil)
	})
	r.POST("/upload", func(context *gin.Context) {
		//从请求中读取文件
		file, err := context.FormFile("f1")//从请求中获取携带的参数一样的
		if err!=nil {
			context.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
		}else {
			//把读取到的文件保存到本地
			dst:= fmt.Sprintf("./%s", file.Filename)
			//或者这样写
			//path.Join("./",file.Filename)
			context.SaveUploadedFile(file,dst)

		}
	})
	/*重定向两种方法*/
	r.GET("/a", func(context *gin.Context) {
		/*方法一，这种方法，地址栏会发生变化*/
		context.Redirect(http.StatusMovedPermanently,"www.baidu.com")
		/*方法二，这种方法地址栏不会发声变化*/
		context.Request.URL.Path="/b"//把请求的url修改
		r.HandleContext(context)//继续后续的处理
	})
	r.GET("/b", func(context *gin.Context) {
		context.JSON(http.StatusOK,gin.H{
			"message":"ok",
		})
	})
	r.Run(":8080")
}
