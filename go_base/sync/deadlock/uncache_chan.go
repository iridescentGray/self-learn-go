package main

import (
	"fmt"
	"time"
)

func main() {
	request := func() int {
		c := make(chan int)
		for i := 0; i < 5; i++ {
			i := i
			go func() {
				select {
				case c <- i:
				default: //非阻塞发送
				}
			}()
		}
		time.Sleep(time.Second)
		return <-c // 错过了发送,接收时,永久阻塞
	}
	fmt.Println(request())
}
