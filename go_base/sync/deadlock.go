package main

import (
	"sync"
	"time"
)

var wg_deadlock sync.WaitGroup

func main() {
	wg_deadlock.Add(1)
	go func() {
		time.Sleep(time.Second * 2)
		wg_deadlock.Wait() // 阻塞在此
	}()
	wg_deadlock.Wait() // 阻塞在此
}
