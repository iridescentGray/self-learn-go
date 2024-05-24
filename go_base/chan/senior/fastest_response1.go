package main

import (
	"fmt"
	"math/rand"
	"time"
)

func source(c chan<- int32) {
	time.Sleep(time.Duration(rand.Intn(3)+1) * time.Second)
	c <- rand.Int31n(100)
}

/*
通道使用最快返回的那个值
*/
func main() {
	startTime := time.Now()
	c := make(chan int32, 5) // 必须用一个缓冲通道
	for i := 0; i < cap(c); i++ {
		go source(c)
	}
	rnd := <-c // 只有第一个回应被使用了
	fmt.Println(rnd)

	fmt.Println(time.Since(startTime))
}
