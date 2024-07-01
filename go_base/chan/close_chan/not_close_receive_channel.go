package main

import "fmt"

func main() {
	ch := make(chan int)
	fmt.Println(ch)
}

func foo(c <-chan int) {
	close(c) // error: 不能关闭单向接收通道
}
