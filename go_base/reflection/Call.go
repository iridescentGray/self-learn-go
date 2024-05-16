package main

import (
	"fmt"
	"reflect"
)

/*
Call
调用函数
*/
type T struct {
	A, b int
}

func (t T) Add(n int) int {
	return n + t.A + t.b
}

func main() {
	t := T{5, 2}
	vt := reflect.ValueOf(t)
	vm := vt.MethodByName("Add")
	results := vm.Call([]reflect.Value{reflect.ValueOf(3)})
	fmt.Println(results[0].Int()) // 5+2+3=10
}
