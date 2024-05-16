package main

import (
	"fmt"
	"reflect"
)

/*
Comparable
判断可比较
*/
func main() {
	type C = map[int]int
	var b int = 1
	c := C{}
	fmt.Println(reflect.TypeOf(b).Comparable()) // true
	fmt.Println(reflect.TypeOf(c).Comparable()) // false  map不可比较
}
