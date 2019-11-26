// map是go内置的关联数据类型
// 在一些成为hash或者字典

package main

import (
	"fmt"
)

func main() {
	// 要创建一个空map,需要使用内建make
	// make(map[key-type]val-type)

	m := make(map[string]int)

	// 使用make[key] = val设置键值对
	m["k1"] = 7
	m["k2"] = 13

	// 使用println打印一个map输出所有的键值对
	fmt.Println("map", m)

	// 使用name[key]获取键值对
	v1 := m["k1"]
	fmt.Println("v1", v1)

	// len 返回键值对的数目
	fmt.Println("len", len(m))

	// delete从map中删除键值对
	delete(m, "k2")
	fmt.Println("map", m)

	// 从一个map取值时候，可选的第二个返回值指示这个键是在这个map中，可以用来消除键不存在或者键有零值，像‘0'或者“”产生歧义
	_, prs := m["k2"]
	fmt.Println("prs", prs)

	// 通过在同一行声明和初始化一个新的map
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map", n)

}

//fmt.println打印的时候，按照 map[k1:7 k2:13]格式输出的
