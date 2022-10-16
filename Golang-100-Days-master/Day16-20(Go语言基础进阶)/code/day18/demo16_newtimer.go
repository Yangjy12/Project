package main

import (
	"fmt"
	"time"
)

func main() {
	/*
	1. func NewTimer(d Duration) *Timer
			创建一个计时器，d时间以后触发


	 */
	 //timer := time.NewTimer(3 *time.Second)//这里数字就是定义阻塞多少秒
	 //fmt.Printf("%T\n",timer) //*time.Timer
	 //fmt.Println(time.Now()) //2022-07-31 08:08:38.3434183 +0800 CST m=+0.002415501
	 ////
	 //////此处等待channel中的数值，会阻塞3秒
	 //ch2 := timer.C
	 //fmt.Println(<-ch2 ) //2022-07-31 08:09:56.1214098 +0800 CST m=+3.011776701


	 //新建一个计时器
	 timer2 := time.NewTimer(5*time.Second)
	 //开始goroutine，来处理触发后的事件
	 go func() {
	 	<- timer2.C
	 	fmt.Println("Timer 2 结束了。。。开始。。。。")
	 }()
	//由于上面的等待信号是在新线程中，所以代码会继续往下执行，停掉计时器
	 time.Sleep(3*time.Second)
	 flag := timer2.Stop()
	 if flag{
	 	fmt.Println("Timer 2 停止了。。。")
	 }
}
