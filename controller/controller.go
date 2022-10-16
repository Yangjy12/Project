package controller

import (
	"bubble/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexHandle(context * gin.Context)  {
	context.HTML(http.StatusOK,"index.html",nil)
}
func Creat(context *gin.Context)  {
	//从前端页面填写待办事项
	//1.从请求中把数据拿出来
	var todo Models.Todo
	context.BindJSON(&todo)
	//2、存入数据库
	err:= Models.CreatTodo(&todo)
	if err!=nil{
		context.JSON(http.StatusOK,gin.H{"error":err.Error()})
	}else {
		context.JSON(http.StatusOK,todo)
	}
	//3、返回响应
}
func Search(context *gin.Context)  {
	//查询todo这个表里的所有数据
	todo, err2 := Models.GetAllTodo()
	if err2!=nil{
		context.JSON(http.StatusOK,gin.H{"error":err2.Error()})
	}else {
		context.JSON(http.StatusOK,todo)
	}
}
func Update(context *gin.Context)  {
	id, ok := context.Params.Get("id")
	if !ok {
		context.JSON(http.StatusOK,gin.H{"error":"无效的id"})
		return
	}
	todo, err := Models.GetTodo(id)
	if err!=nil {
		context.JSON(http.StatusOK,gin.H{"error":err.Error()})
		return
	}
	context.BindJSON(&todo)
	if err:=Models.UpdateTodo(todo);err!=nil {
		context.JSON(http.StatusOK,gin.H{"error":err.Error()})
	}else {
		context.JSON(http.StatusOK,todo)
	}
}
func Deleted(context *gin.Context)  {
	id, ok := context.Params.Get("id")
	if !ok {
		context.JSON(http.StatusOK,gin.H{"error":"无效的id"})
		return
	}
	if err:=Models.DeleteTodo(id);err!=nil{
		context.JSON(http.StatusOK,gin.H{"error":err.Error()})
	}else {
		context.JSON(http.StatusOK,gin.H{id:"deleted"})
	}
}