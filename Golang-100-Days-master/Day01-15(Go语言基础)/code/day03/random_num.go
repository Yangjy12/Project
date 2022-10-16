package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	/*
	生成随机数random：
		伪随机数，根据一定的算法公式算出来的。
	math/rand
	 */

	 //num1 := rand.Intn(20)
	 //fmt.Println(num1)//第二次跟第一次生成的随机数居然是一样的！
	 //
	 //rand.Seed(100)
	 //num2:=rand.Intn(20)
	 //fmt.Println(num2)

	 //t1:=time.Now()
	 //fmt.Println(t1)
	 //fmt.Printf("%T\n",t1)
	 //fmt.Println(t1.Unix())//秒
	 //
	 //timeStamp2:=t1.UnixNano()
	 //fmt.Println(timeStamp2)

	 //step1：设置种子数，可以设置成时间戳
	 rand.Seed(time.Now().UnixNano())
	 //for i:=0;i<10;i++{
	 //	//step2:调用生成随机数的函数
	 //	fmt.Println("-->",rand.Intn(100))
	 //}
	 ///*
	 //[15,76]
	 //	[0,61]+15
	 //[3,48]
	 //	[0,45]+3
	 //
	 //Intn(n) // [0,n)
	 // */
	 //num3:=rand.Intn(46)+3//[3,48]
	 //fmt.Println(num3)
	 //num4:=rand.Intn(62)+15 //[15,76]
	 //fmt.Println(num4)
	source :=rand.New(rand.NewSource(time.Now().Unix()))
	i:=source.Intn(12)
	fmt.Println(i)

}
