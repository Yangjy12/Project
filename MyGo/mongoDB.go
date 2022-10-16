package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type Students struct {
	Name string
	Age int
}

var client *mongo.Client

func initDb()  {
	//设置客户端连接配置
	uri := options.Client().ApplyURI("mongodb://localhost:27017")
	//连接到MongoDB
	connect, err3 := mongo.Connect(context.TODO(), uri)
	if err3!=nil {
		log.Fatal(err3)
	}else {
		fmt.Println(connect)
	}
	//检验连接
	err2:= connect.Ping(context.TODO(), nil)
	if err2!=nil {
		log.Fatal(err2)
	}else {
		fmt.Println("连接成功！")
	}
	client=connect
}
func mongoinsert()  {
	initDb()
	s1:=Students{
		Name: "yjy",
		Age: 21,
	}
	//打开数据库并且打开Syudents的这个Collection集合
	collection:= client.Database("golang_db").Collection("Students")
	one, err := collection.InsertOne(context.TODO(), s1)
	if err!=nil {
		fmt.Println(err)
	}else {
		fmt.Println(one.InsertedID)
	}
}
func insertMany()  {
	initDb()
	c:= client.Database("golang_db").Collection("Students")
	s1:=Students{
		Name: "kitty",
		Age: 21,
	}
	s2:=Students{
		Name: "zqy",
		Age: 27,
	}
	stus:=[]interface{}{s1,s2}
	many, err := c.InsertMany(context.TODO(), stus)
	if err!=nil {
		fmt.Println(err)
	}else {
		fmt.Println(many.InsertedIDs)
	}
}
func find()  {
	initDb()
	defer client.Disconnect(context.TODO())
	c:= client.Database("golang_db").Collection("Students")
	cur, err := c.Find(context.TODO(), bson.D{{"name","zqy"}})
	//bson.D{}里面可写可不写，写了就是条件，不写就是查找全部
	if err!=nil {
		log.Fatal(err)
	}
	defer cur.Close(context.TODO())
	for cur.Next(context.TODO()) {
		var result bson.D
		cur.Decode(&result)
		fmt.Println(result.Map())
		fmt.Println(result)
	}
}
func mongoupdate()  {
	initDb()
	c:= client.Database("golang_db").Collection("Students")
	//将数据要改成
	d := bson.D{{"$set", bson.D{{"name", "qxx"}, {"age", "28"}}}}
	//要改的原数据
	many, err := c.UpdateMany(context.TODO(), bson.D{{"name", "kitty"}}, d)
	if err!=nil {
		log.Fatal(err)
	}
	//打印改变了多少
	fmt.Println(many.ModifiedCount)
}
func del()  {
	initDb()
	c:= client.Database("golang_db").Collection("Students")
	many, err := c.DeleteMany(context.TODO(), bson.D{{"name", "qxx"}})
	if err!=nil {
		log.Fatal(err)
	}
	fmt.Println(many.DeletedCount)
}
func main()  {
	//mongoinsert()
	//insertMany()
	//find()
	//mongoupdate()
	del()
}
