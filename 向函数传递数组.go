package main

import "fmt"

//向函数传递数组参数，你需要在函数定义时，声明形参为数组，我们可以通过以下两种方式来声明

/*
形参设定数组大小：10
void func (param [10]int ){

}
形参未设定数组大小
void func (param [] int){

}
*/

// 函数接收整型数组参数，另一个参数指定了数组元素的个数，并返回平均值
func recIntArrParam() {
	var socre = [5]int{100, 30, 80, 70, 60}
	var avg float32
	//数组作为参数传给函数
	avg = getAverage(socre, 5)
	//输出平均值
	fmt.Println("平均值:", avg)

}

// 获取平均值
func getAverage(arr [5]int, size int) float32 {
	var i, sum int
	var avg float32
	for i = 0; i < size; i++ {
		sum += arr[i]
	}
	avg = float32(sum) / float32(size)
	return avg
}
