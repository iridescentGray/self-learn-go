package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
)

/*
优雅关闭 M个接收者和一个发送者
由发送者关闭
*/
func main() {
	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(100)
	dataCh := make(chan int)

	// 发送者
	go func() {
		for {
			if value := rand.Intn(1000); value == 0 {
				close(dataCh)
				fmt.Println("通道已关闭")
				return
			} else {
				dataCh <- value
			}
		}
	}()

	// 接收者
	for i := 0; i < 100; i++ {
		go func() {
			defer wgReceivers.Done()
			// 接收数据直到通道dataCh已关闭
			for value := range dataCh {
				log.Println(value)
			}
		}()
	}

	wgReceivers.Wait()
}
