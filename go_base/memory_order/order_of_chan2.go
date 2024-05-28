package main

import (
	"fmt"
	"time"
)

/*
赋值语句b = 1肯定在条件b != 1被估值之前执行完毕。
赋值语句a = 1肯定在条件a != 1被估值之前执行完毕。
*/
func main() {

	var n = 123

	c := make(chan int)

	go func() {
		c <- n + 0 //123
		// c <- n //789
	}()

	time.Sleep(time.Second)
	n = 789
	fmt.Println(<-c)
}
