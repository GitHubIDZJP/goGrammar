package main

import "fmt"

// https://www.runoob.com/go/go-structures.html
/**
数组:可以存储同一类型的数据
结构体:可以存储不同类型的数据
type structName struct{
	param1 dataType
	param2 dataType
	param3 dataType
}
*/

type Books struct {
	title    string
	author   string
	subjects string
	book_id  int
}

func structMain() {
	//创建一个结构体
	var s_obj = Books{"GO语言", "邹俊哦", "物理", 23}
	fmt.Println("打印结构体:\n", s_obj) //{GO语言 邹俊哦 物理 23}
	//通过key=value形式赋值
	var structObj = Books{title: "学习", author: "邹俊", subjects: "数学", book_id: 20}
	fmt.Println("打印结构体:\n", structObj) // {学习 邹俊 数学 20}
	//忽略字段为0或为空
	var lgnoreFieldIsZeroOrNULL = Books{title: "学科", book_id: 20}
	fmt.Println("忽略字段:\n", lgnoreFieldIsZeroOrNULL) //{学科   20}

	//访问结构体变量

	var fetchStructValue_author = s_obj.author
	var fetchStructValue_book_id = s_obj.book_id
	fmt.Println("获取结构体作者字段:\n", fetchStructValue_author)
	fmt.Println("获取结构体书本ID字段:\n", fetchStructValue_book_id)

	//结构体作为参数回调
	//var callBackStruct =
	strcutParam(s_obj)

	// 结构体指针
	//fmt.Printf("打印结构体指针:\n", &s_obj)
	strcutParamPointer(&s_obj)

}

//结构体作为函数参数

func strcutParam(obj Books) {
	fmt.Printf("Book title : %s\n", obj.title)
	fmt.Printf("Book author : %s\n", obj.author)
	fmt.Printf("Book subject : %s\n", obj.subjects)
	fmt.Printf("Book book_id : %d\n", obj.book_id)
}
func strcutParamPointer(structPointer *Books) {
	fmt.Printf("结构体指针 title : %s\n", structPointer.title)
	fmt.Printf("结构体指针 author : %s\n", structPointer.author)
	fmt.Printf("结构体指针 subject : %s\n", structPointer.subjects)
	fmt.Printf("结构体指针 book_id : %d\n", structPointer.book_id)
}
