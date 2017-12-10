package main

import (
	"fmt"
)

func main() {
	// go支持类型推导，会自动根据右值推导出左值的类型
	v := 0.867 + 0.5i // 3.14， 0.867 + 0.5i
	fmt.Printf("v is of type: %T\n", v)
}
