package main

import (
	"fmt"
	"sync"
	"time"
)

type rwCounter struct {
	m sync.RWMutex
	n uint64
}

func (c *rwCounter) rwValue() uint64 {

	c.m.RLock()
	defer c.m.RUnlock()
	return c.n
}

func (c *rwCounter) rwIncrease(delta uint64) {
	c.m.Lock()
	c.n += delta
	c.m.Unlock()
}

/*
RWMutex值常称为一个读写互斥锁
内部包含两个锁：一个写锁和一个读锁，Lock()加写锁，RLock()加读锁。

写锁只有在它的写锁和读锁都处于未加锁状态时才能被成功加锁
加写锁后，新的加写/读锁将导致当前协程阻塞
加读锁后，新加写锁的将导致当前协程阻塞。 但加读锁的会成功
*/
func main() {
	var c rwCounter
	for i := 0; i < 100; i++ {
		go func() {
			for k := 0; k < 100; k++ {
				c.rwIncrease(1)
			}
		}()
	}
	time.Sleep(time.Second)
	fmt.Println(c.rwValue()) // 10000
}
