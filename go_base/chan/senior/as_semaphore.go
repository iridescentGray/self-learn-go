package main

import (
	"fmt"
)

/*
容量为1的chan可以作为mutex lock(互斥锁)
*/
func main() {
	mutex := make(chan struct{}, 1) // 容量必须为1

	counter := 0
	increase1000 := func() {
		mutex <- struct{}{} // 加锁
		counter++
		<-mutex // 解锁
	}
	for i := 0; i < 1000; i++ {
		go increase1000()
	}
	select {}
	fmt.Println(counter) // 200000 如果不加锁 结果不会是 200000
}
