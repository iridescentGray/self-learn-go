package main

import (
	"fmt"
	"runtime"
	"time"
)

func memoryLeaking() {
	type T1 struct {
		v [1 << 20]int
		t *T1
	}

	var finalizer = func(t *T1) {
		fmt.Println("finalizer called")
	}

	var x, y T1

	// 此SetFinalizer函数调用将使x逃逸到堆上。
	runtime.SetFinalizer(&x, finalizer)

	// 下面这行将形成一个包含x和y的循环引用值组。
	// 这有可能造成x和y不可回收。
	x.t, y.t = &y, &x // y也逃逸到了堆上。
}

/*
子字符串内存泄露
*/
func main() {
	go memoryLeaking()
	time.Sleep(time.Second)
}
