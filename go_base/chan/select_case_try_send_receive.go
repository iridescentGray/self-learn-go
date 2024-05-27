package main

import "fmt"

func main() {
	c := make(chan string, 2)
	trySend := func(v string) {
		select {
		case c <- v:
		default: // 如果c的缓冲已满,则执行默认分支。
			fmt.Println("trySend:default")
		}
	}
	tryReceive := func() string {
		select {
		case v := <-c:
			return v
		default:
			return "tryReceive:default" // 如果c的缓冲为空,则执行默认分支。
		}
	}
	trySend("Hello!") // 发送成功
	trySend("Hi!")    // 发送成功
	trySend("Bye!")   // trySend:default
	// 下面这两行将接收成功。
	fmt.Println(tryReceive()) // Hello!
	fmt.Println(tryReceive()) // Hi!
	// 下面这行将接收失败。
	fmt.Println(tryReceive()) // tryReceive:default-
}
