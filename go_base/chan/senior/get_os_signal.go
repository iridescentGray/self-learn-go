package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {

	var wantG sync.WaitGroup
	wantG.Add(1)
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("get Signal")
		wantG.Done()
	}()
	fmt.Println("wait sign")
	wantG.Wait()
	fmt.Println("end")
}
