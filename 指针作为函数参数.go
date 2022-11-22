package main

import "fmt"

// 指针作为函数参数
func pointParm() {
	// 向函数传递指针，并在函数调用后修改函数内的值
	/* 定义局部变量 */
	var a int = 100
	var b int = 200

	fmt.Printf("交换前 a 的值 : %d\n", a)
	fmt.Printf("交换前 b 的值 : %d\n", b)

	/* 调用函数用于交换值
	 * &a 指向 a 变量的地址
	 * &b 指向 b 变量的地址
	 */
	swap_pointer(&a, &b)

	fmt.Printf("交换后 a 的值 : %d\n", a)
	fmt.Printf("交换后 b 的值 : %d\n", b)
	/**

	交换前 a 的值 : 100
	交换前 b 的值 : 200
	交换后 a 的值 : 200
	交换后 b 的值 : 100
	*/

}
func swap_pointer(x *int, y *int) {
	var temp int
	temp = *x /* 保存 x 地址的值 */
	*x = *y   /* 将 y 赋值给 x */
	*y = temp /* 将 temp 赋值给 y */
}

// 更简洁的变量交换实例
func moreConciseSwap() {
	var a int = 100
	var b int = 200
	mcSwap(&a, &b)

	fmt.Printf("交换后 a 的值 : %d\n", a)
	fmt.Printf("交换后 b 的值 : %d\n", b)
	/***
	200
	100


	*/

}

func mcSwap(x *int, y *int) {
	*x, *y = *y, *x
}

// 更更简洁的变量交换实例
func superMoreConciseSwap() {
	var a int = 19
	var b int = 3
	a, b = b, a
	fmt.Printf("交换后 a 的值 : %d\n", a) //3
	fmt.Printf("交换后 b 的值 : %d\n", b) //19
}
