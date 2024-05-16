package main

import (
	"fmt"
	"reflect"
)

/*
DeepEqual类型相同的情况下才返回true。 比较元素中含有函数值的容器值或者比较字段中含有函数值的结构体值也是类似的。
若两个同类型切片共享相同的元素序列（长度/元素的地址相同）DeepEqual比较true
*/
func main() {
	type Book struct{ page int } //具名结构体
	x := struct{ page int }{123} //匿名结构体
	y := Book{123}

	fmt.Println("-----important: 具名结构体/匿名结构体------")
	fmt.Println(reflect.DeepEqual(x, y)) // false
	fmt.Println(x == y)                  // true

	fmt.Println("-----具名结构体/具名结构体------")
	z := Book{123}
	fmt.Println(reflect.DeepEqual(z, y)) // true
	fmt.Println(z == y)                  // true

	fmt.Println("-----具名结构体指针/具名结构体指针------")
	fmt.Println(reflect.DeepEqual(&z, &y)) // true
	fmt.Println(&z == &y)                  // false

	fmt.Println("-----具名结构体循环引用------")
	type Node struct{ peer *Node }
	var q, r Node
	q.peer = &q // 形成一个循环引用链

	fmt.Println(reflect.DeepEqual(&q, &r))
	fmt.Println(q == r) // false

	fmt.Println("-----方法------")
	var f1 func() = nil
	var f2 func() = func() {}

	fmt.Println(reflect.DeepEqual(f1, f1)) // true
	fmt.Println(reflect.DeepEqual(f2, f2)) // false 匿名函数比较总是false
	fmt.Println(reflect.DeepEqual(f1, f2)) // false

	fmt.Println("-----切片比较------")
	var a, b interface{} = []int{1, 2}, []int{1, 2}
	fmt.Println(reflect.DeepEqual(a, b)) // true
	fmt.Println(a == b)                  // 产生恐慌
}
