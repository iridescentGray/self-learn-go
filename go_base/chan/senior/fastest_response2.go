package main

import (
	"fmt"
	"math/rand"
	"time"
)

func source(c chan<- int32) {
	time.Sleep(time.Duration(rand.Intn(3)+1) * time.Second)
	select {
	case c <- rand.Int31n(100):
	default:
	}
}

func main() {
	startTime := time.Now()
	c := make(chan int32, 1) // 此通道容量必须至少为1
	for i := 0; i < 5; i++ {
		go source(c)
	}
	rnd := <-c // 只采用第一个成功发送的回应数据
	fmt.Println(rnd)

	fmt.Println(time.Since(startTime))
}
