package main

import (
	"fmt"
	"sync"
)

/*
Cond用于实现多个协程间的通知
*/
func main() {

	cond := sync.NewCond(&sync.Mutex{})
	cond.L.Lock()
	go func() {
		cond.L.Lock()
		go func() {
			cond.L.Lock()
			fmt.Print("0")
			cond.Broadcast()
			cond.L.Unlock()
		}()
		cond.Wait()
		fmt.Print("a")
		cond.L.Unlock()

	}()
	cond.Wait()
	fmt.Print("b")

	cond.L.Unlock()
	fmt.Println("c")
}
