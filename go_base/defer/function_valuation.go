package main

import "fmt"

func main() {
	//will be call
	var f = func() {
		fmt.Println(false)
	}
	defer f()
	//not be call
	f = func() {
		fmt.Println(true)
	}
}
