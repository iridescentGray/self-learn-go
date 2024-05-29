package main

import (
	"fmt"
)

var slice_0 []int // 一个包级变量
func slice_leak(s1 []int) {
	slice_0 = s1[:10]
	//解决方案 不共享内存
	//s0 = string([]byte(s1[:50]))
	// strings.Clone(s1[:50])
}

/*
子字符串内存泄露
*/
func main() {
	s := make([]int, 50)
	fmt.Printf("length: %d\n", len(s))
	slice_leak(s)
}
