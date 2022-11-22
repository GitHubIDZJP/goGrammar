package main

import "fmt"

const MAX int = 3

// 正常数组
func Arr() {
	a := []int{10, 100, 200}
	var i int

	for i = 0; i < MAX; i++ {
		fmt.Printf("a[%d] = %d\n", i, a[i])
	}
	/**
	   a[0] = 10
	a[1] = 100
	a[2] = 200
	*/
}

const MAX_P int = 3

// 指针数组
func pointerArr() {
	//var ptr [MAX]*int;
	//ptr 为整型指针数组。因此每个元素都指向了一个值。以下实例的三个整数将存储在指针数组中
	a := []int{10, 100, 200}
	var i int
	var ptr [MAX_P]*int //新加
	//新加
	for i = 0; i < MAX_P; i++ {
		ptr[i] = &a[i] /* 整数地址赋值给指针数组 */
	}

	for i = 0; i < MAX_P; i++ {
		fmt.Printf("a[%d] = %d\n", i, *ptr[i])
	}
	/**
	 a[0] = 10
	a[1] = 100
	a[2] = 200

	*/
}
