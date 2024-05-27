package main

import (
	"fmt"
	"sync"
)

type MyChannel struct {
	C    chan int
	once sync.Once
}

func (mc *MyChannel) SafeClose() {
	mc.once.Do(func() {
		close(mc.C)
	})
}

func isClosed2(c chan int) bool {
	select {
	case <-c:
		return true
	default:
	}
	return false
}

/*
礼貌关闭

缺点:
SafeClose函数和同一个通道上的发送操作同时运行，则数据竞争可能发生
*/
func main() {
	mc := MyChannel{C: make(chan int)}
	fmt.Println(isClosed2(mc.C)) //false
	mc.SafeClose()
	fmt.Println(isClosed2(mc.C)) //true
}
