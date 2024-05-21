package main

import "fmt"

func main() {
	c := make(chan int, 1)
	c <- 0
	fchan := func(info string, c chan int) chan int {
		fmt.Println(info)
		return c
	}
	fptr := func(info string) *int {
		fmt.Println(info)
		return new(int)
	}

	select {
	case *fptr("aaa") = <-fchan("bbb", nil): // blocking
	case *fptr("ccc") = <-fchan("ddd", c): // non-blocking
	case fchan("eee", nil) <- *fptr("fff"): // blocking
	case fchan("ggg", nil) <- *fptr("hhh"): // blocking
	}
}
