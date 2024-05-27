package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
)

/*
优雅关闭 一个接收者 N个发送者
接收者关闭
*/
func main() {
	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(1)
	dataCh := make(chan int)
	stopCh := make(chan struct{}) //接受关闭信号
	// 发送者
	for i := 0; i < 100; i++ {
		go func() {
			for {
				// 为了让此发送协程尽早退出(编译器有特殊的优化速度很快)
				select {
				case <-stopCh:
					return
				default:
				}

				select {
				case <-stopCh: //从已关闭通道中接收,不会阻塞
					return
				case dataCh <- rand.Intn(3000):
				}
			}
		}()
	}

	// 接收者
	go func() {
		defer wgReceivers.Done()
		for value := range dataCh {
			if value == 3000-1 {
				close(stopCh) //关闭信号
				fmt.Println("通道已关闭")
				return
			}
			log.Println(value)
		}
	}()
	wgReceivers.Wait()
}
