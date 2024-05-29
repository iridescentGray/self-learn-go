package main

import (
	"fmt"
	"unsafe"
)

/*
Go 1.17引入
func Alignof(variable ArbitraryType) uintptr
用来取得一个值在内存中的地址对齐保证（address alignment guarantee）
同一个类型的值做为结构体字段和非结构体字段时地址对齐保证可能是不同的
*/
func main() {
	var x struct {
		a int64
		b bool
		c string
	}

	fmt.Println(unsafe.Alignof(x.a)) // 8 int64
	fmt.Println(unsafe.Alignof(x.b)) // 1 bool
	fmt.Println(unsafe.Alignof(x.c)) // 8 string
	fmt.Println(unsafe.Alignof(x))   // 8  struct

}
