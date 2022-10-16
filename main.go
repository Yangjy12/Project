package main

import (
	"bubble/Models"
	"bubble/dao"
	"bubble/routers"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	//连接数据库
	err := dao.InitMySql()
	if err!=nil {
		fmt.Println(err)
	}
	defer dao.DB.Close()
	//模型绑定
	dao.DB.AutoMigrate(&Models.Todo{})
	r:=routers.SetupRouter()
	r.Run(":8080")
}
