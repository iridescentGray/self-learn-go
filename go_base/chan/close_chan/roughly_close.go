package main

import (
	"fmt"
)

func RoughlyClose(ch chan int) (justClosed bool) {
	defer func() {
		if recover() != nil {
			justClosed = false
		}
	}()
	close(ch)   // 如果ch已关闭，则产生一个恐慌。
	return true // <=> justClosed = true; return
}

/*
使用异常关闭通道
*/
func main() {
	c := make(chan int)
	fmt.Println(RoughlyClose(c)) //true
	fmt.Println(RoughlyClose(c)) //false

}
