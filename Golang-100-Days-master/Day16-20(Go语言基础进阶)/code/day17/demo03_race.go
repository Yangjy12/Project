package main

import (
	"fmt"
	"time"
)

func main() {
	/*
	临界资源：
	 */
	 a := 1
	 go func() {
	 	a = 2
	 	fmt.Println("goroutine中。。",a)
	 }()
//如果多个goroutine在访问同一个数据资源的时候，其中一个线程修改了数据，那么这个数值就被修改了，
//对于其他的goroutine来讲，这个数值可能是不对的。
	 a = 3
	 time.Sleep(1)
	 fmt.Println("main goroutine...",a)
}