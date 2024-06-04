package main

import (
	"container/ring"
	"fmt"
)

/*
环形链表
最后一个元素指向第一个元素，从而形成一个环
在需要循环访问元素的场景中非常有用
*/
func main() {
	// 包含 5 个元素的环
	r := ring.New(5)
	for i := 0; i < r.Len(); i++ {
		r.Value = i
		r = r.Next()
	}

	fmt.Println("-----遍历元素-----")
	r.Do(func(p interface{}) {
		fmt.Println(p.(int))
	})

	fmt.Println("-----移动指针并打印当前元素-----")
	r = r.Move(2)
	fmt.Println("当前元素:", r.Value)

	fmt.Println("-----在当前元素后插入一个新环-----")
	s := ring.New(2)
	s.Value = 100
	s.Next().Value = 200
	r.Link(s)

	fmt.Println("-----遍历元素-----")
	r.Do(func(p interface{}) {
		fmt.Println(p)
	})

	fmt.Println("----- 删除当前元素后的 2 个元素-----")
	r.Unlink(2)
	fmt.Println("-----遍历环并打印每个元素的值-----")

	r.Do(func(p interface{}) {
		fmt.Println(p)
	})

}
