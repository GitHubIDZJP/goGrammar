package main

import "fmt"

// 语言函数
func lanuageFunction() {
	/*
		 函数定于
		 func 函数名([参数列表])  [返回类型]{
			函数体
		 }


	*/

	var callValue int = max(200, 20)
	fmt.Printf("最大值是 : %d\n", callValue)

	var a, b = swap("傻逼", "大傻逼")
	fmt.Print("交换位置  : \n", a, b)

	var j1 int = 90
	var j2 int = 200
	swapValue(j1, j2)
	//fmt.Printf("最大值是 : %d\n", j1, j2)
	fmt.Printf("交换后 j1 的值 : %d\n", j1)
	fmt.Printf("交换后 j2 的值 : %d\n", j2)
}

/*
函数名:max
参数: num1 .num2 (int：是参数类型)
int:函数返回类型
*/
func max(num1, num2 int) int {
	var result int
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

func swap(x, y string) (string, string) {
	return y, x
}

// 值传递:交换常量的值
func swapValue(j1, j2 int) int {
	var temp int
	temp = j1
	j1 = j2
	j2 = temp
	return temp
}
