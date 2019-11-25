// if else分支结构在go中是直截了当

package main

import "fmt"

func main() {
	// 基本例子

	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is old")
	}

	// 只用if 不用else
	if 8%4 == 0 {
		fmt.Println("8 is divisable by 4")
	}

	// 条件语句之前可以有一个语句，任何在这里声明的变量可以在所有条件分支中使用
	if num := 10; num < 0 {
		fmt.Println(num, "is negative")

	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}

// go中可以不是用圆括号，但是需要使用花括号
// go中没有三目运算符，所有需要使用完整if语句
