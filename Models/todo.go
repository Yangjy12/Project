package Models

import (
	"bubble/dao"
)

type Todo struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Status bool `json:"status"`
}
/*
Todo的这些增删查改都放到这里面
*
 */

func CreatTodo(todo *Todo) (err error) {
	if err2 := dao.DB.Create(&todo).Error;err2!=nil{
		return err
	}
	return
}
func GetAllTodo()(todolist []*Todo,err error)  {
	if err:=dao.DB.Find(&todolist).Error; err != nil{
		return nil, err
	}
	return
}
func GetTodo(id string)(todo *Todo,err error)  {
	todo=new(Todo)
	if err:=dao.DB.Where("id=?",id).First(&todo).Error;err!=nil {
		return nil,err
	}
	return 
}
func UpdateTodo(todo *Todo)(err error)  {
	if err=dao.DB.Save(todo).Error;err!=nil{
		return err
	}
	return
}
func DeleteTodo(id string)(err error)  {
	err=dao.DB.Where("id=?",id).Delete(&Todo{}).Error
	return
}