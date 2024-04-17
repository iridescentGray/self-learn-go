package main

import "fmt"

func double(x int) {
	x += x
}

func double_pointer(x *int) {
	*x += *x
}

func main() {
	var a = 3
	double(a)
	fmt.Println(a) // 3

	double_pointer(&a)
	fmt.Println(a) //6
}
