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

/*
sync.WaitGroup用来让某个协程等待其它协程完成它们各自的任务
sync.WaitGroup内部维护着一个计数，此计数的初始默认值为零(负数时产生恐慌)
它有三个方法：
1.Add(delta int)

	用来改变计数

2.Done()

	wg.Done()和wg.Add(-1)等价

3.Wait()

	计数为零,相当于空操作
	计数为正,当前协程阻塞,直到此计数更改为0
*/
func main() {
	wg.Add(2) // 注册两个新任务
	go SayGreetings("hi!", 3)
	go SayGreetings("hello!", 2)
	wg.Wait() // 阻塞在这里,直到所有任务都已完成
}
