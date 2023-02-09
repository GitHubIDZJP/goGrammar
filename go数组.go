package main

import "fmt"

/*
数组定义:
vary 数组名 [数组元素长度] 数组类型
参数注解:
数组名:随便取
数组元素长度: 整形
数组类型: 数组元素的类型(int string 等)
Example：
var arr [20] float32
Example1：
var  arr = [2]float32{12.0,15.0}
假如长度不确定，可以用...代替数组的len,编译器会根据元素个数自动推行数组长度
var arr = [...]float32{1.0,3.0,5.0}
或者
arr := [...]float32{1.0,3.0,5.0}
如果设置了数组的长度，我们还可以通过指定下标来初始化元素：
arr :=[5]float32{1:2.0,3:7.0}
注解:
指定下标1的元素为2.0
指定下标3的元素为7.0
因为未全部指定，剩下的3个则。。。。
*/
func callBack_Arr() {
	var arr_a [10]int //定义一个数组长度为10的数组变量arr_a
	var k, l int
	for k = 0; k < 10; k++ {
		arr_a[k] = k + 100
		fmt.Printf("数组下标[%d] = %d\n", k, arr_a[k])
	}

	//
	for l = 0; l < 10; l++ {
		fmt.Printf("数组下标[%d] = %d\n", l, arr_a[l])
	}

	balance3 := [5]float32{1: 2.0, 3: 7.0}
	for k = 0; k < 5; k++ {
		fmt.Printf("balance3[%d] = %f\n", k, balance3[k])
	}

}

/**
二维数组
Example:
结构定义:
var m_Arr [数组长度][数组长度] 数组类型
var m_Arr [2][4]
三维数组
Example:
结构定义:
var m_Arr [数组长度][数组长度][数组长度] 类型
var m_Arr [2][4][1]
*/

func multidimensionalArray() {
	//多维数组

	m_Arr := [][]int{} //创建2维空数组
	row1 := []int{1, 2, 3}
	row2 := []int{5, 6, 7}
	m_Arr = append(m_Arr, row1, row2)
	fmt.Print("打印二维数组[%d] = %d\n", m_Arr)

	//初始化多维数组:
	m_ArrDef := [3][4]int{
		{0, 1, 2, 3},   /*  第一行索引为 0 */
		{4, 5, 6, 7},   /*  第二行索引为 1 */
		{8, 9, 10, 11}, /* 第三行索引为 2 */
	}

	fmt.Print("打印 = %d\n", m_ArrDef)
	// 以上代码中倒数第二行的 } 必须要有逗号，因为最后一行的 } 不能单独一行，也可以写成这样

	a_m_a := [3][4]int{
		{0, 1, 2, 3},   /*  第一行索引为 0 */
		{4, 5, 6, 7},   /*  第二行索引为 1 */
		{8, 9, 10, 11}} /* 第三行索引为 2 */

	fmt.Println("多维数组打印 = %d\n", a_m_a)

	//创建一个2行2列的二维数组
	m_Arr_Str := [2][2]string{}
	//给二维数组添加元素
	m_Arr_Str[0][0] = "春"
	m_Arr_Str[0][1] = "意"
	m_Arr_Str[1][0] = "昂"
	m_Arr_Str[1][1] = "然"

	//访问数组：通过坐标来访问
	accArrCoord := m_Arr_Str[1][0]
	fmt.Println("访问二维数组坐标:\n", accArrCoord) //昂
	//或者
	var accNewArrEl string = m_Arr_Str[0][1]
	fmt.Println("访问二维数组坐标:\n", accNewArrEl) //意

	//二维数组用for循环打印输出访问
	var forArr = [5][2]int{{0, 0}, {1, 2}, {2, 4}, {3, 6}, {4, 8}}
	var i, j int
	for i = 0; i < 5; i++ {
		for j = 0; j < 2; j++ {
			fmt.Printf("二维数组:forArr[%d][%d] = %d\n", i, j, forArr[i][j])
		}

	}

	// 创建空的二维数组：创建各个维度元素数量不一致的多维数组
	animals := [][]string{}

	// 创建三一维数组，各数组长度不同
	row11 := []string{"fish", "shark", "eel"}
	row22 := []string{"bird"}
	row33 := []string{"lizard", "salamander"}

	// 使用 append() 函数将一维数组添加到二维数组中
	animals = append(animals, row11)
	animals = append(animals, row22)
	animals = append(animals, row33)

	// 循环输出
	for i := range animals {
		fmt.Printf("Row: %v\n", i)
		fmt.Println(animals[i])
	}

}
