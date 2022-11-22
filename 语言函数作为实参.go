package main

import (
	"fmt"
	"math"
)

// 函数定义后可作为另外一个函数的实参数传入
func main_hs() {
	/* 声明函数变量 */
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x) //3
	}

	/* 使用函数 */
	fmt.Println(getSquareRoot(9))

}
