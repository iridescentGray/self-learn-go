package main

import (
	"fmt"
	"unsafe"
)

/*
Go 1.20 引入
func SliceData(slice []ArbitraryType) *ArbitraryType
获取一个切片底层元素序列中的第一个元素的指针
*/
func main() {
	var s = []int{1, 2, 3, 6}

	fmt.Println(s) // [1 2 3 6]

	fmt.Println(unsafe.SliceData(s)) // 0xcxxxxx

}
