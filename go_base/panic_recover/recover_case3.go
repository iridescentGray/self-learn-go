package main

import (
	"log"
	"time"
)

/*
恐慌后,自动重启
*/
func shouldNotExit() {
	for {
		time.Sleep(time.Second)
		if 1 == 1 {
			panic("unexpected situation") // 未预料到的恐慌
		}
	}
}

func NeverExit(name string, shouldNotExit func()) {
	defer func() {
		if v := recover(); v != nil { // 侦测到一个恐慌
			log.Printf("协程%s崩溃了,准备重启一个", name)
			go NeverExit(name, shouldNotExit) // 重启一个同功能协程
		}
	}()
	time.Sleep(time.Second)
	shouldNotExit()
}

func main() {
	go NeverExit("job-------A", shouldNotExit)
	select {} // 永久阻塞主线程
}
