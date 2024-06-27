package main

import (
	"fmt"
	"time"
)

/*
接受一个回调函数，在延迟时间之后调用它
*/
func main() {
	fmt.Println("Starting...")

	time.AfterFunc(2*time.Second, func() {
		fmt.Println("2 seconds passed")
	})
	time.Sleep(time.Minute)
}
