package main

import (
	"fmt"
	"unsafe"
)

/*
Go 1.17引入
func Offsetof(selector ArbitraryType) uintptr
取得一个结构体值的某个字段的地址相对于此结构体值的地址的偏移
同一个结构体类型的不同值的对应相同字段，此函数的返回值总是相同的
*/
func main() {
	var x struct {
		a int64
		b bool
		c string
	}

	fmt.Println(unsafe.Offsetof(x.a)) // 0
	fmt.Println(unsafe.Offsetof(x.b)) // 8
	fmt.Println(unsafe.Offsetof(x.c)) // 16

}
