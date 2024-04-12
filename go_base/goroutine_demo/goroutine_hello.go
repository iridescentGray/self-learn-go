package main

import (
	"log"
	"math/rand"
	"time"
)

func SayGreetings(greeting string, times int) {
	for i := 0; i < times; i++ {
		log.Println(greeting)
		d := time.Second * time.Duration(rand.Intn(5)) / 2
		time.Sleep(d) // 睡眠片刻（随机0到2.5秒）
	}
}

func main() {
	log.SetFlags(0)
	go SayGreetings("1-hi!", 10)
	go SayGreetings("2-hello!", 10)

	//睡眠后,主协程退出，不管其他协程是否结束
	time.Sleep(3 * time.Second)
}
