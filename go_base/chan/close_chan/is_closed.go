package main

import (
	"fmt"
)

func isClosed(c chan int) bool {
	select {
	case <-c:
		return true
	default:
	}
	return false
}

func main() {
	c := make(chan int)
	fmt.Println(isClosed(c)) //false
	close(c)
	fmt.Println(isClosed(c)) //true

}
