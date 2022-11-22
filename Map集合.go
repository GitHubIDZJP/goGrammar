package main

import "fmt"

/*
*

	map是一种"无序"的键值对(key-value)

Map 是无序的，我们无法决定它的返回顺序，这是因为 Map 是使用 hash 表来实现的。
map定义:
//声明map变量，默认map是nil

	var obj map[keyDataType] valueDataType
	Example(举例):
	var mapObj map[string]string
	mapObj: map变量名
	[string]:keyDataType
	string:valueDataType
	使用make函数
	mapVarible := make(map[keyDataType] valueDataType)

	注意: 如果不初始化 map，那么就会创建一个 nil map。nil map 不能用来存放键值对
*/
func mapMain() {

	//创建集合
	var mapObj map[string]string //因为没有初始化map。所有是nil map

	mapObj = make(map[string]string)

	//map插入 key -value
	mapObj["China"] = "中国"
	mapObj["Bhutan"] = "不丹"
	mapObj["India"] = "印度"
	mapObj["Japan"] = "日本"
	mapObj["Laos"] = "老挝"
	mapObj["Nepal"] = "尼泊尔"

	//  for 循环key 打印value：调用语言范围功能：range
	for key := range mapObj {

		fmt.Println(key, "value值:\n", mapObj[key])
	}
	//判断原理在集中里是否存在
	// 包含contain
	contain, ok := mapObj["Japan"] //如果确定是真的存在，则存在，反之
	if ok {
		fmt.Println("存在", contain)
	} else {
		fmt.Println("不存在")
	}

	//delete函数
	//删除函数集合的元素，参数为map和其对用的key
	var setMap = map[string]string{"France": "巴黎", "Italy": "罗马", "Japan": "东京", "India": "新德里"}
	/* 打印地图 */
	for country := range setMap {
		fmt.Println(country, "首都是", setMap[country])
	}
	//删除key
	delete(setMap, "France")
	fmt.Println("删除元素后地图")
	/*打印地图*/
	for country := range setMap {
		fmt.Println(country, "首都是", setMap[country])
	}
}
