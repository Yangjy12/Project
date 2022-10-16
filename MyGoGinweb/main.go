package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func main() {
	r := gin.Default()
	//加载静态文件
	r.Static("/xxx","./statics")
	//这是自定义的一个函数模板，funcmap用来存放模板，safe是模板的名字，然后：是内容为一个匿名函数。
	r.SetFuncMap(template.FuncMap{
		"safe": func(string2 string) template.HTML {
			return template.HTML(string2)
		},
	})
	//r.LoadHTMLGlob("./templates/**/*")
	//r.LoadHTMLFiles("./templates/posts/index.html", "./templates/users/index.html")
	r.LoadHTMLFiles("./templates/index.tmpl")
	r.GET("/index", func(context *gin.Context) {
		context.HTML(200,"index.tmpl","<a href='https://liwenzhou.com'>李文周的博客</a>")
	})
	//r.GET("/posts/index", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "index.html", gin.H{
	//		"title": "posts/index",
	//	})
	//})
	//
	//r.GET("/users/index", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "index.html", gin.H{
	//		"title": "users/index",
	//	})
	//})
	r.LoadHTMLFiles("./templates/index.html")
	r.GET("/home", func(context *gin.Context) {
		context.HTML(http.StatusOK,"index.html",nil)
	})
	r.Run(":9090")
}
