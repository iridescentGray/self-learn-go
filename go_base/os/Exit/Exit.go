// exit-example.go
package main

import (
	"os"
	"time"
)

func main() {
	go func() {
		time.Sleep(time.Second)
		os.Exit(1) // 从任何函数里退出一个程序
	}()
	select {}
}
