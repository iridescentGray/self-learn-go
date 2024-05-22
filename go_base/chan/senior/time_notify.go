package main

import (
	"fmt"
	"time"
)

func AfterDuration(d time.Duration) <-chan struct{} {
	c := make(chan struct{}, 1)
	go func() {
		time.Sleep(d)
		c <- struct{}{}
	}()
	return c
}

/*
等待指定时间后,收到通知
*/
func main() {
	fmt.Println("Hi!")
	<-AfterDuration(time.Second)
	fmt.Println("Hello!")
	<-AfterDuration(time.Second)
	fmt.Println("Bye!")
}
