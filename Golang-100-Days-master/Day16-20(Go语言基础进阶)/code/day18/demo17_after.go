package main

import (
	"fmt"
	"time"
)

func main() {
	/*
	 2. func After(d Duration) <-chan Time
			返回一个通道：chan，存储的是d时间间隔之后的当前时间

		相当于：return NewTimer(d).C
	 */
	ch :=  time.After(3 *time.Second)
	fmt.Printf("%T\n",ch) //<-chan time.Time
	fmt.Println(time.Now()) //2022-07-31 08:14:21.8453578 +0800 CST m=+0.002258801


	time2 := <-ch
	fmt.Println(time2) //2022-07-31 08:14:24.8462755 +0800 CST m=+3.003176501
}
