package main

import "fmt"

func main() {
	fmt.Println("测试go打印")
	channelMain()           //通道
	concurrenceMain()       //并发
	errHandle()             //错误处理
	portMain()              //接口
	typeConverionMain()     //类型转换
	recursionMain()         //递归调用
	mapMain()               //map
	RangeMain()             //范围
	sliceMain()             //切片
	structMain()            //结构体
	pointerMain()           //指针
	recIntArrParam()        //向函数传递数组参数
	multidimensionalArray() //多为数组
	callBack_Arr()          //数组
	CallParameter()
	main_hs()
	swapMain()        //传递的是参数地址
	lanuageFunction() //语言函数

	fmt.Print("\n\n\n")

	study()
	fmt.Print("\n\n\n")
	customType()
	//不推荐
	print("测试输出\n") //不带换行

	//推荐
	fmt.Println("dat02")
	fmt.Println("dat02", "收到火炬大厦", "sdsd") //字符串多个，会自动空格

	//fmt扩展:格式化输出
	//"老汉开着%s,去接ALex的媳妇"  %s占位

	// fmt包 扩展：格式化输出
	//%s，占位符“文本®
	//%d，占位符 整数
	//%f，占位符 小数（淨点数）
	//%.2f，占位符 小数（保留小数点后2位，四含五入）
	//百分比

	fmt.Printf("老汉开着%s，去接%s的媳妇，多少钱一次？%d块。娘子给打折吧，%.2f怎么样？小哥哥包你100%%满意”，“车", "老男孩", 100, 3.889)

	//注释
	//单行 //
	//多行 /**/
	//选中要注释的代码，直接com+?就行

	//整形
	fmt.Println(666)
	fmt.Println(6 + 9)
	fmt.Println(9 - 6)
	fmt.Println(9 * 6)
	fmt.Println(9 / 3)  //商
	fmt.Println(10 % 3) //余数

	//字符串
	fmt.Println("时代光华感受到")
	fmt.Println("我是" + "你")
	fmt.Println(1 + 2)     // 3
	fmt.Println("1" + "2") // "12"

	//布尔类型:真假
	fmt.Println(1 > 2) //false
	fmt.Println(1 > 0) //true

	if 1 > 2 {
		fmt.Println("假")
	} else {
		fmt.Println("真")
	}

	fmt.Println("qw" == "sd") //false

	//变量
	//理解为昵称

	var sd string = "水鬼 " //起一个常量为字符串的SD变量
	fmt.Println(sd)       //申明了必须使用，或者打印不然会报错

	var flag bool = true
	fmt.Println(flag)

	//先声明，后赋值
	var fetch string
	fetch = "获取"
	fmt.Println(fetch)

	//存储结果，方便之后使用
	var v1 int = 1
	var v2 int = 1
	var v3 int = v1 + v2
	fmt.Println(v3)

	var name string
	fmt.Scanf("%s", &name)
	if name == "aq" {
		fmt.Println("输入正常")
	} else {
		fmt.Println("输入错误")
	}
	/*变量名要求
	必须包含字母，数字，下划线
	var a string
	var 1a string
	var _ string
	都可以
	*/
	//不能使用go内置的关键字
	// var  var String ="sdsd"
	//可以去官网看官方的关键字 25个

	//ide申明一个变量不合法时会报错

	/*
		合法
		var n1 int
		var dart bool
		var _9 string
	*/

	//建议
	//2规则：
	//见名知意
	//驼峰命名：函数名或者变量名很长时适合用驼峰命名
	fmt.Println("输入正常sdsddssd")
	print(string('A' + 2))
	fmt.Print("sdsd")
}
