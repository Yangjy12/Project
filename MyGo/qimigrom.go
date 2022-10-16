package main

import (
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

type user struct {
	ID int
	Name string
	Age int
}
func main() {
	open, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/go?charset=utf8mb4&parseTime=True&loc=Local")
	if err!=nil {
		//fmt.Println(err.Error())、
		log.Fatal(err)
	}
	defer open.Close()
	//把模型与数据库中的表对应起来
	open.AutoMigrate(&User{})
}
