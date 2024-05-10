package main

import (
	"fmt"
	"unsafe"
)

/*
Go 1.17引入
func Slice(ptr *ArbitraryType, len IntegerType) []ArbitraryType
从一个任意（安全）指针派生出一个指定长度的切片

用途
1.转换数组类型
*/
func main() {
	a := [16]int{3: 3, 9: 9, 11: 11}
	fmt.Println(a)
	fmt.Println(len(a), cap(a)) // 16 16

	p9 := &a[9]
	s := unsafe.Slice(p9, 5)
	fmt.Println(s)              // [9 0 11]
	fmt.Println(len(s), cap(s)) // 3 5

	//零值指针
	t := unsafe.Slice((*int)(nil), 0)
	fmt.Println(t == nil) // true

}
