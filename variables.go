package main

import (
	"fmt"
)

// 声明变量，变量可以是包级别的，也可以是函数级别的
var c, python, java bool

// 变量初始化
var x, y int = 10, 20

// “短变量声明”语法不支持包级别变量，也就是不能放到函数外
// z := 10086 // 会报错

func main() {

	// 函数级别的变量
	var i int

	// 变量支持类型推导，会自动根据初始值推导出变量的类型
	var objc, swift, ruby = true, false, "no"

	// 函数级别的变量支持“短变量声明”语法，省略了var和类型
	a := 100
	b, c := false, "good!"

	fmt.Println(c, java, python, i)
	fmt.Println(x, y)
	fmt.Println(objc, swift, ruby)
	fmt.Println(a, b, c)
}
