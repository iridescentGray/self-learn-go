package main

import (
	"fmt"
	"time"
)

/*
若longRunning函数被调用并且在一分钟内有一百万条消息到达， 那么在某个特定的很小时间段（大概若干秒）内将存在一百万个活跃的Timer值，即使其中只有一个是真正有用的
Go 1.23开始的Go运行时做了改进，从而使得当一个Timer对象不再被使用时将立即变得可以垃圾回收。
*/
func longRunning(messages <-chan int) {
	for {
		select {
		case <-time.After(time.Second): //如果message没有消息超过一分钟，会导致函数return
			return
		case msg := <-messages:
			fmt.Println(msg)
		}
	}
}

func main() {

	msg := make(chan int, 10)
	for i := 0; i < 10; i++ {
		msg <- i
		longRunning(msg)
	}

}
