package main

import (
	"fmt"
	"reflect"
)

/*
-DeepEqual类型相同时=true，不相同是总=false
-然后DeepEqual将比较x和y所引用的两个值

	-若两个同类型切片、共享相同的元素序列(长度/元素的地址相同),则DeepEqual=true

DeepEqual和==的区别
1.DeepEqual(x, y)无论如何不产生恐慌，x == y在比较不可比较类型时将产生一个恐慌
2.x,y类型不相同，DeepEqual(x, y)总为false，但x == y可能为true
3.x,y类型相同，值不同，DeepEqual(x, y)可能为true，x == y总为false
4.x,y处于循环引用链中时，DeepEqual调用的结果可能为true(防止死循环)
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

	fmt.Println(reflect.DeepEqual(&q, &r)) // false
	fmt.Println(q == r)                    // false

	fmt.Println("-----方法------")
	var f1 func() = nil
	var f2 func() = func() {}

	fmt.Println(reflect.DeepEqual(f1, f1)) // true
	fmt.Println(reflect.DeepEqual(f2, f2)) // 匿名函数总是false
	fmt.Println(reflect.DeepEqual(f1, f2)) // false

	fmt.Println("-----切片比较------")
	var a, b interface{} = []int{1, 2}, []int{1, 2}
	fmt.Println(reflect.DeepEqual(a, b)) // true
	fmt.Println(a == b)                  // 恐慌
}
