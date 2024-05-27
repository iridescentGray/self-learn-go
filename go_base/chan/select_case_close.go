package main

import "fmt"

/*
对select使用已关闭通道
恐慌随机出现(执行顺序不定)
*/
func main() {
	var c chan struct{} = make(chan struct{})
	close(c)
	select {
	case <-c:
		fmt.Println("receive")
	case c <- struct{}{}:
		fmt.Println("send")
	default:
		fmt.Println("default")
	}
}
