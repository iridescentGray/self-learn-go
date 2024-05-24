package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
Cond用于实现多个协程间的通知
sync.Cond值拥有一个sync.Locker类型的名为L的字段(Mutex/RWMutex),并且维护着一个先进先出等待协程队列

它有三个方法：
1.Wait()

	必须在Cond.L加锁时调用(否则造成一个恐慌)
	a.将当前协程推入到c所维护的等待协程队列；
	b.调用c.L.Unlock()对c.L的锁解锁；
	c.然后使当前协程进入阻塞状态；
	d.当前协程将被另一个协程通过Signal/Broadcast唤醒后重新进入运行状态，将调用c.L.Lock()重新加锁。加锁成功后wait逻辑结束

2.Signal()

	调用将唤醒并移除c所维护的等待协程队列中的第一个协程

3.Broadcast()

	唤醒并移除c所维护的等待协程队列中的所有协程
*/
func main() {
	const N = 10
	var values [N]string
	cond := sync.NewCond(&sync.Mutex{})

	for i := 0; i < N; i++ {
		go func(i int) {
			time.Sleep(time.Second * time.Duration(rand.Intn(10)) / 10)
			//等待通知
			cond.L.Lock()

			values[i] = string('a' + i)
			cond.Broadcast()
			cond.L.Unlock()
		}(i)
	}

	checkCondition := func() bool {
		fmt.Println(values)
		for i := 0; i < N; i++ {
			if values[i] == "" {
				return false
			}
		}
		return true
	}

	cond.L.Lock()
	defer cond.L.Unlock()
	for !checkCondition() {
		cond.Wait()
	}
}
