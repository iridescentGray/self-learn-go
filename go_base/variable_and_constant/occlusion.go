package main

import "fmt"

var f = func(b bool) {
	fmt.Print("Goat")
}

/*
若旧标识符被新标识符遮挡，新标识符内部的作用域不会被遮挡,依旧是旧标识符生效
*/
func main() {
	var f = func(b bool) {
		fmt.Print("Sheep ")
		if b {
			f(!b) //被包变量 f 覆盖
		}
	}
	f(true) //Sheep Goat
}
