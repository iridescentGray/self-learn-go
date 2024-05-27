package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

/*
优雅关闭 M个接收者、一个发送者 和 N个第三方
引入 中介通道
*/
func main() {
	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(100)
	dataCh := make(chan int)
	mid_chain := make(chan struct{}) // 中介通道
	closed := make(chan struct{})

	// 此stop函数能安全地多次调用
	stop := func() {
		select {
		case mid_chain <- struct{}{}:
			<-closed
		case <-closed: //第三方阻塞等待直到发送者关闭closed通道
		}
	}

	// 第三方协程
	for i := 0; i < 15; i++ {
		go func() {
			time.Sleep(time.Duration(1+rand.Intn(3)) * time.Second)
			stop()
		}()
	}

	// 发送者
	go func() {
		defer func() {
			close(closed) //关闭此通道 第三方协程才不会阻塞
			close(dataCh)
			fmt.Println("发送通道已关闭")
		}()

		for {
			select {
			case <-mid_chain:
				return
			default:
			}

			select {
			case <-mid_chain: //关闭信号
				return
			case dataCh <- rand.Intn(1000):
			}
		}
	}()

	// 接收者
	for i := 0; i < 100; i++ {
		go func() {
			defer wgReceivers.Done()
			for value := range dataCh {
				log.Println(value)
			}
		}()
	}
	wgReceivers.Wait()
}
