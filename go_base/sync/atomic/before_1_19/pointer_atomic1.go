package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Page struct {
	views uint32
}

func (page *Page) AddViews(n uint32) {
	atomic.AddUint32(&page.views, n)
}

func (page *Page) Views() uint32 {
	return atomic.LoadUint32(&page.views)
}

/*
Go1.19前
*/
func main() {
	n := Page{views: 123}
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			n.AddViews(uint32(i))
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println(n) // 499623
}
