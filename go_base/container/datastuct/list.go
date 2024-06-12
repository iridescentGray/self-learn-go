package main

import (
	"container/list"
	"fmt"
)

/*
双向链表

* list没有泛型,取值时要强制类型转换
*/
func main() {
	// 创建一个新的链表
	myList := list.New()

	// 添加元素到链表
	myList.PushBack("元素1")
	myList.PushFront("元素2")
	myList.PushBack("元素3")

	fmt.Println("-----第一次遍历-----")
	// 遍历链表
	for element := myList.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}

	// 删除元素
	element := myList.Front() // 获取第一个元素
	myList.Remove(element)    // 删除该元素
	fmt.Println("-----第二次遍历-----")
	// 再次遍历链表
	for element := myList.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}
}
