package main

import (
	"fmt"
	"time"
)

func main() {

	go func() {
		time.Sleep(time.Second)
		defer func() {
			v := recover()
			fmt.Println("恐慌被恢复了：", v)
		}()
		panic(123)
	}()

	for {
		fmt.Println("main loop!")
		time.Sleep(time.Second)
	}
}
