package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func SayGreetings(greeting string, times int) {
	for i := 0; i < times; i++ {
		log.Println(greeting)
		d := time.Second * time.Duration(rand.Intn(3)) / 2
		time.Sleep(d)
	}
	wg.Done() // 通知当前任务已经完成。
}

func main() {
	log.SetFlags(0)
	wg.Add(2) // 注册两个新任务。
	go SayGreetings("hi!", 5)
	go SayGreetings("hello!", 5)
	wg.Wait() // 阻塞在这里,直到所有任务都已完成。
}
