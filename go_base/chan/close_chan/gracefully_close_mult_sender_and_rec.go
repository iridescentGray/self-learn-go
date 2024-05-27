package main

import (
	"log"
	"math/rand"
	"strconv"
	"sync"
)

/*
优雅关闭 M个接收者 N个发送者
通过一个中间调解者通道,决定关闭
*/
func main() {
	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(10)
	dataCh := make(chan int)
	stopCh := make(chan struct{})
	toStop := make(chan string, 1) //中间调停者
	var stoppedBy string

	// 中间调解者负责关闭stopCh通道
	go func() {
		stoppedBy = <-toStop
		close(stopCh)
	}()

	// 发送者
	for i := 0; i < 100; i++ {
		go func(id string) {
			for {
				value := rand.Intn(1000)
				if value == 0 {
					select {
					case toStop <- "发送者#" + id:
					default:
					}
					return
				}

				// 为了让此发送协程尽早退出(编译器有特殊的优化速度很快)
				select {
				case <-stopCh:
					return
				default:
				}

				select {
				case <-stopCh:
					return
				case dataCh <- value:
				}
			}
		}(strconv.Itoa(i))
	}

	// 接收者
	for i := 0; i < 10; i++ {
		go func(id string) {
			defer wgReceivers.Done()
			for {
				select {
				case <-stopCh:
					return
				default:
				}

				select {
				case <-stopCh:
					return
				case value := <-dataCh:
					if value == 1000-1 {
						select {
						case toStop <- "接收者#" + id:
						default:
						}
						return
					}
					log.Println(value)
				}
			}
		}(strconv.Itoa(i))
	}

	wgReceivers.Wait()
	log.Println("被" + stoppedBy + "终止了")
}
