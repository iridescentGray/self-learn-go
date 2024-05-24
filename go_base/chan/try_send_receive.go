package main

import "fmt"

/*
带一个default+case分支的select代码块可以被尝试发送/接收
标准编译器对尝试发送/接收特别的优化，它们的执行效率比多case分支的普通select代码块执行效率高
*/
func main() {
	type Book struct{ id int }
	bookshelf := make(chan Book, 3)

	//尝试发送
	for i := 0; i < 6; i++ {
		select {
		case bookshelf <- Book{id: i}:
			fmt.Println("成功将书放在书架上", i)
		default:
			fmt.Println("书架已经被占满了")
		}
	}

	//尝试接收
	for i := 0; i < 6; i++ {
		select {
		case book := <-bookshelf:
			fmt.Println("成功从书架上取下一本书", book.id)
		default:
			fmt.Println("书架上已经没有书了")
		}
	}
}
