package main

import (
	"fmt"
	"time"
)

var x, y int

func main() {

	var x, y int
	f2 := func() {
		go func() {
			fmt.Println(x) // 可能打印出0、123，或其它值

		}()
		go func() {
			fmt.Println(y) // 可能打印出0、789，或其它值

		}()
		x, y = 123, 789
	}

	f2()
	time.Sleep(time.Second / 4)
}
