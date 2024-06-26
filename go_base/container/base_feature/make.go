package main

import "fmt"

func main() {
	// 创建映射。
	fmt.Println(make(map[string]int)) // map[]

	m := make(map[string]int, 3)
	fmt.Println(m, len(m)) // map[] 0

	// 创建切片。
	s := make([]int, 3, 5)
	fmt.Println(s, len(s), cap(s)) // [0 0 0] 3 5

	s = make([]int, 2)
	fmt.Println(s, len(s), cap(s)) // [0 0] 2 2
}
