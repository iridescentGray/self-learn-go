package main

import (
	"fmt"
	"math/rand"
	"time"
)

func source() <-chan int32 {
	c := make(chan int32, 1) // 必须为一个缓冲通道
	go func() {
		time.Sleep(time.Duration(rand.Intn(3)+1) * time.Second)
		c <- rand.Int31n(100)
	}()
	return c
}

func main() {
	var rnd int32
	// 阻塞在此直到某个数据源率先回应
	select {
	case rnd = <-source():
	case rnd = <-source():
	case rnd = <-source():
	}
	fmt.Println(rnd)
}
