package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)
//定义一个全局对象db
var  db  *sql.DB

func initDB() (err error) {
	dsn:="root:123456@tcp(127.0.0.1:3306)/go?charset=utf8mb4&parseTime=True"
	//注意这里不能用：=，我们是给全局变量赋值，然后在main函数中使用全局变量db
	db,err=sql.Open("mysql",dsn)
	if err!=nil {
		return err
	}
	//尝试与数据库建立链接（校验dsn是否正确）
	err=db.Ping()
	if err!=nil {
		return err
	}
	return nil
}
func insert()  {
	s := "insert into users (username,password) values(?,?)"
	r, err := db.Exec(s, "yjy", "123456")
	if err!=nil {
		println(err)
	}else {
		id, _ := r.LastInsertId()//默认情况下他是返回 主键且自增长的字段值，否则返回0
		println(id)
	}
}
func insert2(username string,password string)  {
	s := "insert into users (username,password) values(?,?)"
	r, err := db.Exec(s, username, password)
	if err!=nil {
		println(err)
	}else {
		id, _ := r.LastInsertId()
		println(id)
	}
}

type User struct {
	id int
	username string
	password string
}
func queryOneRow()  {
	s := "select * from users where id=?"
	var u User
	err := db.QueryRow(s, 3).Scan(&u.username, &u.password, &u.id) //scan就是将查询的结果赋值给另外一个结果,注意写地址
	if err!=nil {
		println(err)
	}else {
		fmt.Println(u)
	}
}
func queryManyRow()  {
	s := "select * from users"
	query, err := db.Query(s)
	var u User
	defer query.Close()
	if err!=nil{
		print(err)
	}else {
		for query.Next(){
			err := query.Scan(&u.username, &u.password, &u.id)
			if err!=nil {
				println(err)
			}else {
				fmt.Println(u)
			}
		}
	}
}
func update()  {
	s := "update users set username=?,password=? where id=?"
	exec, err := db.Exec(s, "yes", "123456", 2)
	if err!=nil {
		println(err)
	}else {
		affected, _ := exec.RowsAffected()//改变了几行
		fmt.Println(affected)
	}
}
func delete(id int)  {
	s := "delete from users where id=?"
	exec, err := db.Exec(s, id)
	if err!=nil{
		fmt.Println(err)
	}else {
		affected, _ := exec.RowsAffected()
		fmt.Println(affected)
	}
}
func main()  {
	fmt.Println("hellp")
	////定义一个全局变量DB
	db, err := sql.Open("mysql", "root:123456@/go")
	if err != nil {
		fmt.Println("cuowu ")
	}else {
		fmt.Println("chengg")
	}
	//// See "Important settings" section.
	////最大连接时长
	db.SetConnMaxLifetime(time.Minute * 3)
	//最大连接数
	db.SetMaxOpenConns(10)
	//空闲连接数
	db.SetMaxIdleConns(10)
	//
	println(db)
	err = initDB()
	if err!=nil {
		fmt.Println(err)
	}else {
		fmt.Println("连接成功！")
	}
	insert()
	//insert2("qxx","123")
	//queryOneRow()
	//queryManyRow()
	//update()
	//delete(2)
}
