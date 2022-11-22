package main

import "fmt"

/*
对数组的抽象
数组的长度不可改变，在特定的场景里集合就不太适用，go提供了一种灵活且强悍的内置类型切片("动态数组"),
与数组相比切片的长度是不固定的，可以追加元素，在追加是可能使切片的容量"增大"
*/
/***

定义数组:
vary 数组名 [数组元素长度] 数组类型
参数注解:
数组名:随便取
数组元素长度: 整形
数组类型: 数组元素的类型(int string 等)
Example：
var arr [20] float32

定义切片:
var  indentfier []Type
切片不需要说明长度
或者用make()函数来创建切片

var objSelice []Type  = make([]type,len)
也可以简写为
objSlice := make([]type,len)
也可以指定容量,其中capacity为可选参数
make([]T,length,capacity)
//这里的length指的是数组的长度并且也是切片的初始长度

切片初始化
s :=[] int{1,2,3}
直接初始化切片，[]代表切片类型，{1,2,3}初始值一次是1,2，3 其中cappacity= len = 3
s :=arr[:]

初始化切片s,是数组Arr的引用
s :=arr[startIndex:endIndex]

将 arr 中从下标 startIndex 到 endIndex-1 下的元素创建为一个新的切片
s := arr[startIndex:]

默认 endIndex 时将表示一直到arr的最后一个元素
s := arr[:endIndex]

默认 startIndex 时将表示从 arr 的第一个元素开始。
s1 := s[startIndex:endIndex]

通过切片 s 初始化切片 s1，通过内置函数 make() 初始化切片s，[]int 标识为其元素类型为 int 的切片。
s :=make([]int,len,cap)


len() 和 cap() 函数
切片是可索引的，并且可以由 len() 方法获取长度。
切片提供了计算容量的方法 cap() 可以测量切片最长可以达到多少
*/
func sliceMain() {
	// arr := []int{10, 100, 200}
	// s := arr[:endIndex]
	// fmt.Printf(s)
	var numbers = make([]int, 3, 5)

	printSlice(numbers) //len=3 cap=5 slice=[0 0 0]

	//---> 切片截取--:可以通过设置下限及上限来设置截取切片 [lower-bound:upper-bound]，
	/* 创建切片 */
	numbersFirst := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(numbersFirst) //调用下面函数

	/* 打印原始切片 */
	fmt.Println("numbersFirst ==", numbersFirst) //len=9 cap=9 slice=[0 1 2 3 4 5 6 7 8]

	/* 打印子切片从索引1(包含) 到索引4(不包含)*/
	fmt.Println("numbersFirst[1:4] ==", numbersFirst[1:4]) //[1 2 3]

	/* 默认下限为 0*/
	fmt.Println("numbersFirst[:3] ==", numbersFirst[:3]) //[0 1 2]

	/* 默认上限为 len(s)*/
	fmt.Println("numbersFirst[4:] ==", numbersFirst[4:]) //[4 5 6 7 8]

	numbersFirst1 := make([]int, 0, 5)
	printSlice(numbersFirst1)

	/* 打印子切片从索引  0(包含) 到索引 2(不包含) */
	number2 := numbersFirst[:2]
	printSlice(number2) //[0 1]

	/* 打印子切片从索引 2(包含) 到索引 5(不包含) */
	number3 := numbersFirst[2:5]
	printSlice(number3) //[2 3 4]

	//append() 和 copy() 函数
	// 如果想增加切片的容量，我们必须创建一个新的更大的切片并把原分片的内容都拷贝过

	var numbersSecond []int
	printSlice(numbersSecond) //len=0 cap=0 slice=[]

	/* 允许追加空切片 */
	numbersSecond = append(numbersSecond, 0)
	printSlice(numbersSecond) //len=1 cap=1 slice=[0]

	/* 向切片添加一个元素 */
	numbersSecond = append(numbersSecond, 1)
	printSlice(numbersSecond) //len=2 cap=2 slice=[0 1]

	/* 同时添加多个元素 */
	numbersSecond = append(numbersSecond, 2, 3, 4)
	printSlice(numbersSecond) //len=5 cap=6 slice=[0 1 2 3 4]

	/* 创建切片 numbersSecond1 是之前切片的两倍容量*/
	numbersSecond1 := make([]int, len(numbersSecond), (cap(numbersSecond))*2)

	/* 拷贝 numbersSecond 的内容到 numbersSecond1 */
	copy(numbersSecond1, numbersSecond)
	printSlice(numbersSecond1) //len=5 cap=12 slice=[0 1 2 3 4]

}
func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
