// 'for' 是go中唯一的循环结构，for循环的三种使用方式

package main

import "fmt"

func main() {

	// 最常用的方式，带单个循环条件

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// 经典的初始化、条件，后续形式for循环
	for j := 7; j <= 9; j++ {

		fmt.Println(j)
	}

	// 不带条件的for循环将一直执行，直到循环内部使用了break或者break跳出循环
	for {
		fmt.Println("loop")
		break
	}
}
