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
	for i, n := range []int{2, 3, 4} {
		defer func() {
			fmt.Println(i, n)
		}()
	}
	fmt.Println("--------------")

	var m = map[*int]int{}
	for i, n := range []int{5, 6, 7} {
		m[&i]++
		m[&n]++
	}
	fmt.Println(len(m)) //6

}
