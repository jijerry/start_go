package main

import (
	"fmt"
	"math"
)

// 常量, Go支持数字，字符串，字符，__常量__
// const 声明一个常量
const s string = "constant"

func main() {
	fmt.Println(s)
	// const语句可以出现在任何var语句出现的地方
	const n = 5000

	// 常数表达式可以进行任何精度的计算
	const d = 3e20 / n
	fmt.Println(d)

	// 数值型的常量是没有确定类型的，直到被给定一个类型
	fmt.Println(int64(d))

	// 上下文需要时候，一个数可以被指定一个类型
	fmt.Println(math.Sin(n))

}
