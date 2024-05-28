package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
Go1.19前
*/
func main() {
	var n int32
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt32(&n, 1)
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println(atomic.LoadInt32(&n)) // 1000
}
