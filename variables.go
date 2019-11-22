package main

import "fmt"

//变量
func main() {

	// 'var' 声明1个或者多个变量
	var a string = "inital"
	fmt.Println(a)

	// 声明多个变量
	var b, c int = 1, 2
	fmt.Println(b, c)

	// 自动推断初始化的变量类型
	var d = true
	fmt.Println(d)

	// 声明了变量但是没有初始值，会被初始化为0
	var e int
	fmt.Println(e)

	// := 声明并初始化的简写
	f := "short"
	fmt.Println(f)
}
