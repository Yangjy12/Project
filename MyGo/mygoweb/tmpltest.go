package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func sayhi(w http.ResponseWriter,r *http.Request)  {
	//解析模板
	files, err := template.ParseFiles("./mygoweb/hello.tmpl")
	if err!=nil {
		fmt.Println("Parse template failed,err:%v\n",err)
		return
	}
	name:="yjy"
	err = files.Execute(w, name)
	if err!=nil {
		fmt.Println("render template faild,err:%v\n",err)
		return
	}
}
func main() {
	http.HandleFunc("/",sayhi)
	err := http.ListenAndServe(":9090", nil)
	if err!=nil {
		fmt.Println("Http server stort faild,err:%v\n",err)
		return
	}
}
