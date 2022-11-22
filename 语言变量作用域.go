package main

import "fmt"

/***

作用域为已声明标识符所表示的常量、类型、变量、函数或包在源代码中的作用范围。
Go 语言中变量可以在三个地方声明：
函数内定义的变量称为局部变量
函数外定义的变量称为全局变量
函数定义中的变量称为形式参数

局部变量
全局变量
形式参数
*/

// 函数使用了局部变量 a, b, c：
func localVariableDef() {

	/* 声明局部变量 */
	var a, b, c int

	/* 初始化参数 */
	a = 10
	b = 20
	c = a + b

	fmt.Printf("结果： a = %d, b = %d and c = %d\n", a, b, c)
	//a= 10
	// b = 20
	c = 30

}

// 声明全局变量
var g int

func invoking() {

	//声明局部变量
	var l_a, l_b int
	l_a = 1
	l_b = 11
	g = l_a + l_b
	fmt.Printf("结果： l_a = %d, l_b = %d and g = %d\n", l_a, l_b, g) //1 11 12

}

var g_g int = 20

// 全局级别调用
func Global_level_Invocation() {

	//局部
	var g_g int = 21

	fmt.Print("打印全局gg纸:\n", g_g) // 21

}

//形式参数：函数的局部变量来使用

var g_p int = 23

// 调用形参
func CallParameter() {
	var l_s int = 9
	var l_s_s int = 8
	var f_sum int = 0

	f_sum = fetchSum(l_s, l_s_s)
	fmt.Print("值:\n", f_sum) //17

	fmt.Print("\n\n")

}
func fetchSum(xp_a, xp_b int) int {
	return xp_a + xp_b
}
