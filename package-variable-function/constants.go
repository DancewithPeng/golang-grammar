package main

import "fmt"

// go中常量用const关键字声明
const Pi = 3.14

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	/*
		常量可以是字符、字符串、布尔或数字类型的值。
		常量不能使用 := 语法定义。
	*/
	const Truth = true
	fmt.Println("Go rules?", Truth)
}
