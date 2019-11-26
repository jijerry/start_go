// go中，数组是一个固定长度的数列
package main

import "fmt"

func main() {

	// 创建了一个数组a,存放5个int
	// 元素的类型和长度都是数组的一部分，数组默认是零值

	var a [5]int
	fmt.Println(a)

	// 可以使用array[index] = value语法设置数组
	// 指定位置的值，可以用array[index]得到
	a[4] = 100
	fmt.Println("set", a)
	fmt.Println("get", a[4])

	// 内置函数len获取长度
	fmt.Println("len", len(a))

	// 使用语法在一行初始化一个数组
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl", b)

	// 数组的存储类型是单一的，但是可以组合数组来构建多维数据结构
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {

			twoD[i][i] = i + j
		}
	}
	fmt.Println("2d", twoD)
}

// 使用fmt.println打印数组时候，会使用[v1 v2 v3]格式显示
