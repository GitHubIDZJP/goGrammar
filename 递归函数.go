package main

import "fmt"

//递归：运行的过程中调用自己

func recursionMain() {
	recursion() //递归调用

	var i int = 15
	fmt.Printf("%d 的阶乘是 %d\n", i, Factorial(uint64(i))) //1307674368000

	var ii int
	for ii = 0; ii < 10; i++ {
		//fmt.Printf("%d\t", fibonacci(ii))
	}

}
func recursion() {
	fmt.Println("递归调用")
}

// 阶层:
func Factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}

// 斐波那契数列
func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}
