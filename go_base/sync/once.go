package main

import (
	"log"
	"sync"
)

/*
sync.Once用来确保一段代码在并发程序中仅被执行一次
Once有一个Do(f func())方法,参数为func()
*/
func main() {
	x := 0
	doSomething := func() {
		x++
		log.Println("Only Once Time")
	}

	var wg sync.WaitGroup
	var once sync.Once
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			once.Do(doSomething)
			log.Println("dodododododoo!")
		}()
	}

	wg.Wait()
	log.Println("x =", x) // x = 1
}
