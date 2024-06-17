package main

import (
	"fmt"
	"runtime"
)

func main() {
	c := make(chan int)
	go func() {
		defer func() { c <- 1 }()
		defer fmt.Println("Go") //正常打印
		func() {
			defer fmt.Println("C") //正常打印
			runtime.Goexit()       //退出当前携程
		}()
		fmt.Println("Java") // 已退出,不可达
	}()
	<-c
}
