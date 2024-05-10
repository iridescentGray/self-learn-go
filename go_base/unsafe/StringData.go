package main

import (
	"fmt"
	"unsafe"
)

/*
Go 1.20 引入
func StringData(str string) *byte
获取一个字符串底层字节序列中的第一个byte的指针。
*/
func main() {
	var s = "123"

	fmt.Println(s) // 123

	fmt.Println(unsafe.StringData(s)) // 0xxxxxx

}
