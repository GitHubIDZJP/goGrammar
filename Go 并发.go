package main

import (
	"fmt"
	"time"
)

/**


Go 语言支持并发，我们只需要通过 go 关键字来开启 goroutine 即可。
goroutine 是轻量级线程，goroutine 的调度是由 Golang 运行时进行管理的。
goroutine 语法格式：
go 函数名( 参数列表 )
例如：
go f(x, y, z)
开启一个新的 goroutine:
f(x, y, z)

*/
//并发
func concurrenceMain() {

	go say("world") //涉及并发
	say("hello")
	/*
	   输出的 hello 和 world 是没有固定先后顺序。因为它们是两个 goroutine 在执行
	*/
}
func say(s string) {
	//Microsecond微妙
	for i := 0; i < 5; i++ {
		time.Sleep(10 * time.Microsecond)
		fmt.Println(s)
	}
}
