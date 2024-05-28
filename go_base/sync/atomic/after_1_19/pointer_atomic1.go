package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Page struct {
	views atomic.Uint32
}

func (page *Page) AddViews(n uint32) {
	page.views.Add(n)
}

func (page *Page) Views() uint32 {
	return page.views.Load()
}

/*
Go1.19前
*/
func main() {
	n := Page{views: atomic.Uint32{}}
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			n.AddViews(uint32(i))
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println(n.views) // 4950
}
