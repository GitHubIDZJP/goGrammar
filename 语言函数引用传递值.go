package main

//引用传递值
import "fmt"

//引用传递是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数。
//引用传递指针参数传递到函数内，以下是交换函数 swap() 使用了引用传递

func swapMain() {

	// 调用
	var aa int = 100
	var bb int = 20
	a_swap(&aa, &bb)
	fmt.Print("交换后的aa的值:\n", aa) //20
	fmt.Print("交换后的bb的值:\n", bb) //100

}

func a_swap(x *int, y *int) {

	var temp int
	temp = *x /* 保持 x 地址上的值 */
	*x = *y   /* 将 y 值赋给 x */
	*y = temp /* 将 temp 值赋给 y */
}
