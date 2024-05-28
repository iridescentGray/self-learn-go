package main

import (
	"fmt"
	"time"
)

/*
 */
func longRunning2(messages <-chan int) {
	timer := time.NewTimer(time.Second)
	defer timer.Stop()

	for {
		select {
		case <-timer.C: // 过期了
			return
		case msg := <-messages:
			fmt.Println(msg)

			// 此if代码块很重要。
			if !timer.Stop() {
				<-timer.C
			}
		}

		// 必须重置以复用。
		timer.Reset(time.Second)
	}
}
func main() {

	msg := make(chan int, 10)
	for i := 0; i < 10; i++ {
		msg <- i
		longRunning2(msg)
	}

}
