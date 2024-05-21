package main

import "fmt"

// x是否依赖于a和b，不同的编译器有不同的见解，所以不推荐这样使用
var x = I(T{}).ab()

var q = w
var w = 42

type I interface{ ab() []int }
type T struct{}

func (T) ab() []int { return []int{q, w} }

/*
x的初始化 并不依赖q 所以x中的q=0
*/
func main() {
	fmt.Println(q) //42
	fmt.Println(w) //42
	fmt.Println(x) //[0 42]
}
