package main

import "fmt"

//是用来传递数据的一个数据结构。
//通道可用于两个 goroutine 之间通过传递一个指定类型的值来同步运行和通讯。
//操作符 <- 用于指定通道的方向，发送或接收。如果未指定方向，则为双向通道。
/**
ch <- v     把 v 发送到通道 ch
v := <-ch   从 ch 接收数据
            并把值赋给 v
*/
//通道
func channelMain() {
	//默认情况下，通道是不带缓冲区的。发送端发送数据，同时必须有接收端相应的接收数据。
	//以下实例通过两个 goroutine 来计算数字之和，在 goroutine 完成计算后，它会计算两个结果的和
	//var numbers = make([]int, 3, 5)
	//numbersFirst := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	s := []int{7, 2, 8, -9, 4, 0} //定义切片

	c := make(chan int) //

	var sliceValue = s[:len(s)]
	fmt.Println("切片值:\n", sliceValue) //[7 2 8 -9 4 0]

	var sliceValue1 = s[len(s)/2:]
	fmt.Println("切片值1:\n", sliceValue1) //[-9 4 0]

	//go：并发
	go sum(s[:len(s)/2], c) //切片: a:取的是从len(s)/2后面的index位其实质
	go sum(s[len(s)/2:], c) //切片获取3位， -9,4,0
	x, y := <-c, <-c        // 从通道 c 中接收
	
	fmt.Println(x, y, x+y)
	/*
		结果:
		x: -5
		y: 17
		x+y: 12

	*/

}

// chan:通道-》用来传值的
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 把 sum 发送到通道 c
}
