package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start")
	c := time.After(time.Second) //阻塞一秒
	fmt.Println(<-c)
}
