package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/panjf2000/ants/v2"
)

func main() {
	defer ants.Release()

	runTimes := 1000
	var wg sync.WaitGroup
	syncCalculateSum := func() {
		time.Sleep(10 * time.Millisecond)
		fmt.Println("Hello World!")
		wg.Done()
	}
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		_ = ants.Submit(syncCalculateSum)
	}
	wg.Wait()
	fmt.Printf("running goroutines: %d\n", ants.Running())
	fmt.Printf("finish all tasks.\n")
}
