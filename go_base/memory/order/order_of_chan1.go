package main

import (
	"time"
)

/*
赋值语句b = 1肯定在条件b != 1被估值之前执行完毕。
赋值语句a = 1肯定在条件a != 1被估值之前执行完毕。
*/
func main() {

	f3 := func() {
		var a, b int
		var c = make(chan bool)

		go func() {
			a = 1
			c <- true
			if b != 1 {
				panic("b != 1") // 绝不可能发生
			}
		}()

		go func() {
			b = 1
			<-c //阻塞
			if a != 1 {
				panic("a != 1") // 绝不可能发生
			}
		}()
	}

	f3()
	time.Sleep(time.Second / 4)
}
