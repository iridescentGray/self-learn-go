package main

import "fmt"

var (
	_ = add_a() //4
	a = b / 2
	b = 6
	c = add_a() //5 在 _ 之后
)

func add_a() int {
	a++
	return a
}

func main() {
	fmt.Println(a, b, c) // 5 6 5
}
