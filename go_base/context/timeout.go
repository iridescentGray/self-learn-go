package main

import (
	"context"
	"fmt"
	"time"
)

func doSomething(ctx context.Context) {
	ctx_timeout, cancel_timeout := context.WithTimeout(ctx, time.Second)
	defer cancel_timeout()
	select {
	case <-time.After(2 * time.Second):
		fmt.Println("Completed")
	case <-ctx_timeout.Done():
		fmt.Println("Context Canceled")
	}
}

/*
Timeout 处理并发编程中的超时
*/
func main() {
	//创建一个具有截止时间的context
	ctx_bak := context.Background()

	go doSomething(ctx_bak)

	time.Sleep(5 * time.Second)

}
