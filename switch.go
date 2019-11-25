// switch 方便的条件分支语句
package main

import (
	"fmt"
	"time"
)

func main() {

	// 基本的switch语句
	i := 2
	fmt.Println("write", i, "as")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")

	case 3:
		fmt.Println("three")

	}

	// 一个case语句中，使用逗号分隔多个表达式，可以使用可选的default分支

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it is the weekwend")

	default:
		fmt.Println("it is weedday")
	}

	// 不带switch是实现if else逻辑的另一种
	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("it is before noon")
	default:
		fmt.Println("it is after noon")
	}

}
