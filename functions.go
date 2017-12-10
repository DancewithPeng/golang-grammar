package main

import (
	"fmt"
)

// go的函数，
// 带返回值的函数
func add(x int, y int) int {
	return x + y
}

// 如果参数的返回值类型相同可以这样写
func xAdd(x, y int) int {
	return x + y
}

// go函数支持多返回值
func swap(x, y string) (string, string) {
	return y, x
}

// go的函数返回值可以被命名
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x

	// 使用命名返回值之后，同样需要写return关键字，但不需要再指定返回的变量了
	return
}

func main() {
	fmt.Println(add(5, 10))
	fmt.Println(xAdd(10, 20))
	fmt.Println(swap("Hello", "World!"))

	a, b := split(20)
	fmt.Println(a, b)
}
