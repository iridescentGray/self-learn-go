package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	var x int32 = 0
	for i := 0; i < 100; i++ {
		go func() {
			wg.Add(1) //this bad
			atomic.AddInt32(&x, 1)
			wg.Done()
		}()
	}

	fmt.Println("等待片刻...")
	wg.Wait()
	fmt.Println(atomic.LoadInt32(&x))
}
