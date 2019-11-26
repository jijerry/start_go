// slice 是go中一个关键的数据类型，比数组更强大的序列接口
package main

import "fmt"

func main() {

	// slice的类型仅有所包含的元素决定（不像数组中需要元素的个数）
	// 创建一个长度非零的空slice,使用内建的make
	// 这里创建一个长度3的string类型的slice,初始化为0值

	s := make([]string, 3)
	fmt.Println("emp", s)

	// 可以和数组一起设置和得到值
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set", s)
	fmt.Println("get", s[2])

	// len返回slice长度
	fmt.Println("len", len(s))

	// slice支持比数组更多的操作
	// 内建append， 返回一个包含了一个或者多个新值的slice,注意我们接受返回由append返回的新slice值
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println(s)

	// 使用copy，创建一个空的和s相同长度的c,并且将s赋值给c
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println(c)

	// slice支持通过slice[low:high]语法进行切片
	l := s[2:5]
	fmt.Println("sli", l)

	// 这个slice从s[0]到s[5]
	l = s[:5]
	fmt.Println("sl2", l)

	// 这个slice从s[2]到slice的最后一个值
	l = s[2:]
	fmt.Println("sl3", l)

	// 在一行代码汇总声明并初始化一个slice变量
	t := []string{"g", "h", "i"}
	fmt.Println("dcl", t)

	// slice可以组成多维数据结构，内部的slice长度可以不同
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 1; j < innerLen; j++ {

			twoD[i][j] = i + j

		}
	}
	fmt.Println("2d", twoD)
}
