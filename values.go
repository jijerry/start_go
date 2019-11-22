package main

import "fmt"

//类型

func main() {
	//字符串可以通过'+'连接
	fmt.Println("go" + "lang")

	//整形和浮点数
	fmt.Println("1+1 = ", 1+1)
	fmt.Println("7.0/3.0 = ", 7.0/3.0)

	//布尔型,逻辑运算符
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

}
