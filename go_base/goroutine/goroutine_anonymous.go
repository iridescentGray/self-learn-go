package main

import (
	"log"
	"math/rand"
	"time"
)

func main() {
	//启动一个匿名协程
	go func() {
		for i := 0; i < 10; i++ {
			log.Println("1-hi!")
			d := time.Second * time.Duration(rand.Intn(5)) / 2
			time.Sleep(d) // 睡眠片刻（随机0到2.5秒）
		}
	}()

	time.Sleep(10 * time.Second)

}
