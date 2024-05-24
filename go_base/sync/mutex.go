package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	m sync.Mutex
	n uint64
}

func (c *Counter) Value() uint64 {
	c.m.Lock()
	defer c.m.Unlock()
	return c.n
}

func (c *Counter) Increase(delta uint64) {
	c.m.Lock()
	c.n += delta
	c.m.Unlock()
}

/*
Mutex是互斥锁,一旦Mutex被加了锁，其他新的加锁调用将导致当前协程阻塞
Mutex的零值是一个尚未加锁的互斥锁
*/
func main() {
	var c Counter
	for i := 0; i < 100; i++ {
		go func() {
			for k := 0; k < 100; k++ {
				c.Increase(1)
			}
		}()
	}
	time.Sleep(time.Second)
	fmt.Println(c.Value()) // 10000
}
