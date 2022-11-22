package main

import (
	"fmt"
	"strconv"
	/***
	  %T:输出值的类型
	*/)

/*
类型转换
将一种数据类型转为另外一种类型的变量
type_name(expression)
type_name 为类型
expression 为表达式
*/
func typeConverionMain() {
	var a int = 17
	var c int8 = 50
	var h string = "xsddfdf"
	b := string(a)
	d := int16(c)
	s := []byte(h) //先将 string 类型通过 []byte 类型函数转为 []byte （等同于 []uint8）
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", d)
	fmt.Printf("%T \n", s)

	//字符串类型转int

	aStr := "100"
	bInt, err := strconv.Atoi(aStr)
	if err == nil {
		fmt.Printf("字符串类型:aStr：%T %s,整型bInt：%T %d", aStr, aStr, bInt, bInt)
	} else {
		fmt.Printf("err: %s\n", err)
	}
	//整型转字符串

	cInt := 200
	dStr := strconv.Itoa(cInt)
	fmt.Printf("整型cInt：%T %d,字符串dStr：%T %s", cInt, cInt, dStr, dStr)

}
