package main

import "fmt"

/*
for 循环的 range 格式可以对 slice、map、数组、字符串等进行迭代循环。格式如下：

	 for key,value :=range oldMap{
		newMap[key] = value
	 }
	 以上代码中的 key 和 value 是可以省略。
	 如果只想读取 key，格式如下：
	 for key := range oldMap
	 或者这样：
	 for key, _ := range oldMap
	 如果只想读取 value，格式如下：

for _, value := range oldMap
*/
func RangeMain() {

	// var forArr = [5][2]int{{0, 0}, {1, 2}, {2, 4}, {3, 6}, {4, 8}}
	// var i, j int
	// for i = 0; i < 5; i++ {
	// 	for j = 0; j < 2; j++ {
	// 		fmt.Printf("二维数组:forArr[%d][%d] = %d\n", i, j, forArr[i][j])
	// 	}

	// }

	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Println("pow[i],pow[v] %d\n", i, v)
		fmt.Printf("2**%d = %d\n", i, v)
		/*
					2**0 = 1
			2**1 = 2
			2**2 = 4
			2**3 = 8
			2**4 = 16
			2**5 = 32
			2**6 = 64
			2**7 = 128

		*/
	}

	//2

	map1 := make(map[int]float32)
	map1[1] = 1.0
	map1[2] = 2.0
	map1[3] = 3.0
	map1[4] = 4.0

	// 读取 key 和 value
	for key, value := range map1 {
		fmt.Printf("key is: %d - value is: %f\n", key, value)
	}

	// 读取 key
	for key := range map1 {
		fmt.Printf("key is: %d\n", key)
	}

	// 读取 value
	for _, value := range map1 {
		fmt.Printf("value is: %f\n", value)
	}
	/*
			key is: 4 - value is: 4.000000
		key is: 1 - value is: 1.000000
		key is: 2 - value is: 2.000000
		key is: 3 - value is: 3.000000
		key is: 1
		key is: 2
		key is: 3
		key is: 4
		value is: 1.000000
		value is: 2.000000
		value is: 3.000000
		value is: 4.000000

	*/

	//3
	//这是我们使用 range 去求一个 slice 的和。使用数组跟这个很类似
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum) //sum: 9
	//在数组上使用 range 将传入索引和值两个变量。上面那个例子我们不需要使用该元素的序号，所以我们使用空白符"_"省略了。有时侯我们确实需要知道它的索引。
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i) //1
		}
	}
	//range 也可以用在 map 的键值对上。
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
		/*
					a -> apple
			b -> banana
		*/
	}

	//range也可以用来枚举 Unicode 字符串。第一个参数是字符的索引，第二个是字符（Unicode的值）本身。
	for i, c := range "go" {
		fmt.Println(i, c)
		/**
				0 103
		1 111
		*/
	}

}
