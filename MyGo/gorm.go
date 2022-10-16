package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Products struct {
	gorm.Model//继承
	Code string
	Price string
}

func main()  {
	dsn:="root:123456@tcp(127.0.0.1:3306)/go?charset=utf8mb4&parseTime=True&loc=Local"
	db,err:=gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err!=nil {
		fmt.Println(err)
	}
	//创建表
	//db.AutoMigrate(&Products{})
	//插入数据
	//products := Products{
	//	Code:  "1001",
	//	Price: "100",
	//}
	//db.Create(&products)
	//查询
	var p Products
	//db.First(&p,1)//1代表主键的意思,根据整形主键查找
	//db.First(&p,"Code=?","1001")//查找code字段值为1001的记录
	//fmt.Println(p)
	//更新（更新前先查找）
	db.First(&p,1)
	db.Model(&p).Update("Price",500)
	db.Model(&p).Updates(Products{Price: "900", Code: "8908"})
	//删除（删除前先查询）
	db.Delete(&p)//软删除，只增加删除标记
}