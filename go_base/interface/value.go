package main

import "fmt"

type Aboutable interface {
	About() string
}

// 类型*Book实现了接口类型Aboutable。
type Book struct {
	name string
}

func (book *Book) About() string {
	return "Book: " + book.name
}

func main() {
	// 一个*Book值被包裹在了一个Aboutable值中。
	var a Aboutable = &Book{"Go语言101"}
	fmt.Println(a) // &{Go语言101}

	// i是一个空接口值。类型*Book实现了任何空接口类型。
	var i interface{} = &Book{"Rust 101"}
	fmt.Println(i) // &{Rust 101}

	// Aboutable实现了空接口类型interface{}。
	i = a
	fmt.Println(i) // &{Go语言101}

	var i interface{}
	i = []int{1, 2, 3}
	fmt.Println(i) // [1 2 3]
	i = map[string]int{"Go": 2012}
	fmt.Println(i) // map[Go:2012]
}
