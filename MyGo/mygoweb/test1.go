package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"runtime"
)

func getCurrentPath() string {
	_, file, _, _ := runtime.Caller(1)
	return path.Dir(file)
}
func hello(w http.ResponseWriter,r *http.Request){
	currentPath := getCurrentPath()
	fmt.Fprintln(w,currentPath)
	//file, _ := ioutil.ReadFile("D:\\MyGo\\mygoweb\\hello.txt")
	file, _ := ioutil.ReadFile("./mygoweb/hello.txt")
	//file, _ := ioutil.ReadFile("./hello.txt")
	fmt.Fprintln(w,string(file))
	fmt.Fprintln(w,"yes")
}
func main() {
	http.HandleFunc("/hi",hello)
	err := http.ListenAndServe(":9090", nil)
	if err!=nil {
		fmt.Printf("http server failed,err:%v\n",err)
		return
	}
}
