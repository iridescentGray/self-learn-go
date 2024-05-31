package main

import "fmt"

type Books struct {
	name string
}

func main() {
	// 编译器将把123的类型推断为内置类型int
	var x interface{} = 123

	//
	n, ok := x.(int)
	fmt.Println(n, ok) // 123 true
	n = x.(int)
	fmt.Println(n) // 123

	//
	a, ok := x.(float64)
	fmt.Println(a, ok) // 0 false

	var boo interface{} = Books{name: "???"}
	fmt.Println(boo) // ???
	c, ok := boo.(*Books)
	fmt.Println(c, ok) // <nil> false

	d, ok := boo.(Books)
	fmt.Println(d, ok) // {???} true

	// 将产生一个恐慌
	a = x.(float64)
}
