package main

import (
	"fmt"
	"strings"
	"unsafe"
)

func Float64frombits(b uint64) float64 {
	return *(*float64)(unsafe.Pointer(&b))
}

// 类似于strings标准库包中的Builder
func ByteSlice2String(bs []byte) string {
	return *(*string)(unsafe.Pointer(&bs))
}

/*
用途
1.转换原本不可能转换的类型
2.将一个不再使用的字节切片转换为一个字符串（从而避免对底层字节序列的一次开辟和复制）
*/
func main() {
	//把uint64  转换为 float64,两者尺寸应该相仿
	var u1 uint64 = 1000
	fmt.Println(u1)                  //1000
	fmt.Println(Float64frombits(u1)) //4.94e-321
	fmt.Println(float64(u1))         //1000

	//把[]MyString 转换为 []string
	type MyString string
	ms := []MyString{"C", "C++", "Go"}
	// ss := ([]string)(ms) // 编译错误,只能通过unsafe转换
	var ss []string = *(*[]string)(unsafe.Pointer(&ms))
	ss[1] = "Zig"
	fmt.Printf("%s\n", ms) // [C Zig Go]

}
