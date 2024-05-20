package main

import (
	"fmt"
	"unsafe"
)

/*
Go 1.17引入
func Sizeof(variable ArbitraryType) uintptr
用来取得一个值的尺寸（亦即此值的类型的尺寸
同一个类型的不同值,此函数的返回值总是相同的
*/
func main() {
	var x struct {
		a int64
		b bool
		c string
	}

	fmt.Println(unsafe.Sizeof(x.a)) // 8
	fmt.Println(unsafe.Sizeof(x.b)) // 1
	fmt.Println(unsafe.Sizeof(x.c)) // 16
	fmt.Println(unsafe.Sizeof(x))   // 32

}
