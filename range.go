// 用于迭代各种数据结构
package main

import "fmt"

func main() {

	//这里使用range来统计一个slice各元素之和
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums{
		sum += num
	}

	fmt.Println("sum", sum)

	// range和slice中同样提供每个项的索引和值，上面我们不需要索引，可以使用_空值定义符忽略它，有时候实际是需要的
	for i, num := range nums{

		if num == 3 {
			fmt.Println("index", i)
		}

	}

	// range在map中迭代键值对
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs{
		fmt.Printf("%s -> %s\n", k, v)


	}

	//range在字符串中迭代unicode编码，第一个返回的是go的起始位置，第二个是go自己
	for i, c := range "go"{
		fmt.Println(i, c)
	}
}