package main

import (
	"log"
	"time"
)

/*
一对多 等待通道接收 实现通知
*/
func worker(id int, ready <-chan struct{}, done chan<- struct{}) {
	<-ready // 阻塞在此，等待通知
	log.Print("Worker#", id, "开始工作")
	time.Sleep(time.Second * time.Duration(id+1))
	log.Print("Worker#", id, "工作完成")
	done <- struct{}{} // 通知主协程（N-to-1）
}

func main() {
	log.SetFlags(0)

	ready, done := make(chan struct{}), make(chan struct{})
	go worker(0, ready, done)
	go worker(1, ready, done)
	go worker(2, ready, done)

	ready <- struct{}{}
	ready <- struct{}{}
	ready <- struct{}{}
	// close(ready)  关闭等价于上述三条,因为关闭会发送零值
	<-done
	<-done
	<-done
	time.Sleep(time.Second * 3)

}
