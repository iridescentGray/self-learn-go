package main

import (
	"fmt"
	"time"
)

func main() {
	// 假设此切片的长度很大，以至于它的元素
	// 将被开辟在堆上。
	bs := make([]byte, 1<<31)

	// 一个聪明的编译器将觉察到bs的底层元素不会再被使用,底层元素部分在此刻可以被安全地回收了
	time.Sleep(time.Second)
	fmt.Println("Recycle")
	fmt.Println(len(bs))
}
