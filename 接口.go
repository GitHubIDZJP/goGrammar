package main

import (
	"fmt"
)

//Go 语言提供了另外一种数据类型即接口，它把所有的具有共性的方法定义在一起，任何其他类型只要实现了这些方法就是实现了这个接口

/**

   定义接口
type interface_name interface {
   method_name1 [return_type]
   method_name2 [return_type]
   method_name3 [return_type]
   ...
   method_namen [return_type]
}

/  定义结构体   /
type struct_name struct {
      variables   /
}

/   实现接口方法   /
func (struct_name_variable struct_name) method_name1() [return_type] {
   //   方法实现   /
}

func (struct_name_variable struct_name) method_namen() [return_type] {
   //  方法实现  /
}

*/

//接口port

func portMain() {

	//调用接口
	var phone Phone

	phone = new(NokiaPhone)
	phone.call()
	//phone.title = "诺基亚"

	phone = new(IPhone)
	phone.call()
}

type Phone interface {
	call()
}

// 诺基亚手机
type NokiaPhone struct {
	title  string
	author string
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("我是诺基亚, 我可以打电话给你!")
	nokiaPhone.title = "诺基亚"
	nokiaPhone.author = "测试机"
	//var a = nokiaPhone{nokiaPhone.author:"sd"}
	fmt.Println("测试接口:\n", nokiaPhone.title)
	fmt.Println("测试接口:\n", nokiaPhone.author)
}

type IPhone struct {
}

func (iPhone IPhone) call() {
	fmt.Println("我是苹果手机,我能给你打电话")
}
