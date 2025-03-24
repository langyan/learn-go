package main

import "fmt"

var globalVar int // 包级别变量，只能用 var 声明

func main() {
	globalVar = 100 // 修改全局变量
	fmt.Println(globalVar)

	var localVar int
	localVar = 200 // 修改局部变量
	fmt.Println(localVar)

	x := 5
	//x := 10 // 错误: no new variables on left side of :=

	fmt.Println(x)

}
