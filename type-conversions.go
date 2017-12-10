package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4

	// go的类型转换：表达式 T(v) 将值 v 转换为类型 T 。
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)

	// go不支持隐式转换，必须显示表示
	// var t float32 = z // 会报错

	fmt.Println(x, y, z, f)
}
