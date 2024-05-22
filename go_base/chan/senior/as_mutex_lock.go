package main

import "fmt"

/*
容量为1的chan可以作为mutex lock(互斥锁)
*/
func main() {
	mutex := make(chan struct{}, 1) // 容量必须为1

	counter := 0
	increase1000 := func(done chan<- struct{}) {
		for i := 0; i < 100000; i++ {
			mutex <- struct{}{} // 加锁
			counter++
			<-mutex // 解锁
		}
		done <- struct{}{}
	}

	done := make(chan struct{})
	go increase1000(done)
	go increase1000(done)
	<-done //因为有两个 increase1000 所以有两个<-done
	<-done
	fmt.Println(counter) // 200000 如果不加锁 结果不会是 200000
}
