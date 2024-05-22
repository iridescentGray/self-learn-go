package main

import (
	"fmt"
	"time"
)

/*
一对一 等待通道接收 实现通知
*/
func main() {
	done := make(chan struct{})

	go func() {
		fmt.Print("Hello")
		time.Sleep(time.Second * 2)
		<-done
	}()

	done <- struct{}{} // 阻塞在此，等待通知
	fmt.Println(" world!")
}
