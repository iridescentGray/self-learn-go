package main

import (
	"fmt"
	"runtime"
	"time"
)

func f() {
	defer func() {
		fmt.Println(recover()) //<nil>
	}()

	// Goexit会恢复之前产生的恐慌。
	defer runtime.Goexit()
	panic("bye")
}

func main() {
	go f()
	time.Sleep(time.Second)
}
