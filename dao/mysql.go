package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)
func InitMySql() (err error) {
	DB,err=gorm.Open("mysql","root:123456@tcp(127.0.0.1:3306)/bubble?charset=utf8mb4&parseTime=True&loc=Local")
	if err!=nil {
		return
	}
	//err = DB.DB().Ping()
	return DB.DB().Ping()
}