package main

import (
	"crypto/rand"
	"fmt"
	"sort"
)

/*
一对一 等待通道发送 实现通知
*/
func main() {
	values := make([]byte, 1024)
	//write random byte
	rand.Read(values)
	fmt.Println(values[0], values[len(values)-1])

	done := make(chan struct{})
	go func() {
		sort.Slice(values, func(i, j int) bool {
			return values[i] < values[j]
		})
		done <- struct{}{} // 通知排序已完成
	}()
	<-done //wait for notic
	fmt.Println(values[0], values[len(values)-1])
}
