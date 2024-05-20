package main

import (
	"fmt"
)

func main() {
	type Person struct {
		name string
		age  int
	}
	persons := map[string]int{"Alice": 28, "Bob": 25}
	// persons := [2]Person{{"Alice", 28}, {"Bob", 25}}
	for k, v := range persons {
		fmt.Println(k, v)

		// 此修改不会反映到persons数组中,因为p
		// 是persons数组的副本中的一个元素的副本。
		v = 31
	}
	fmt.Println("persons:", persons)
	for k := range persons {
		fmt.Println(k)
	}
}
