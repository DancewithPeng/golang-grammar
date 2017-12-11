package main

import (
	"fmt"
	"math/cmplx"
)

/*
Go 的基本类型(Basic types)有:

bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // uint8 的别名

rune // int32 的别名
     // 代表一个Unicode码

float32 float64

complex64 complex128

int，uint 和 uintptr 类型在32位的系统上一般是32位，而在64位系统上是64位。
当你需要使用一个整数类型时，你应该首选 int，仅当有特别的理由才使用定长整数类型或者无符号整数类型。
*/

// 类似于“导入包语法”，变量声明也可以用“变量包”语法声明：
var (
	Tobe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {

	// const 声明常量
	// 输出格式中，%T输出类型，%v输出值
	const f = "%T(%v)\n"
	fmt.Printf(f, Tobe, Tobe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)
}
