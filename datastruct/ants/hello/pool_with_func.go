package main

import (
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/panjf2000/ants/v2"
)

var sum1 int32

/*
Use the pool with a function
set 10 to the capacity of goroutine pool and 1 second for expired duration.
*/
func main() {
	defer ants.Release()

	runTimes := 1000
	var wg sync.WaitGroup
	p, _ := ants.NewPoolWithFunc(10, func(i interface{}) {
		n := i.(int32)
		atomic.AddInt32(&sum1, n)
		fmt.Printf("run with %d\n", n)
		wg.Done()
	})
	defer p.Release()
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		p.Invoke(int32(i))
	}
	wg.Wait()
	fmt.Printf("running goroutines: %d\n", p.Running())
	fmt.Printf("finish all tasks, result is %d\n", sum1)
	if sum1 != 499500 {
		panic("the final result is wrong!!!")
	}
}
