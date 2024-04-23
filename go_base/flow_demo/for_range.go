package main

import (
	"fmt"
)

func main() {

	//support by go1.22
	for i := range 10 {
		fmt.Println(i)
	}
	fmt.Println("--------------")
	for i, n := range []int{0, 1, 2} {
		defer func() {
			fmt.Println(i, n)
		}()
	}
	fmt.Println("--------------")

	var m = map[*int]int{}
	for i, n := range []int{1, 2, 3} {
		m[&i]++
		m[&n]++
	}
	fmt.Println(len(m)) //6

}
