package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/panjf2000/ants/v2"
)

var sum int32

/*
Use the MultiPoolFunc and set the capacity of 10 goroutine pools to (runTimes/10).
*/
func main() {
	defer ants.Release()

	runTimes := 1000
	var wg sync.WaitGroup
	mpf, _ := ants.NewMultiPoolWithFunc(10, runTimes/10, func(i interface{}) {
		n := i.(int32)
		atomic.AddInt32(&sum, n)
		fmt.Printf("run with %d\n", n)
		wg.Done()
	}, ants.LeastTasks)
	defer mpf.ReleaseTimeout(5 * time.Second)
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		_ = mpf.Invoke(int32(i))
	}
	wg.Wait()
	fmt.Printf("running goroutines: %d\n", mpf.Running())
	fmt.Printf("finish all tasks, result is %d\n", sum)
	if sum != 499500 {
		panic("the final result is wrong!!!")
	}
}
