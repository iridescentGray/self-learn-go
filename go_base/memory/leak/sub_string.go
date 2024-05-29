package main

import (
	"fmt"
	"strings"
)

var s0 string // 一个包级变量
func f(s1 string) {
	s0 = s1[:50]
	// 目前，s0和s1共享同一个内存块, s1到不再被使用了，但是s0仍然在使用中,所以它们共享的内存块将不会被回收

	//解决方案 不共享内存
	//s0 = string([]byte(s1[:50]))
	// strings.Clone(s1[:50])
}

/*
子字符串内存泄露
*/
func main() {
	s := strings.Repeat("a", 1*1024*1024)
	fmt.Printf("length: %d\n", len(s))
	f(s)
}
