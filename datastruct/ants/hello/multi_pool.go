package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/panjf2000/ants/v2"
)

/*
Use the MultiPool and set the capacity of the 10 goroutine pools to unlimited.
If you use -1 as the pool size parameter, the size will be unlimited.
There are two load-balancing algorithms for pools: ants.RoundRobin and ants.LeastTasks.
*/
func main() {
	defer ants.Release()

	runTimes := 1000
	var wg sync.WaitGroup
	syncCalculateSum := func() {
		time.Sleep(10 * time.Millisecond)
		fmt.Println("Hello World!")
		wg.Done()
	}
	mp, _ := ants.NewMultiPool(10, -1, ants.RoundRobin)
	defer mp.ReleaseTimeout(5 * time.Second)
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		_ = mp.Submit(syncCalculateSum)
	}
	wg.Wait()
	fmt.Printf("running goroutines: %d\n", mp.Running())
	fmt.Printf("finish all tasks.\n")
}
