// 每个 Go 程序都是由包组成的。
// 程序运行的入口是包 main 。

package main

import (
	"fmt"
	"math/rand"
)

// 程序入口
func main() {
	// 这里的随机数每次是固定的，并不是随机
	fmt.Println("My favorite number is", rand.Intn(10))
}
