package main

import (
	"fmt"
	"unsafe"
)

/*
Go 1.20 引入
func String(ptr *byte, len IntegerType) string
从一个任意（安全）byte指针派生出一个指定长度的字符串。
*/
func main() {
	var s = "123"
	fmt.Println(s) // 123

	str_byte := unsafe.StringData(s)
	fmt.Println(unsafe.String(str_byte, len(s))) // 123

}
