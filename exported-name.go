package main

import (
	"fmt"
	"math"
)

func main() {
	// 在go的包中，首字母大写的名称是被导出的，能被包外的代码访问
	// 小字母小写不是能被访问的
	fmt.Println("包的访问控制")
	fmt.Println(math.Pi)
}
