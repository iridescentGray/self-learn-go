package main

import (
	"errors"
	"fmt"
	"time"
)

func doRequest(c chan int) {
	time.Sleep(time.Second * 2)
	c <- 3
}

func requestWithTimeout(timeout time.Duration) (int, error) {
	c := make(chan int)
	go doRequest(c) // 可能需要超出预期的时长回应

	select {
	case data := <-c:
		return data, nil
	case <-time.After(time.Second * timeout):
		return 0, errors.New("超时了！")
	}
}

/*
执行超时
*/
func main() {
	fmt.Println(requestWithTimeout(1))
}
