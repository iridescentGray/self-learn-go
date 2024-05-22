package main

import (
	"fmt"
	"math/rand"
	"time"
)

func longTimeRequest() <-chan int32 {
	r := make(chan int32)
	go func() {
		time.Sleep(time.Second * 3) // 模拟一个工作负载
		r <- rand.Int31n(100)
	}()

	return r
}

/*
通道也能作为函数的返回值
*/
func main() {
	a, b := longTimeRequest(), longTimeRequest()
	c, d := <-a, <-b
	fmt.Println(c + d)
}
