package main

import "fmt"

// import "fmt"

func study() {

	/**
	var a *int地方
	var a []int
	var a map[string] int
	var a chan int
	var a func(string) int
	var a error // error 是接口
	*/
	var i int     //0
	var f float64 //0
	var b bool    //false
	var s string  // ""
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	/*定义常量
	const b type  = value
	b:常量名
	type:常量名
	value:常量值
	*/
	//显式类型定义： const b string = "abc"
	//隐式类型定义： const b = "abc"

	//多个相同类型的声明可以简写为：const c_name1, c_name2 = value1, value2

	/**
	  & 求变量地址

	*/

	var a_jk int = 10
	if a_jk < 20 {
		fmt.Print("a < 20\n")
	}
	/*

	   0值等于:
	   1值等于:
	   2值等于:
	   3值等于:
	   4值等于:
	   5值等于:
	   6值等于:
	   7值等于:
	   8值等于:
	   9值等于:
	   10值等于:
	   11值等于:
	   12值等于:
	   13值等于:
	   14值等于:
	   15值等于:
	   16值等于:
	   17值等于:
	   18值等于:
	   19
	*/
	for i := 0; i < 20; i++ {
		fmt.Print("值等于:\n", i)

	}
	fmt.Print("\n\n\n")
	sum := 0
	for hj := 0; hj <= 10; hj++ {
		sum += hj
		fmt.Print("和:\n", sum) //55
	}

	/**	sum1 := 0
	for {
		sum1++ // 无限循环下去
	}
	fmt.Println(sum1) // 无法输出
	*/
	//For-each range 循环
	//这种格式的循环可以对字符串、数组、切片等进行迭代输出元素。

	strings := []string{"google", "runoob"} //字符串数组
	for i, s := range strings {
		fmt.Println(i, s)
		/**
				0 google
		1 runoob
		*/
	}

	numbers := [6]int{1, 2, 3, 5}
	for i, x := range numbers {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i, x)
	}
	/**

	第 0 位 x 的值 = 1
	第 1 位 x 的值 = 2
	第 2 位 x 的值 = 3
	第 3 位 x 的值 = 5
	第 4 位 x 的值 = 0
	第 5 位 x 的值 = 0
	*/

	//for 循环的 range 格式可以省略 key 和 value，如下实例
	map1 := make(map[int]float32)
	map1[1] = 1.0
	map1[2] = 2.0
	map1[3] = 3.0
	map1[4] = 4.0

	// 读取 key 和 value
	for key, value := range map1 {
		fmt.Printf("key is: %d - value is: %f\n", key, value)
	}
	/**
	  key is: 4 - value is: 4.000000
	  key is: 1 - value is: 1.000000
	  key is: 2 - value is: 2.000000
	  key is: 3 - value is: 3.000000

	*/
	// 读取 key
	for key := range map1 {
		fmt.Printf("key is: %d\n", key)
	}
	/**

	  key is: 1
	  key is: 2
	  key is: 3
	  key is: 4
	*/
	// 读取 value
	for _, value := range map1 {
		fmt.Printf("value is: %f\n", value)
	}
	/**
	  value is: 1.000000
	  value is: 2.000000
	  value is: 3.000000
	  value is: 4.000000
	*/

}
