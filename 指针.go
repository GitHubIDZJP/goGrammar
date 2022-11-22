package main

import "fmt"

//指针
/*
变量是一种使用方便的占位符，用于引用计算机内存地址

GO语言的取地址符是&,放到一个变量前使用就会返回变量的”内存地址“

var a int = 10
fmt.Printf("获取变量的地址:%x\n",&a) ////变量的地址: 20818a220

指针:
一个指针变量指向了一个只的内存地址
var var_name *var-Type
注释:
*:指定变量是作为一个指针
var_name : 指针变量名
var-Type ：为指针类型
Example：
var ip *int        指向整型
var fp *float32    指向浮点型

*/
func pointerMain() {
	var a int = 20
	var ip *int
	ip = &a /* 指针变量的存储地址 */

	fmt.Printf("a的值:%x\n", &a)
	/* 指针变量的存储地址 */
	fmt.Printf("ip变量的指针地址:%x\n", ip)
	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip)
	// 空指针:当一个指针被定义后没有分配到任何变量时，值为nil
	// nil：也叫空指针
	// 一个指针变量通常叫ptr
	// nil在概念上和其它语言的null、None、nil、NULL一样，都指代零值或空值。
	var ptr *int
	fmt.Printf("ptr 的值为 : %x\n", ptr) //0
	/**
	空指针判断：
	if(ptr != nil)      ptr 不是空指针
	if(ptr == nil)     ptr 是空指针
	*/

	//指向指针的指针: **用双星号
	var aa int
	var ptra *int
	var pptra **int

	aa = 3000

	/* 指针 ptr 地址 */
	ptra = &aa

	/* 指向指针 ptr 地址 */
	pptra = &ptra

	/* 获取 pptr 的值 */
	fmt.Printf("变量 a = %d\n", aa)
	fmt.Printf("指针变量 *ptr = %d\n", *ptra)
	fmt.Printf("指向指针的指针变量 **pptr = %d\n", **pptra)

}
